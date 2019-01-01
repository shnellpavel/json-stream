package filter

import (
	"strings"
	"text/scanner"
	"unicode"

	"github.com/pkg/errors"
)

// ErrInvalidOperator appears when pass a string that doesn't match any known operator
var ErrInvalidOperator = errors.New("invalid operator")

type operator string

// Available operators to use in conditions
const (
	OpEq      = operator("=")
	OpNotEq   = operator("!=")
	OpLt      = operator("<")
	OpLte     = operator("<=")
	OpGt      = operator(">")
	OpGte     = operator(">=")
	OpLike    = operator("~")
	OpNotLike = operator("!~")
	OpUnknown = operator("")
)

var validOperators = []operator{OpEq, OpNotEq, OpLt, OpLte, OpGt, OpGte, OpLike, OpNotLike}

func newOperator(op string) operator {
	res := operator(strings.TrimSpace(op))
	for _, validOp := range validOperators {
		if validOp == res {
			return validOp
		}
	}

	return OpUnknown
}

func (o operator) String() string {
	return string(o)
}

const specialIdentRunes = "._=!<>~"

// Condition is parsed string expression. It used to solve inclusion of stream elem
type Condition struct {
	path     string
	operator operator
	value    string
}

// NewConditionFromStr builds condition object from string representation
func NewConditionFromStr(conditionStr string) (*Condition, error) {
	var scan scanner.Scanner
	scan.Init(strings.NewReader(conditionStr))

	scan.IsIdentRune = func(ch rune, i int) bool {
		return strings.IndexRune(specialIdentRunes, ch) > 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}
	scan.Whitespace ^= 1 << ' '

	var path, operatorStr, value string
	partNum := 0
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		if strings.TrimSpace(scan.TokenText()) == "" {
			partNum++
		}

		switch partNum {
		case 0:
			path += scan.TokenText()
		case 1:
			operatorStr += scan.TokenText()
		case 2:
			value += scan.TokenText()
		default:
			break
		}
	}

	op := newOperator(operatorStr)
	if op == OpUnknown {
		return nil, errors.Wrapf(ErrInvalidOperator, "found operator %s", operatorStr)
	}

	return &Condition{
		path:     strings.TrimSpace(path),
		operator: op,
		value:    strings.TrimSpace(value),
	}, nil
}
