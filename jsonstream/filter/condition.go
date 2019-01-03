package filter

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// ErrInvalidOperator appears when pass a string that doesn't match any known operator
var ErrInvalidOperator = errors.New("invalid operator")

// Operator is part of condition expression that is used to compare two operands
type Operator string

// Available operators to use in conditions
const (
	OpEq      = Operator("=")
	OpNotEq   = Operator("!=")
	OpLt      = Operator("<")
	OpLte     = Operator("<=")
	OpGt      = Operator(">")
	OpGte     = Operator(">=")
	OpLike    = Operator("~")
	OpNotLike = Operator("!~")
	OpUnknown = Operator("")
)

var validOperators = []Operator{OpEq, OpNotEq, OpLt, OpLte, OpGt, OpGte, OpLike, OpNotLike}

var conditionRegexp = regexp.MustCompile(`^\s*(.+?)\s*(!?=|=|<=?|>=?|!?~)\s*(.*?)\s*$`)

func newOperator(op string) Operator {
	res := Operator(strings.TrimSpace(op))
	for _, validOp := range validOperators {
		if validOp == res {
			return validOp
		}
	}

	return OpUnknown
}

// String casts operator to string
func (o Operator) String() string {
	return string(o)
}

// Condition is parsed string expression. It used to solve inclusion of stream elem
type Condition struct {
	path     string
	operator Operator
	value    string
}

// Path returns path to left operand of condition
func (c Condition) Path() string {
	return c.path
}

// Operator returns operator of condition
func (c Condition) Operator() Operator {
	return c.operator
}

// Value returns value (right operand of condition)
func (c Condition) Value() string {
	return c.value
}

// NewConditionFromStr builds condition object from string representation
func NewConditionFromStr(conditionStr string) (*Condition, error) {
	found := conditionRegexp.FindStringSubmatch(conditionStr)
	if found == nil {
		return nil, errors.New("invalid expression")
	}

	path := found[1]
	operatorStr := found[2]
	value := found[3]

	op := newOperator(operatorStr)
	if op == OpUnknown {
		return nil, errors.Wrapf(ErrInvalidOperator, "found operator %s", operatorStr)
	}

	return &Condition{
		path:     strings.TrimSpace(path),
		operator: op,
		value: strings.TrimFunc(value, func(c rune) bool {
			return unicode.IsSpace(c) || c == '\'' || c == '"'
		}),
	}, nil
}
