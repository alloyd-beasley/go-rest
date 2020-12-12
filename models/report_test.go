package models

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseReportData(t *testing.T) {
	data, err := ioutil.ReadFile("../db/testdata/testRecord.json")

	if err != nil {
		log.Fatal("There was a problem reading your file", err)
	}

	response := FDAResponse{}
	parsed, err := response.Parse(data)
	var want []Report

	if err != nil {
		t.Error("There was an error parsing response: ", err.Error())
	}

	log.Printf("This is the parsed report: %+v", parsed)

	assert.IsType(t, parsed, want)
}
