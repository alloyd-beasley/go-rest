package logic

import "time"

type EventType uint8

const (
	Y EventType = 1
	N EventType = 0
)

type DeviceClass string

const (
	I             DeviceClass = "Class I (low to moderate risk): general controls"
	II            DeviceClass = "Class II (moderate to high risk): general controls and special controls"
	III           DeviceClass = "Class III (high risk): general controls and Premarket Approval (PMA)"
	UN            DeviceClass = "Unclassified"
	F             DeviceClass = "HDE"
	NotClassified DeviceClass = "Not Classified"
)

type MAUDE struct {
	EventlLocation             string
	ReporToFda                 bool
	EvenType                   EventType
	ManufacturerContactAddress string
	DateOfEvent                time.Time
	DeviceDateOfManufacture    time.Time
}

type DeviceOpenFda struct {
	DeviceName         string
	RegistrationNumber int
	FeiNumber          int
	device_class       DeviceClass
}

type Device struct {
	ManufacturerDAddress1   string
	device_sequence_number  uint16
	ManufacturerDState      string
	ManufacturerDPostalCode uint16
	ManufacturerDCity       string
	LotNumber               string
	ModelNumber             string
	DateReceived            time.Time
	DeviceReportPorducCode  string
	DeviceOperator          string
	OtherIdNumber           uint16
	GenericName             string
	ManufacturerDName       string
	ManufacturerDCountry    string
	BrandName               string
}
