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
	Results LocationData `json:"results"`
}
