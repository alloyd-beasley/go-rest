package MAUDE

import (
	"github.com/krdo-93/go-rest.git/server/models"
)

type sourceType string

const (
	Other                 sourceType = "Other"
	Foreign               sourceType = "Foreign"
	Study                 sourceType = "Study"
	Literature            sourceType = "Literature"
	Consumer              sourceType = "Consumer"
	HealthProfessional    sourceType = "Health Professional"
	UserFacility          sourceType = "User facility"
	CompanyRepresentation sourceType = "Company representation"
	Distributor           sourceType = "Distributor"
	Unknown               sourceType = "Unknown"
	InvalidData           sourceType = "Invalid data"
)

type productProblemFlag uint8

const (
	Y productProblemFlag = 1 //about defect or malfunction
	N productProblemFlag = 0 //not about defect or malfunction
)

type typeOfReport string

const (
	InitialSubmission         typeOfReport = "Initial report of an event"
	Followup                  typeOfReport = "Additional or corrected information"
	ExtraCopyReceived         typeOfReport = "Documentation forthcoming"
	OtherInformationSubmitted typeOfReport = "Documentation forthcoming"
)

type MAUDEResponse struct {
	EventLocation        string           `json:"event_location"`
	ReportToFda          string           `json:"report_to_fda"`
	EventType            string           `json:"event_type"`
	ReportNumber         string           `json:"report_number"`
	TypeOfReport         []string         `json:"type_of_report"`
	ProductProblemFlag   string           `json:"product_problem_flag"`
	DateReceived         string           `json:"date_received"`
	DateOfEvent          string           `json:"date_of_event"` //estimate of first onset of event
	ReportDate           string           `json:"report_date"`   //date report was forwarded to manu. or FDA
	DateFacilityAware    string           `json:"date_facility_aware"`
	Device               []models.Device  `json:"device"`
	Patient              []models.Patient `json:"patient"`
	NumberDevicesInEvent string           `json:"number_devices_in_event"`
	MdrText              []models.MdrText `json:"mdr_text"`
	ManufacturerName     string           `json:"manufacturer_name"`
}

type MAUDEResults struct {
	EventLocation        string           
	ReportToFda          string           
	EventType            string           
	ReportNumber         string           
	TypeOfReport         []string         
	ProductProblemFlag   string           
	DateReceived         string           
	DateOfEvent          string           
	ReportDate           string           
	DateFacilityAware    string           
	Device               []models.Device  
	Patient              []models.Patient 
	NumberDevicesInEvent string           
	MdrText              []models.MdrText 
	ManufacturerName     string           
}

type Response struct {
	Results []MAUDEResponse
}



