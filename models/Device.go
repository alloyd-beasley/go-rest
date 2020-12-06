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
	DeviceName                  string      `json:"device_name"`
	MedicalSpecialtyDescription string      `json:"medical_specialty_description"`
	DeviceClass                 DeviceClass `json:"device_class"`
	RegulationNumber            string      `json:"regulation_number"`
}

type Device struct {
	ManufacturerDAddress1   string `json:"manufacturer_d_address_1"`
	ManufacturerDState      string `json:"manufacturer_d_state"`
	ManufacturerDPostalCode string `json:"manufacturer_d_Postal_code"`
	ManufacturerDCity       string `json:"manufacturer_d_City"`
	ManufacturerDCountry    string `json:"manufacturer_d_Country"`
	ManufacturerDName       string `json:"manufacturer_d_Name"`
	LotNumber               string `json:"lot_number"`
	ModelNumber             string `json:"model_number"`
	GenericName             string `json:"generic_name"`
	BrandName               string `json:"brand_name"`
}
