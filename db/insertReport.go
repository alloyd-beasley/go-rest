package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/alloyd-beasley/go-rest.git/models"
)

//ParseTime parses MAUDE dates to "YY-MM-DD"
func ParseDate(item string) string {
	//MAUDE dates look like: "YYMMDD"

	if len(item) < 1 {
		log.Fatal("Tried to parse empty string")
	}

	itemSlice := strings.Split(item, "")

	year := strings.Join(itemSlice[:4], "")
	month := strings.Join(itemSlice[4:6], "")
	day := strings.Join(itemSlice[6:8], "")

	date := fmt.Sprintf("%v-%v-%v", year, month, day)

	return date
}

//InsertReportFromFile inserts raw data from given file
func InsertReport(data []byte) error {
	db := Initialize()
	defer db.Close()

	response := models.Response{}
	results := response.Parse(data)

	reportStatment, _ := ioutil.ReadFile("./statements/insert_report.sql")
	deviceStatement, _ := ioutil.ReadFile("./statements/insert_device.sql")

	for _, v := range results {

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
			ParseDate(v.Date_received),
			ParseDate(v.Date_of_event),
			ParseDate(v.Report_date),
			ParseDate(v.Date_facility_aware),
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
