package filter_test

import (
	"testing"

	"github.com/shnellpavel/json-stream/jsonstream/filter"
	"github.com/stretchr/testify/assert"
)

func TestNewConditionFromStr_Positive(t *testing.T) {
	cases := []struct {
		name             string
		inputExpr        string
		expectedPath     string
		expectedOperator filter.Operator
		expectedValue    string
	}{
		{
			name:             "Simple equation",
			inputExpr:        "attr = value",
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Nested path",
			inputExpr:        "attr.subAttr.sub-sub_attR = value",
			expectedPath:     "attr.subAttr.sub-sub_attR",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Nested path with multi word key",
			inputExpr:        "attr.'sub attr'.sub-sub_attR = value",
			expectedPath:     "attr.'sub attr'.sub-sub_attR",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Nested path with unicode key",
			inputExpr:        "attr.'ключ #1'.sub-sub_attR = value",
			expectedPath:     "attr.'ключ #1'.sub-sub_attR",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name: "Many spaces",
			inputExpr: "attr     = \t	value",
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Empty value",
			inputExpr:        "attr =",
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "",
		},
		{
			name:             "Array attr",
			inputExpr:        "attr[0] = value",
			expectedPath:     "attr[0]",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Array attr nested",
			inputExpr:        "x.attr[0].y = value",
			expectedPath:     "x.attr[0].y",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Multi words value single quotes",
			inputExpr:        "attr = 'value.part1 value_part2 value@part3'",
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "value.part1 value_part2 value@part3",
		},
		{
			name:             "Multi words value double quotes",
			inputExpr:        `attr = "value.part1, value_part2 value@part3"`,
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "value.part1, value_part2 value@part3",
		},
		{
			name:             "Operator: equal",
			inputExpr:        "attr = value",
			expectedPath:     "attr",
			expectedOperator: filter.OpEq,
			expectedValue:    "value",
		},
		{
			name:             "Operator: not equal",
			inputExpr:        "attr != value",
			expectedPath:     "attr",
			expectedOperator: filter.OpNotEq,
			expectedValue:    "value",
		},
		{
			name:             "Operator: less than",
			inputExpr:        "attr < value",
			expectedPath:     "attr",
			expectedOperator: filter.OpLt,
			expectedValue:    "value",
		},
		{
			name:             "Operator: less than or equal",
			inputExpr:        "attr <= value",
			expectedPath:     "attr",
			expectedOperator: filter.OpLte,
			expectedValue:    "value",
		},
		{
			name:             "Operator: greater than",
			inputExpr:        "attr > value",
			expectedPath:     "attr",
			expectedOperator: filter.OpGt,
			expectedValue:    "value",
		},
		{
			name:             "Operator: greater than or equal",
			inputExpr:        "attr >= value",
			expectedPath:     "attr",
			expectedOperator: filter.OpGte,
			expectedValue:    "value",
		},
		{
			name:             "Operator: like",
			inputExpr:        "attr ~ value",
			expectedPath:     "attr",
			expectedOperator: filter.OpLike,
			expectedValue:    "value",
		},
		{
			name:             "Operator: not like",
			inputExpr:        "attr !~ value",
			expectedPath:     "attr",
			expectedOperator: filter.OpNotLike,
			expectedValue:    "value",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			condition, err := filter.NewConditionFromStr(testCase.inputExpr)
			assert.NoError(t, err)
			if !assert.NotNil(t, condition) {
				assert.Fail(t, "condition is nil")
				return
			}

			assert.Equal(t, testCase.expectedPath, condition.Path(), "path part hasn't expected value")
			assert.Equal(t, testCase.expectedOperator.String(), condition.Operator().String(), "operator hasn't expected value")
			assert.Equal(t, testCase.expectedValue, condition.Value(), "value part hasn't expected value")

		})
	}
}

func TestNewConditionFromStr_Negative(t *testing.T) {
	cases := []struct {
		name      string
		inputExpr string
	}{
		{
			name:      "Empty expression",
			inputExpr: "",
		},
		{
			name:      "Any string",
			inputExpr: "asdlkaj sahd kajsdh kajhds kjahdkj ahsdkjljaldj asd",
		},
		{
			name:      "Empty path",
			inputExpr: "= value",
		},
		{
			name:      "Empty operator",
			inputExpr: "attr value",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			condition, err := filter.NewConditionFromStr(testCase.inputExpr)
			assert.Nil(t, condition)
			assert.Error(t, err)
		})
	}
}
