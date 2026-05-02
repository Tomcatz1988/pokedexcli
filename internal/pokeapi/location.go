package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/Tomcatz1988/pokedexcli/internal/pokecache"
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


func GetLocationBatch(url string, cache pokecache.Cache) (batch LocationBatch, err error) {
	data, exists := cache.Get(url)
	if !exists {
		data, err = getBatchFromWeb(url)
		if err != nil {
			return LocationBatch{}, err
		}
		cache.Add(url, data)
	} else {
		fmt.Println("result from cache")
	}
	if err = json.Unmarshal(data, &batch); err != nil {
		return LocationBatch{}, fmt.Errorf("pokeapi.GetLocationBatch() - unmarshal error: %w", err)
	}

	return batch, nil
}


func getBatchFromWeb(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("pokeapi.GetLocationBatch() - response error: %w", err)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("pokeapi.GetLocationBatch() - read error: %w", err)
	}

	return data, nil
}
