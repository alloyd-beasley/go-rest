package worker

import "net/http"

func worker(jobs <-chan *http.Response, results chan<- *http.Response) {
	//do some work on http resonses 
}
