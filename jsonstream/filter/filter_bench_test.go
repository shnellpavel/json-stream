package filter_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/shnellpavel/json-stream/jsonstream/filter"
)

func getFileContent(b *testing.B, path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		b.Fatalf("Fail to load file by path '%s'", path)
		b.FailNow()
	}

	res, err := ioutil.ReadAll(file)
	if err != nil {
		b.Fatalf("Fail to read file by path '%s'", path)
		b.FailNow()
	}

	return res
}

func BenchmarkProcessElem_200B_FirstLevel(b *testing.B) {

	testElem := getFileContent(b, "./test/200b.json")
	testCondition, _ := filter.NewConditionFromStr("name = John")

	for n := 0; n < b.N; n++ {
		_, isOk, err := filter.ProcessElem(*testCondition, testElem)
		if !isOk || err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkProcessElem_200B_NestedField(b *testing.B) {
	testElem := getFileContent(b, "./test/200b.json")
	testCondition, _ := filter.NewConditionFromStr("children.age > 5")

	for n := 0; n < b.N; n++ {
		_, isOk, err := filter.ProcessElem(*testCondition, testElem)
		if !isOk || err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkProcessElem_16KB_FirstLevel(b *testing.B) {

	testElem := getFileContent(b, "./test/16Kb.json")
	testCondition, _ := filter.NewConditionFromStr("mix = use")

	for n := 0; n < b.N; n++ {
		_, isOk, err := filter.ProcessElem(*testCondition, testElem)
		if !isOk || err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkProcessElem_16KB_NestedField(b *testing.B) {
	testElem := getFileContent(b, "./test/16Kb.json")
	testCondition, _ := filter.NewConditionFromStr("coast.construction.specific >= 1219861845")

	for n := 0; n < b.N; n++ {
		_, isOk, err := filter.ProcessElem(*testCondition, testElem)
		if !isOk || err != nil {
			b.FailNow()
		}
	}
}
