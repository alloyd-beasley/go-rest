package db

import (
	"io/ioutil"
	"testing"
)

func TestInsertRawReport(t *testing.T) {
	data, _ := ioutil.ReadFile("./testData/testRecord.json")

	if err := InsertReport(data); err != nil {
		t.Error("Error inserting report from test: ", err)
	}
}
