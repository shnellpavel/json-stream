package filter

import (
	"strconv"

	"github.com/Jeffail/gabs"
	"github.com/pkg/errors"
)

var (
	// ErrUnsupportedOperator represents case of using unsupported operator (app or type of value by path)
	ErrUnsupportedOperator = errors.New("unsupported operator")

	// ErrUnsupportedType represents case appearing unknown type of field value
	ErrUnsupportedType = errors.New("unsupported type")
)

// ProcessElem solves accordance of stream element to condition
func ProcessElem(condition Condition, elem []byte) (resElem string, isOk bool, err error) {
	jsonParsed, err := gabs.ParseJSON(elem)
	if err != nil {
		return resElem, false, errors.Wrap(err, "parse json error")
	}

	isOk = false

	switch val := jsonParsed.Path(condition.path).Data().(type) {
	case string:
		isOk, err := checkString(val, condition)
		if err != nil {
			return resElem, false, errors.Wrapf(err, "error process path as string")
		}
		return jsonParsed.String(), isOk, nil
	case int64:
		isOk, err = checkInt64(val, condition)
		if err != nil {
			return resElem, false, errors.Wrapf(err, "error process path as int64")
		}
	case float64:
		isOk, err = checkFloat64(val, condition)
		if err != nil {
			return resElem, false, errors.Wrapf(err, "error process path as float64")
		}
	case nil:
		isOk, err = checkNil(condition)
		if err != nil {
			return resElem, false, errors.Wrapf(err, "error process path as nil")
		}
	default:
		return resElem, false, errors.Wrapf(ErrUnsupportedType, "unsupported type of val by path '%s'", condition.path)
	}

	return jsonParsed.String(), isOk, nil
}

func checkString(checkVal string, condition Condition) (bool, error) {
	switch condition.operator {
	case OpEq:
		return checkVal == condition.value, nil
	case OpNotEq:
		return checkVal != condition.value, nil
	default:
		return false, errors.Wrapf(ErrUnsupportedOperator, "passed %s", condition.operator.String())
	}
}

func checkInt64(checkVal int64, condition Condition) (bool, error) {
	conditionVal, err := strconv.ParseFloat(condition.value, 64)
	if err != nil {
		return false, errors.Wrapf(err, "fail to parse '%s' as number", condition.value)
	}

	switch condition.operator {
	case OpEq:
		return checkVal == int64(conditionVal), nil
	case OpNotEq:
		return checkVal != int64(conditionVal), nil
	case OpLt:
		return checkVal < int64(conditionVal), nil
	case OpLte:
		return checkVal <= int64(conditionVal), nil
	case OpGt:
		return checkVal > int64(conditionVal), nil
	case OpGte:
		return checkVal >= int64(conditionVal), nil
	default:
		return false, errors.Wrapf(ErrUnsupportedOperator, "passed %s", condition.operator.String())
	}
}

func checkFloat64(checkVal float64, condition Condition) (bool, error) {
	conditionVal, err := strconv.ParseFloat(condition.value, 64)
	if err != nil {
		return false, errors.Wrapf(err, "fail to parse '%s' as number", condition.value)
	}

	switch condition.operator {
	case OpEq:
		return checkVal == conditionVal, nil
	case OpNotEq:
		return checkVal != conditionVal, nil
	case OpLt:
		return checkVal < conditionVal, nil
	case OpLte:
		return checkVal <= conditionVal, nil
	case OpGt:
		return checkVal > conditionVal, nil
	case OpGte:
		return checkVal >= conditionVal, nil
	default:
		return false, errors.Wrapf(ErrUnsupportedOperator, "passed %s", condition.operator.String())
	}
}

func checkNil(condition Condition) (bool, error) {
	return false, nil
}
