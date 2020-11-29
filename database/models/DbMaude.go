package models

type textTypeCode string

const (
	DescriptionOfEvent              textTypeCode = "Description of Event or Problem"
	ManufacturerEvaluationSummary   textTypeCode = "Manufacturer Evaluation Summary"
	AdditionalManufacturerNarrative textTypeCode = "Additional Manufacturer Narrative"
)

type MdrText struct {
	TextTypeCode string `json:"text_type_code"`
	Text         string `json:"text"`
}

type MAUDE struct {
	event_location          string
	report_to_fda           string
	event_type              string
	report_number           string
	type_of_report          []string
	product_problem_flag    string
	date_received           string
	date_of_event           string //estimate of first onset of event
	report_date             string //date report was forwarded to manu. or FDA
	date_facility_aware     string
	device                  []Device
	patient                 []Patient
	number_devices_in_event string
	mdr_text                []MdrText
	manufacturer_name       string
}

type MAUDEResults struct {
	Results []MAUDE
}
