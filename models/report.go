package models

import (
	"encoding/json"
	"fmt"
	"log"
)

//Report defines MAUDE report
type Report struct {
	Event_location          string
	Report_to_fda           string
	Event_type              string
	Report_number           string
	Type_of_report          []string
	Product_problem_flag    string
	Date_received           string
	Date_of_event           string
	Report_date             string
	Date_facility_aware     string
	Device                  []Device
	Number_devices_in_event string
	Mdr_text                []MdrText
	Manufacturer_name       string
}

type Response struct {
	Results []Report
}

func (r Response) Parse(data []byte) ([]Report, error) {

	var report []Report

	if err := json.Unmarshal(data, &r); err != nil {
		log.Println("There was problem parsing a report: ", err)
		return nil, fmt.Errorf("There was a problem parsing your report: %v", err)
	}

	for _, v := range r.Results {
		report = append(report, v)
	}

	fmt.Printf("%+v", report)

	return report, nil
}
