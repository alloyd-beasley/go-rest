package models

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseReportData(t *testing.T) {
	data, err := ioutil.ReadFile("../testdata/testRecord.json")

	if err != nil {
		log.Fatal("There was a problem reading your file", err)
	}

	response := Response{}
	_, pErr := response.Parse(data)

	if pErr != nil {
		t.Error("TestParseReportData: ", err)
	}
}
