package rickAndMortyApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Origin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Character struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Origin    `json:"origin"`
	Location Location  `json:"location"`
	Image    string    `json:"image"`
	Episode  []string  `json:"episode"`
	URL      string    `json:"url"`
	Created  time.Time `json:"created"`
}

type rickAndMortyCharacterResponse struct {
	Results []Character `json:"results"`
}

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
