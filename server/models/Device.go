package models

type DeviceClass string

const (
	I             DeviceClass = "Class I (low to moderate risk): general controls"
	II            DeviceClass = "Class II (moderate to high risk): general controls and special controls"
	III           DeviceClass = "Class III (high risk): general controls and Premarket Approval (PMA)"
	UN            DeviceClass = "Unclassified"
	F             DeviceClass = "HDE"
	NotClassified DeviceClass = "Not Classified"
)

type DeviceOpenFda struct {
	DeviceName                  string
	MedicalSpecialtyDescription string
	DeviceClass                 DeviceClass
	RegulationNumber            float32
}

type Device struct {
	ManufacturerDAddress1         string
	DeviceSequenceNumber          uint16
	ManufacturerDState            string
	ManufacturerDPostalCode       uint16
	ManufacturerDCity             string
	LotNumber                     string
	ModelNumber                   string
	DeviceReportProductCode       string
	DeviceOperator                string
	OtherIdNumber                 uint16
	GenericName                   string
	ManufacturerDName             string
	ManufacturerDCountry          string
	BrandName                     string
	OpenFda                       DeviceOpenFda
	DeviceAgeText                 string
	DeviceEvaluatedByManufacturer string
	CatalogNumber                 string
	Baseline510kExemptFlag        string
	ImplantFlag                   string
	DateRemovedFlag               string
}
