package shakespearean

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

const incomingUrl string = "/pokemon"

type responseVals struct {
	Name        string `json:"name"`
	Translation string `json:"translation"`
}

func HealthEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"Status":"UP"}`)
}

func HandleIncoming(w http.ResponseWriter, r *http.Request, incomingMetrics *prometheus.CounterVec) {

	var pokemonDataRequest string
	var rspStatusCode int
	var rspBody string

	urlPath := r.URL.Path

	log.Println("Got a pokemon request: ", urlPath)

	if strings.HasPrefix(urlPath, incomingUrl) {

		pokemonDataRequest = path.Base(urlPath)

		log.Println(pokemonDataRequest)

		body, val := processIncoming(pokemonDataRequest)

		if val < 0 {
			rspStatusCode = http.StatusInternalServerError
			rspBody = "Internal error"
			incomingMetrics.WithLabelValues("500", "Internal Error").Inc()
		} else if val == 1 {
			rspStatusCode = http.StatusNotFound
			rspBody = "Not Found"
			incomingMetrics.WithLabelValues("404", "Not Found").Inc()
		} else {
			rspStatusCode = val
			rspBody = body
			incomingMetrics.WithLabelValues(strconv.Itoa(val), body).Inc()
		}

	} else {
		rspStatusCode = http.StatusNotFound
		rspBody = "Not Found"
		incomingMetrics.WithLabelValues("404", "Not Found").Inc()
	}

	rsp := responseVals{pokemonDataRequest, rspBody}

	jsonData, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rspStatusCode)
	val, err := w.Write(jsonData)

	if err != nil {
		log.Println("Could not send back response, will timeout at other end: ", err)
	}

	log.Println(val)
}
