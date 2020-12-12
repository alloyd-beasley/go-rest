package fdahandlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alloyd-beasley/go-rest.git/httperror"
	"github.com/alloyd-beasley/go-rest.git/models"
)

const deviceEventURL = "https://api.fda.gov/device/event.json"

//GetLimit retrieves records by limit
func GetLimit(limit string) ([]models.Report, error) {

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

	model := &models.FDAResponse{}
	response, err := model.Parse(body)

	if err != nil {
		return nil, httperror.NewHTTPError(err, "Something went wrong while Unmarshalling response json", 400)
	}

	return response, nil
}
