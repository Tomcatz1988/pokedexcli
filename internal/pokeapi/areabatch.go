package pokeapi

import (
	"encoding/json"
	"fmt"

	pokecache "internal/pokecache"
)


type AreaBatch struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


func GetAreaBatch(url string, cache *pokecache.Cache) (batch AreaBatch, err error) {
	data, err := getJsonData(url, cache)
	if err != nil  {
		return batch, fmt.Errorf("GetAreaBatch: %w", err)
	}

	if err = json.Unmarshal(data, &batch); err != nil {
		return AreaBatch{}, fmt.Errorf("GetAreaBatch: %w", err)
	}

	return batch, nil
}

