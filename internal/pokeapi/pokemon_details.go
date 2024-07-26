package pokeapi

import "encoding/json"

func (c *PokeAPIClient) PokemonDetails(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	body, err := c.request(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
