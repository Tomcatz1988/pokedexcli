package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationBatch struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationBatch(url string) (batch LocationBatch, err error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationBatch{}, fmt.Errorf("pokeapi.GetLocationBatch() - response error: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationBatch{}, fmt.Errorf("pokeapi.GetLocationBatch() - read error: %w", err)
	}

	if err = json.Unmarshal(data, &batch); err != nil {
		return LocationBatch{}, fmt.Errorf("pokeapi.GetLocationBatch() - unmarshal error: %w", err)
	}

	return batch, nil
}
