package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Bydate queries openFDA device events by date
func Bydate(w http.ResponseWriter, r *http.Request) {

	fromdate := r.URL.Query().Get("fromdate")
	todate := r.URL.Query().Get("to")
	
	requesturl := fmt.Sprintf("https://api.fda.gov/device/event.json?search=date_received:[%s+TO+%s]&limit=1", fromdate, todate)
	
	resp, err := http.Get(requesturl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
