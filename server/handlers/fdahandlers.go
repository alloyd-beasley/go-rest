package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/krdo-93/go-rest.git/server/models/MAUDE"
	"github.com/krdo-93/go-rest.git/server/util/httperror"
)

//GetLimit retrieves records by limit
func GetLimit(limit string) ([]MAUDE.MAUDEResults, error) {

	requestURL := fmt.Sprintf("https://api.fda.gov/device/event.json?limit=%s", limit)
	resp, err := http.Get(requestURL)

	if err != nil {
		log.Println("Error making request to FDA API")
		return nil, fmt.Errorf("Something went wrong when making the request %v: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error reading response body from FDA API")
		return nil, fmt.Errorf("Something went wrong when reading the response body%v: ", err)
	}

	response := MAUDE.Response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, httperror.NewHTTPError(err, "Something went wrong while Unmarshaling response json", 400)
	}

	var results []MAUDE.MAUDEResults

	for _, v := range response.Results {
		r := MAUDE.MAUDEResults{
			v.EventLocation,
			v.ReportToFda,
			v.EventType,
			v.ReportNumber,
			v.TypeOfReport,
			v.ProductProblemFlag,
			v.DateReceived,
			v.DateOfEvent,
			v.ReportDate,
			v.DateFacilityAware,
			v.Device,
			v.Patient,
			v.NumberDevicesInEvent,
			v.MdrText,
			v.ManufacturerName,
		}

		results = append(results, r)
	}

	return results, nil
}
