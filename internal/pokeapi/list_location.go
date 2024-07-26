package pokeapi

import (
	"encoding/json"
	"errors"
)

func GetInitialLocationURL() string {
	return "https://pokeapi.co/api/v2/location-area"
}

func (c *PokeAPIClient) GetLocationList(url string) (Locations, error) {
	body, err := c.request(url)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Locations{}, errors.New("Unable to format Data")
	}

	return locations, nil
}
