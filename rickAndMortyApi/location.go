package rickAndMortyApi

import "time"

type LocationData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	Residents []string  `json:"residents"`
	URL       string    `json:"url"`
	Created   time.Time `json:"created"`
}

type rickAndMortyLocationResponse struct {
	Info struct {
		Count int `json:"count"`
		Pages int `json:"pages"`
		Next  interface {
		} `json:"next"`
		Prev interface {
		} `json:"prev"`
	} `json:"info"`
	Results []LocationData `json:"results"`
}
