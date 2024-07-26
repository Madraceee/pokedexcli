package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

func (c *PokeAPIClient) request(url string) ([]byte, error) {
	ans, present := c.cache.Get(url)
	if present == true {
		return ans, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Failed to fetch data")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New("Failed to fetch data")
	}

	if err != nil {
		return nil, errors.New("Failed to fetch data")
	}

	c.cache.Add(url, body)
	return body, nil
}
