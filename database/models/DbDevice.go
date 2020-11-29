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
	device_name                   string
	medical_specialty_description string
	device_class                  DeviceClass
	regulation_number             string
}

type Device struct {
	manufacturer_d_address_1   string
	device_sequence_number     string
	manufacturer_d_state       string
	manufacturer_d_Postal_code string
	manufacturer_d_City        string
	lot_number                 string
	model_number               string
	device_report_productCode  string
	device_operator            string
	other_id_number            string
	generic_name               string
	manufacturer_d_Name        string
	manufacturer_d_Country     string
	brand_name                 string
	openfda                    DeviceOpenFda
}
