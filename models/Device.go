package models

//Device defines MAUDE report device section
type Device struct {
	Manufacturer_d_address_1   string
	Manufacturer_d_state       string
	Manufacturer_d_Postal_code string
	Manufacturer_d_City        string
	Manufacturer_d_Country     string
	Manufacturer_d_Name        string
	Lot_number                 string
	Model_number               string
	Generic_name               string
	Brand_name                 string
}
