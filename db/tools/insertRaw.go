package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alloyd-beasley/go-rest.git/db/models"
)

func parseTime(item string) string {
	layout := "2006-01-02"

	return layout
}

func ParseRawData(data []byte) models.Report {
	var report models.Report

	if err := json.Unmarshal(data, &report); err != nil {
		log.Fatalf("There was a problem setting data to map %v", err)
	}
	
	return report
}

func InsertRawData() {
	data, err := ioutil.ReadFile("./data/testRecord.json")

	if err != nil {
		log.Fatal("There was a problem reading your file", err)
	}

	response := models.Response{}
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatalf("There was a problem setting data to map %v", err)
	}

	fmt.Println("report \n", response)
}
