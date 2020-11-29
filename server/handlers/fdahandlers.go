package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/krdo-93/go-rest.git/server/models/MAUDE"
)

func GetLimit(w http.ResponseWriter, r *http.Request) {

	limit := r.URL.Query().Get("limit")
	requestUrl := fmt.Sprintf("https://api.fda.gov/device/event.json?limit=%s", limit)
	resp, err := http.Get(requestUrl)

	if err != nil {
		fmt.Errorf("Something went wrong %v: ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Errorf("Something went wrong %v: ", err)
	}

	response := MAUDE.Response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Errorf("Something went wrong %v: ", err)
	}

	var results []MAUDE.MAUDEResults

	for i, v := range response.Results {
		fmt.Printf("i %v", string(i))
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
	output, err := json.Marshal(results)

	if err != nil {
		fmt.Errorf("Something went wrong %v: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
