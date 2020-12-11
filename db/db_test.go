package db

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseDate(t *testing.T) {
	date := ParseDate("19930316")
	_, err := regexp.Match("([0-9]+)-([0-9]+)-([0-9])", []byte(date))
	if err != nil {
		t.Error(err)
	}

	t.Logf("This is the date: %v", date)
}

func TestInsertRawReport(t *testing.T) {
	data, _ := ioutil.ReadFile("./testData/testRecord.json")
	
	if err := InsertReport(data); err != nil {
		t.Error("Error inserting report from test: ", err)
	}
}
