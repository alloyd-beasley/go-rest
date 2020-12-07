package models

import "github.com/alloyd-beasley/go-rest.git/models"

//Report defines MAUDE report
type Report struct {
	event_location          string
	report_to_fda           string
	event_type              string
	report_number           string
	type_of_report          []string
	product_problem_flag    string
	date_received           string
	date_of_event           string
	report_date             string
	date_facility_aware     string
	device                  []models.Device
	patient                 []models.Patient
	number_devices_in_event string
	mdr_text                []models.MdrText
	manufacturer_name       string
}

type Response struct {
	Results []Report
}
