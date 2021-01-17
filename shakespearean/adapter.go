package shakespearean

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/structs"
	"github.com/mtslzr/pokeapi-go"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

const shakespeareEndpoint string = "https://api.funtranslations.com/translate/shakespeare.json"

type shakespeareResponse struct {
	Contents struct {
		Translated string `json:"translated"`
	} `json:"contents"`
}

func processIncoming(name string) (string, int) {

	var id int
	var err error
	var flavourText string

	resource, _ := pokeapi.Resource("pokemon")

	for i := range resource.Results {
		if resource.Results[i].Name == name {

			idStr := path.Base(resource.Results[i].URL)

			id, err = strconv.Atoi(idStr)

			if err != nil {
				log.Println(err)
				return "", -1
			}

			log.Println("Found corresponding ID ", id, resource.Results[i].Name)

		}
	}

	p, _ := pokeapi.PokemonSpecies(name)

	m := structs.Map(p)

	for k, v := range m {

		if k == "FlavorTextEntries" {

			for k2, v2 := range v.([]interface{}) {

				if k2 == id {

					for k3, v3 := range v2.(map[string]interface{}) {

						if k3 == "FlavorText" {
							flavourText = v3.(string)
							log.Println("To Translate:", flavourText)
							return callTranslator(flavourText)
						}
					}
				}
			}
		}
	}

	return "", 1
}

func callTranslator(strToTranslate string) (string, int) {

	data := make(map[string]string)
	data["text"] = strToTranslate

	jsonString, err := json.Marshal(data)

	if err != nil {
		log.Println("Error with json creation: ", err)
		return "", -1
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", shakespeareEndpoint, bytes.NewBuffer(jsonString))

	if err != nil {
		log.Println("Error with request creation: ", err)
		return "", -1
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error with connection: ", err)
		return "", -1
	}
	defer resp.Body.Close()

	log.Println("\nResponse Status   ---->:", resp.Status)

	if resp.StatusCode == http.StatusOK {

		body, _ := ioutil.ReadAll(resp.Body)

		log.Println(string(body))

		shResp := shakespeareResponse{}

		err := json.Unmarshal(body, &shResp)

		if err != nil {
			log.Println("Error with response handling: ", err)
			return "", -1
		}

		log.Println("Got translated: ", shResp.Contents.Translated)

		return shResp.Contents.Translated, resp.StatusCode
	} else {
		return resp.Status, resp.StatusCode
	}
}
