package shakespearean

import (
	"log"
	"fmt"
	"net/http"
	"strings"
	"path"
	"io/ioutil"
)

func HealthEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"Status":"UP"}`)
}


func PokemonRequest(w http.ResponseWriter, r *http.Request) {
	//name := r.URL.Path[len("/pokemon/"):]

	var pokemonDataRequest string

	urlPath := r.URL.Path

	log.Println("Got a pokemon request: ", urlPath)

	if strings.HasPrefix(urlPath, "/pokemon") {

		pokemonDataRequest = path.Base(urlPath)

		log.Println(pokemonDataRequest)

		body, val := callPokeApi(urlPath)

		if val < 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))			
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))		
	}

	//fmt.Fprintf(w, `{"Status":"OK"}`)
}

func callPokeApi(receivedPath string) (string, int) {

	callPath := "https://pokeapi.co/api/v2" + receivedPath

	log.Println("Calling ", callPath)

	req, err := http.NewRequest("GET", callPath, nil)

	if err != nil {
		return "", -1
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", -1
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", -1
	}

	log.Println(body)

	return string(body), 0

}
