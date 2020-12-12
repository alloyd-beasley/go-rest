package models

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/alloyd-beasley/go-rest.git/common"
)

//Report defines MAUDE report from FDA API
type FDAReport struct {
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

//Response defines MAUDE response from FDA API includeing results field for unmarshalling
type FDAResponse struct {
	Results []FDAReport
}

//Report defines MAUDE report with type conversions for certain fields
type Report struct {
	Event_location          string
	Report_to_fda           bool
	Event_type              string
	Report_number           string
	Type_of_report          string
	Product_problem_flag    string
	Date_received           string
	Date_of_event           string
	Report_date             string
	Date_facility_aware     string
	Device                  []Device
	Number_devices_in_event int
	Mdr_text                []MdrText
	Manufacturer_name       string
}

func (fr FDAResponse) Parse(data []byte) ([]Report, error) {

	if err := json.Unmarshal(data, &fr); err != nil {
		log.Printf("There was a problem parsing your report: %v, %v", err, err.Error())
		return nil, err
	}

	parsed := []Report{}

	for _, v := range fr.Results {
		v.Date_received = common.ParseDate(v.Date_received)
		v.Date_of_event = common.ParseDate(v.Date_of_event)
		v.Report_date = common.ParseDate(v.Report_date)
		v.Date_facility_aware = common.ParseDate(v.Date_facility_aware)

		reportToFda := false

		if v.Report_to_fda == "Y" {
			reportToFda = true
		}

		numberDevicesInEvent, err := strconv.Atoi(v.Number_devices_in_event)

		if err != nil {
			log.Print("Error converting Number_devices_in_event to int: ", err)
		}

		r := Report{
			Event_location:          v.Event_location,
			Report_to_fda:           reportToFda,
			Event_type:              v.Event_type,
			Report_number:           v.Report_number,
			Type_of_report:          v.Type_of_report[0],
			Product_problem_flag:    v.Product_problem_flag,
			Date_received:           common.ParseDate(v.Date_received),
			Date_of_event:           common.ParseDate(v.Date_of_event),
			Report_date:             common.ParseDate(v.Report_date),
			Date_facility_aware:     common.ParseDate(v.Date_facility_aware),
			Device:                  v.Device,
			Number_devices_in_event: numberDevicesInEvent,
			Mdr_text:                v.Mdr_text,
			Manufacturer_name:       v.Manufacturer_name,
		}

		parsed = append(parsed, r)
	}

	return parsed, nil
}
