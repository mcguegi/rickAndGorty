package rickAndMortyApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllCharacters() ([]Character, error) {
	res, err := http.Get(fmt.Sprintf("%s/character", baseUrl))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response rickAndMortyCharacterResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Results, err
}

func getCharacterByName(name string) (*Character, error) {
	res, err := http.Get(fmt.Sprintf("%s/character/?name=%s", baseUrl, name))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response *Character
	err = json.NewDecoder(res.Body).Decode(&response)

	return response, err
}
