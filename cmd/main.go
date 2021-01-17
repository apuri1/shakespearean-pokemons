package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"shakespeareanpokemons/shakespearean"
	"time"
)

var (
	incomingMetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "shakespeare_requests",
			Help: "Count of all requests",
		}, []string{"status_code", "info"})
)

func main() {

	log.Println("Starting server..")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shakespearean.HandleIncoming(w, r, incomingMetrics)
	})

	mux.HandleFunc("/health", shakespearean.HealthEndPoint)
	mux.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(incomingMetrics)

	s := &http.Server{
		Addr:         ":5000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())

}
