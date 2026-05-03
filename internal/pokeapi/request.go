package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	pokecache "internal/pokecache"
)


func getJsonData(url string, cache *pokecache.Cache) (data []byte, err error) {
	data, exists := cache.Get(url)
	if !exists {
		data, err = getFromAPI(url)
		if err != nil {
			return data, fmt.Errorf("getJsonData: %w", err)
		}

		cache.Add(url, data)
	}
	return data, nil
}


func getFromAPI(url string) (data []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("getFromAPI: %w", err)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("getFromAPI: %w", err)
	}

	return data, nil
}
