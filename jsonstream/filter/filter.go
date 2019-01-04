package filter

import (
	"strconv"
	"strings"

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
func ProcessElem(condition Condition, elem []byte) (resElem []byte, isOk bool, err error) {
	resElem = elem
	jsonParsed, err := gabs.ParseJSON(elem)
	if err != nil {
		return resElem, false, errors.Wrap(err, "parse json error")
	}

	isOk, err = chechkValue(jsonParsed.Path(condition.path).Data(), condition)
	if err != nil {
		return resElem, false, errors.Wrap(err, "error check path")
	}

	return resElem, isOk, nil
}

func chechkValue(checkVal interface{}, condition Condition) (isOk bool, err error) {
	switch val := checkVal.(type) {
	case string:
		isOk, err = checkString(val, condition)
		if err != nil {
			return false, errors.Wrapf(err, "error process path as string")
		}
	case float64:
		isOk, err = checkFloat64(val, condition)
		if err != nil {
			return false, errors.Wrapf(err, "error process path as number")
		}
	case bool:
		isOk, err = checkBool(val, condition)
		if err != nil {
			return false, errors.Wrapf(err, "error process path as boolean")
		}
	case nil:
		isOk, err = checkNil(condition)
		if err != nil {
			return false, errors.Wrapf(err, "error process path as nil")
		}
	case []interface{}:
		for _, valElem := range val {
			elemIsOk, elemErr := chechkValue(valElem, condition)
			if elemErr != nil {
				return false, errors.Wrapf(elemErr, "error process elems of array in path as string")
			}

			if elemIsOk {
				return true, nil
			}
		}
	default:
		return false, errors.Wrapf(ErrUnsupportedType, "unsupported type of val by path '%s'", condition.path)
	}

	return isOk, err
}

func checkString(checkVal string, condition Condition) (bool, error) {
	switch condition.operator {
	case OpEq:
		return strings.Compare(checkVal, condition.value) == 0, nil
	case OpNotEq:
		return strings.Compare(checkVal, condition.value) != 0, nil
	case OpGt:
		return strings.Compare(checkVal, condition.value) > 0, nil
	case OpGte:
		return strings.Compare(checkVal, condition.value) >= 0, nil
	case OpLt:
		return strings.Compare(checkVal, condition.value) < 0, nil
	case OpLte:
		return strings.Compare(checkVal, condition.value) <= 0, nil
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

func checkBool(checkVal bool, condition Condition) (bool, error) {
	conditionVal, err := strconv.ParseBool(condition.value)
	if err != nil {
		return false, errors.Wrapf(err, "fail to parse '%s' as bool", condition.value)
	}

	switch condition.operator {
	case OpEq:
		return checkVal == conditionVal, nil
	case OpNotEq:
		return checkVal != conditionVal, nil
	default:
		return false, errors.Wrapf(ErrUnsupportedOperator, "passed %s", condition.operator.String())
	}
}
