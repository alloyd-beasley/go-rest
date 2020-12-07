package models

//Device defines MAUDE report device section
type Device struct {
	manufacturer_d_address_1   string
	manufacturer_d_state       string
	manufacturer_d_Postal_code string
	manufacturer_d_City        string
	manufacturer_d_Country     string
	manufacturer_d_Name        string
	lot_number                 string
	model_number               string
	generic_name               string
	brand_name                 string
}
