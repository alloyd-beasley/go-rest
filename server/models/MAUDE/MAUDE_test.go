package MAUDE

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	// "github.com/krdo-93/go-rest.git/server/models"
)

func TestDecodeJsonToMaude(t *testing.T) {
	jsonMap := unmarshalHelper("maude_results.json", t)
	ri := jsonMap["results"].([]interface{})
	// rs := ri[0].(map[string]interface{})

	var records []MAUDE
	for i := range ri {
		val := ri[i].(map[string]interface{})
		fmt.Println(val["type_of_report"])

		record := MAUDE{
			EventLocation:        val["event_location"].(string),
			ReportToFda:          val["report_to_fda"].(string),
			EventType:            val["event_type"].(string),
			ReportNumber:         val["report_number"].(string),
			TypeOfReport:         val["type_of_report"].([]string),
			// ProductProblemFlag:   val["Product_problem_flag"].(string),
			// DateReceived:         val["date_received"].(string),
			// DateOfEvent:          val["date_of_event"].(string),
			// ReportDate:           val["report_date"].(string),
			// DateFacilityAware:    val["date_facility_aware"].(string),
			// Device:               val["device"].([]models.Device),
			// Patient:              val["patient"].([]models.Patient),
			// NumberDevicesInEvent: val["number_devices_in_event"].(string),
			// MdrText:              val["mdr_text"].([]models.MdrText),
			// ManufacturerName:     val["manufacturer_name"].(string),
		}

		records = append(records, record)
	}
}

func unmarshalHelper(file string, t *testing.T) map[string]interface{} {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()

	if err != nil {
		t.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Fatal(err)
	}

	var interfaceValue interface{}
	err = json.Unmarshal(byteValue, &interfaceValue)

	if err != nil {
		t.Fatal(err)
	}

	return interfaceValue.(map[string]interface{})
}
