package db

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/alloyd-beasley/go-rest.git/models"
)

//InsertReportFromFile inserts raw data from given file
func InsertReport(data []byte) error {
	db := Initialize()
	defer db.Close()

	response := &models.Response{}

	if err := response.Parse(data); err != nil {
		log.Printf("Error inserting device: %v, %v", err, err.Error())
		return err
	}

	reportStatment, _ := ioutil.ReadFile("./statements/insert_report.sql")
	deviceStatement, _ := ioutil.ReadFile("./statements/insert_device.sql")

	for _, v := range response.Results {

		var deviceID int

		deviceErr := db.QueryRow(string(deviceStatement),
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
		).Scan(&deviceID)

		if deviceErr != nil {
			log.Printf("Error inserting device: %v, %v", deviceErr, deviceErr.Error())
			return deviceErr
		}

		numberDevicesInEvent, _ := strconv.Atoi(v.Number_devices_in_event)

		_, reportErr := db.Exec(string(reportStatment),
			v.Event_location,
			v.Report_to_fda,
			v.Event_type,
			v.Report_number,
			v.Type_of_report[0],
			v.Product_problem_flag,
			v.Date_received,
			v.Date_of_event,
			v.Report_date,
			v.Date_facility_aware,
			numberDevicesInEvent,
			v.Manufacturer_name,
			deviceID,
			v.Mdr_text[0].Text_type_code,
			v.Mdr_text[0].Text,
		)

		if reportErr != nil {
			log.Printf("Error inserting report: %v. %v", reportErr, reportErr.Error())
			return reportErr
		}
	}

	return nil
}
