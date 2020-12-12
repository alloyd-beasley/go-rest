package db

import (
	"io/ioutil"
	"testing"
)

var db DB

func TestMain(t *testing.T) {
	db = Initialize()
}

func TestInsertRawReport(t *testing.T) {
	data, _ := ioutil.ReadFile("./testData/testRecord.json")

	if err := db.InsertReport(data); err != nil {
		t.Error("Error inserting report from test: ", err)
	}
}

func TestGetAllReports(t *testing.T) {
	if err := db.GetAllReports(); err != nil {
		t.Error("Error selecting all reports: ", err)
	}
}

func TestGetReportById(t *testing.T) {
	id := 3
	if err := db.GetReportById(); err != nil {
		t.Error("Error selectingreports by id: ", id, err)
	}
}

func TestGetReportByDate(t *testing.T) {
	date := "1992-03-10"
	if err := db.GetReportByDate(); err != nil {
		t.Error("Error selectingreports by date: ", date, err)
	}
}
