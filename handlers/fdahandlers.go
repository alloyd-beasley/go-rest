package fdahandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alloyd-beasley/go-rest.git/models/MAUDE"
	"github.com/alloyd-beasley/go-rest.git/util"
)

const deviceEventURL = "https://api.fda.gov/device/event.json"

//GetLimit retrieves records by limit
func GetLimit(limit string) ([]MAUDE.MAUDEResults, error) {

	query := fmt.Sprintf("?limit=%s", limit)
	requestURL := fmt.Sprintf(deviceEventURL+"%s", query)
	resp, err := http.Get(requestURL)

	if err != nil {
		log.Println("Error making request to FDA API")
		return nil, fmt.Errorf("Something went wrong when making the request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error reading response body from FDA API")
		return nil, fmt.Errorf("Something went wrong when reading the response body: %v", err)
	}

	response := MAUDE.Response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, httperror.NewHTTPError(err, "Something went wrong while Unmarshalling response json", 400)
	}

	var results []MAUDE.MAUDEResults

	for _, v := range response.Results {
		r := MAUDE.MAUDEResults{
			EventLocation:        v.EventLocation,
			ReportToFda:          v.ReportToFda,
			EventType:            v.EventType,
			ReportNumber:         v.ReportNumber,
			TypeOfReport:         v.TypeOfReport,
			ProductProblemFlag:   v.ProductProblemFlag,
			DateReceived:         v.DateReceived,
			DateOfEvent:          v.DateOfEvent,
			ReportDate:           v.ReportDate,
			DateFacilityAware:    v.DateFacilityAware,
			Device:               v.Device,
			Patient:              v.Patient,
			NumberDevicesInEvent: v.NumberDevicesInEvent,
			MdrText:              v.MdrText,
			ManufacturerName:     v.ManufacturerName,
		}

		results = append(results, r)
	}

	return results, nil
}
