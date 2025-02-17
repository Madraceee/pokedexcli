package pokeapi

import (
	"encoding/json"
	"errors"
)

func (c *PokeAPIClient) ExploreLocation(name string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + name
	body, err := c.request(url)
	if err != nil {
		return LocationDetails{}, err
	}

	locationDetails := LocationDetails{}
	err = json.Unmarshal(body, &locationDetails)
	if err != nil {
		return LocationDetails{}, errors.New("Unable to format Data")
	}

	return locationDetails, nil
}
