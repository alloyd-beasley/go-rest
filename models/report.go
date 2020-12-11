package models

import (
	"encoding/json"
	"log"

	"github.com/alloyd-beasley/go-rest.git/common"
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

func (r *Response) Parse(data []byte) error {

	if err := json.Unmarshal(data, &r); err != nil {
		log.Printf("There was a problem parsing your report: %v, %v", err, err.Error())
		return err
	}

	for _, v := range r.Results {
		v.Date_received = common.ParseDate(v.Date_received)
		v.Date_of_event = common.ParseDate(v.Date_of_event)
		v.Report_date = common.ParseDate(v.Report_date)
		v.Date_facility_aware = common.ParseDate(v.Date_facility_aware)
	}

	return nil
}
