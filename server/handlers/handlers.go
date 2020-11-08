package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Test() {
	fmt.Println("this is a test handler")
}

func Placeholder(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(body))
}
