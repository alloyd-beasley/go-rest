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
func ParseDate(item string) (string, error) {
	//MAUDE dates look like: "YYMMDD"

	if len(item) < 1 {
		log.Print("Empty string passed to parser ParseTime")
		return "", fmt.Errorf("Tried to parse empty string")
	}

	itemSlice := strings.Split(item, "")

	year := strings.Join(itemSlice[:4], "")
	month := strings.Join(itemSlice[4:6], "")
	day := strings.Join(itemSlice[6:8], "")

	date := fmt.Sprintf("%v-%v-%v", year, month, day)

	return date, nil
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

		_, dateReceived := ParseDate(v.Date_received)
		_, dateOfEvent := ParseDate(v.Date_of_event)
		_, reportDate := ParseDate(v.Report_date)
		_, dateFacilityAware := ParseDate(v.Date_facility_aware)
		_, numberDevicesInEvent := strconv.Atoi(v.Number_devices_in_event)

		_, rErr := db.Exec(string(reportStatment),
			v.Event_location,
			v.Report_to_fda,
			v.Event_type,
			v.Report_number,
			v.Type_of_report,
			v.Product_problem_flag,
			dateReceived,
			dateOfEvent,
			reportDate,
			dateFacilityAware,
			numberDevicesInEvent,
			v.Manufacturer_name,
			deviceID,
			textID,
		)

		if rErr != nil {
			log.Fatalf("Error inserting report: %v", rErr)
		}
	}
}
