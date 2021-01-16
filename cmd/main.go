package main

import (
	"log"
	"net/http"
    "shakespeareanpokemons/shakespearean"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//func usage() {
//	log.Println("Args..")
//}

func main() {

	log.Println("Starting server..")

	mux := http.NewServeMux()

	mux.HandleFunc("/", shakespearean.PokemonRequest)

	http.HandleFunc("/health", shakespearean.HealthEndPoint)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":5000", mux)

}
