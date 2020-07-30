package rickAndMortyApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func GetLocationData(character Character) ([]LocationData, error) {
	res, err := http.Get(fmt.Sprintf("%slocation/?name=%s", baseUrl, url.QueryEscape(character.Location.Name)))
	log.Printf(fmt.Sprintf("%s/location/?name=%s", baseUrl, url.QueryEscape(character.Location.Name)))
	if err != nil {
		return []LocationData{}, err
	}
	defer res.Body.Close()

	var response rickAndMortyLocationResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	return response.Results, err
}
