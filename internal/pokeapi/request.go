package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

func request(url string) ([]byte, error) {

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

	return body, nil
}
