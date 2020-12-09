package db

import (
	"io/ioutil"
	"log"

	"github.com/alloyd-beasley/go-rest.git/models"
)

func parseTime(item string) string {
	layout := "2006-01-02"

	return layout
}

func insertReport() {
	db := Initialize()

	data, err := ioutil.ReadFile("../tools/data/testRecord.json")

	if err != nil {
		log.Fatalf("There was a problem reading your file: %v", err)
	}

	response := models.Response{}
	results, rErr := response.Parse(data)

	if rErr != nil {
		log.Fatalf("Parser Error: %v", rErr)
	}

	reportStatment, _ := ioutil.ReadFile("./statements/insert_report.sql")
	deviceStatement, _ := ioutil.ReadFile("./statements/insert_device.sql")
	mdrTextStatement, _ := ioutil.ReadFile("./statements/insert_mdr_text.sql")

	for _, v := range results {
		deviceID, dErr := db.Exec(string(deviceStatement),
			v.Device[0].Manufacturer_d_address_1,
			v.Device[0].Manufacturer_d_state,
			v.Device[0].Manufacturer_d_Postal_code,
			v.Device[0].Manufacturer_d_City,
			v.Device[0].Manufacturer_d_Country,
			v.Device[0].Manufacturer_d_Name,
			v.Device[0].Lot_number,
			v.Device[0].Model_number,
			v.Device[0].Generic_name,
			v.Device[0].Brand_name,
		)

		if dErr != nil {
			log.Fatalf("Error inserting device: %v", rErr)
		}

		textID, tErr := db.Exec(string(mdrTextStatement),
			v.Mdr_text[0].Text_type_code,
			v.Mdr_text[0].Text,
		)

		if tErr != nil {
			log.Fatalf("Error inserting text: %v", tErr)
		}

		_, rErr := db.Exec(string(reportStatment),
			v.Event_location,
			v.Report_to_fda,
			v.Event_type,
			v.Report_number,
			v.Type_of_report,
			v.Product_problem_flag,
			v.Date_received,
			v.Date_of_event,
			v.Report_date,
			v.Date_facility_aware,
			v.Number_devices_in_event,
			v.Manufacturer_name,
			deviceID,
			textID,
		)

		if rErr != nil {
			log.Fatalf("Error inserting report: %v", rErr)
		}
	}
}
