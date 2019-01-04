package filter_test

import (
	"testing"

	"github.com/shnellpavel/json-stream/jsonstream/filter"
	"github.com/stretchr/testify/assert"
)

func TestProcessElem_TypesAndOperators(t *testing.T) {
	type args struct {
		condition filter.Condition
		elem      []byte
	}

	buildCondition := func(expr string) filter.Condition {
		res, _ := filter.NewConditionFromStr(expr)
		return *res
	}

	cases := []struct {
		name         string
		args         args
		expectedIsOk bool
	}{
		{
			name: "Strings. Equal. Ok",
			args: args{
				condition: buildCondition("attr = value"),
				elem:      []byte(`{"attr": "value"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Equal. Not ok",
			args: args{
				condition: buildCondition("attr = value"),
				elem:      []byte(`{"attr": "value1"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Not equal. Ok",
			args: args{
				condition: buildCondition("attr != value"),
				elem:      []byte(`{"attr": "value1"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Not equal. Not ok",
			args: args{
				condition: buildCondition("attr != value"),
				elem:      []byte(`{"attr": "value"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Greater than. Ok",
			args: args{
				condition: buildCondition("attr > a"),
				elem:      []byte(`{"attr": "bar"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Greater than. Not ok 1",
			args: args{
				condition: buildCondition("attr > b"),
				elem:      []byte(`{"attr": "aka"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Greater than. Not ok 2",
			args: args{
				condition: buildCondition("attr > b"),
				elem:      []byte(`{"attr": "b"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Greater than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr >= a"),
				elem:      []byte(`{"attr": "a"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Greater than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr >= a"),
				elem:      []byte(`{"attr": "bar"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Greater than or equal. Not ok",
			args: args{
				condition: buildCondition("attr >= b"),
				elem:      []byte(`{"attr": "aka"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Less than. Ok",
			args: args{
				condition: buildCondition("attr < b"),
				elem:      []byte(`{"attr": "aka"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Less than. Not ok 1",
			args: args{
				condition: buildCondition("attr < b"),
				elem:      []byte(`{"attr": "b"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Less than. Not ok 2",
			args: args{
				condition: buildCondition("attr < b"),
				elem:      []byte(`{"attr": "bar"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Strings. Less than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr <= b"),
				elem:      []byte(`{"attr": "b"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Less than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr <= b"),
				elem:      []byte(`{"attr": "aka"}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Strings. Less than or equal. Not ok",
			args: args{
				condition: buildCondition("attr <= bar"),
				elem:      []byte(`{"attr": "baz"}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Equal. Ok",
			args: args{
				condition: buildCondition("attr = 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Equal. Not ok",
			args: args{
				condition: buildCondition("attr = 26"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Not equal. Ok",
			args: args{
				condition: buildCondition("attr != 25"),
				elem:      []byte(`{"attr": -25}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Not equal. Not ok",
			args: args{
				condition: buildCondition("attr != 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Greater than. Ok",
			args: args{
				condition: buildCondition("attr > 25"),
				elem:      []byte(`{"attr": 26}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Greater than. Not ok 1",
			args: args{
				condition: buildCondition("attr > 25"),
				elem:      []byte(`{"attr": 2}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Greater than. Not ok 2",
			args: args{
				condition: buildCondition("attr > 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Greater than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr >= 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Greater than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr >= 25"),
				elem:      []byte(`{"attr": 26}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Greater than or equal. Not ok",
			args: args{
				condition: buildCondition("attr >= 25"),
				elem:      []byte(`{"attr": 2}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Less than. Ok",
			args: args{
				condition: buildCondition("attr < 25"),
				elem:      []byte(`{"attr": 24}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Less than. Not ok 1",
			args: args{
				condition: buildCondition("attr < 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Less than. Not ok 2",
			args: args{
				condition: buildCondition("attr < 25"),
				elem:      []byte(`{"attr": 26}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Integers. Less than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr <= 25"),
				elem:      []byte(`{"attr": 25}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Less than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr <= 25"),
				elem:      []byte(`{"attr": 24}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Integers. Less than or equal. Not ok",
			args: args{
				condition: buildCondition("attr <= 25"),
				elem:      []byte(`{"attr": 26}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Equal. Ok",
			args: args{
				condition: buildCondition("attr = 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Equal. Not ok",
			args: args{
				condition: buildCondition("attr = 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Not equal. Ok",
			args: args{
				condition: buildCondition("attr != 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Not equal. Not ok",
			args: args{
				condition: buildCondition("attr != 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Greater than. Ok",
			args: args{
				condition: buildCondition("attr > 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Greater than. Not ok 1",
			args: args{
				condition: buildCondition("attr > 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Greater than. Not ok 2",
			args: args{
				condition: buildCondition("attr > 25.5"),
				elem:      []byte(`{"attr": 25.4}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Greater than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr >= 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Greater than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr >= 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Greater than or equal. Not ok",
			args: args{
				condition: buildCondition("attr >= 25.5"),
				elem:      []byte(`{"attr": 25.4}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Less than. Ok",
			args: args{
				condition: buildCondition("attr < 25.5"),
				elem:      []byte(`{"attr": 25.4}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Less than. Not ok 1",
			args: args{
				condition: buildCondition("attr < 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Less than. Not ok 2",
			args: args{
				condition: buildCondition("attr < 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Floats. Less than or equal. Ok 1",
			args: args{
				condition: buildCondition("attr <= 25.5"),
				elem:      []byte(`{"attr": 25.5}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Less than or equal. Ok 2",
			args: args{
				condition: buildCondition("attr <= 25.5"),
				elem:      []byte(`{"attr": -25.5}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Floats. Less than or equal. Not ok",
			args: args{
				condition: buildCondition("attr <= 25.5"),
				elem:      []byte(`{"attr": 25.6}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Booleans. Equal. Ok",
			args: args{
				condition: buildCondition("attr = true"),
				elem:      []byte(`{"attr": true}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Booleans. Equal. Not ok",
			args: args{
				condition: buildCondition("attr = true"),
				elem:      []byte(`{"attr": false}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Booleans. Not equal. Ok",
			args: args{
				condition: buildCondition("attr != true"),
				elem:      []byte(`{"attr": false}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Booleans. Not equal. Not ok",
			args: args{
				condition: buildCondition("attr != true"),
				elem:      []byte(`{"attr": true}`),
			},
			expectedIsOk: false,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, actualIsOk, err := filter.ProcessElem(testCase.args.condition, testCase.args.elem)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedIsOk, actualIsOk)
		})
	}
}

func TestProcessElem_DiffferentPathsAndValues(t *testing.T) {
	type args struct {
		condition filter.Condition
		elem      []byte
	}

	buildCondition := func(expr string) filter.Condition {
		res, _ := filter.NewConditionFromStr(expr)
		return *res
	}

	cases := []struct {
		name         string
		args         args
		expectedIsOk bool
	}{
		{
			name: "Only direct objects. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = value"),
				elem:      []byte(`{"attr1": {"attr2": {"attr3": "value"}}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Only direct objects. Not equal. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = value"),
				elem:      []byte(`{"attr1": {"attr2": {"attr3": "value1"}}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Only direct objects. Absent attribute. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = value"),
				elem:      []byte(`{"attr1": {"attr2": {}}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Only direct objects. Unexpected type attribute. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = value"),
				elem:      []byte(`{"attr1": {"attr2": 3}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Unicode multiword attr. Ok",
			args: args{
				condition: buildCondition("attr1.Некий атрибут.attr3 = value"),
				elem:      []byte(`{"attr1": {"Некий атрибут": {"attr3": "value"}}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Unicode multiword attr. Not equal. Not Ok",
			args: args{
				condition: buildCondition("attr1.Некий атрибут.attr3 = value"),
				elem:      []byte(`{"attr1": {"Некий атрибут": {"attr3": "value1"}}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Unicode multiword value. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = 'Некое значение'"),
				elem:      []byte(`{"attr1": {"attr2": {"attr3": "Некое значение"}}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Unicode multiword value. Not equal. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.attr3 = 'Некое значение'"),
				elem:      []byte(`{"attr1": {"attr2": {"attr3": "Другое значение"}}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Array elems. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.id = id2"),
				elem:      []byte(`{"attr1": {"attr2": [{"id": "id1"}, {"id": "id2"}, {"id": "id3"}]}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Array elems nested. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.ids = id2"),
				elem:      []byte(`{"attr1": {"attr2": [{"ids": ["id1", "id2", "id3"]}]}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Array elems. Not equal. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.id = id0"),
				elem:      []byte(`{"attr1": {"attr2": [{"id": "id1"}, {"id": "id2"}, {"id": "id3"}]}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Array elems nested. Not equal. Not Ok",
			args: args{
				condition: buildCondition("attr1.attr2.ids = id0"),
				elem:      []byte(`{"attr1": {"attr2": [{"ids": ["id1", "id2", "id3"]}]}}`),
			},
			expectedIsOk: false,
		},
		{
			name: "Array elems different types 1. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.id = 25"),
				elem:      []byte(`{"attr1": {"attr2": [{"id": "25"}, {"id": 26}, {"id": 25.050}]}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Array elems different types 1. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.id = 26"),
				elem:      []byte(`{"attr1": {"attr2": [{"id": "25"}, {"id": 26}, {"id": 25.050}]}}`),
			},
			expectedIsOk: true,
		},
		{
			name: "Array elems different types 1. Ok",
			args: args{
				condition: buildCondition("attr1.attr2.id = 25.050"),
				elem:      []byte(`{"attr1": {"attr2": [{"id": "25"}, {"id": 26}, {"id": 25.050}]}}`),
			},
			expectedIsOk: true,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			_, actualIsOk, err := filter.ProcessElem(testCase.args.condition, testCase.args.elem)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedIsOk, actualIsOk)
		})
	}
}
