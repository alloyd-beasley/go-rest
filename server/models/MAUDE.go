package models

import "time"

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

type MAUDE struct {
	EventLocation           string
	ReportToFda             bool
	DeviceDateOfManufacture time.Time
	EventType               string
	ReportNumber            string
	TypeOfReport            []typeOfReport
	ProductProblemFlag      productProblemFlag
	DateReceived            time.Time
	DateOfEvent             time.Time //estimate of first onset of event
	ReportDate              time.Time //date report was forwarded to manu. or FDA
	DateFacilityAware       time.Time
	Device                  []Device
	Patient                 []Patient
	NumberDevicesInEvent    uint
	MdrText                 []MdrText
	ManufacturerName        string
	SourceType              []sourceType
}
