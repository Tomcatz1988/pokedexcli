package main

import (
	"math/rand"
	"time"

	pokeapi "internal/pokeapi"
)

func catch(pokemon pokeapi.Pokemon, pokedex map[string]pokeapi.Pokemon) bool {
	clock := time.Now()
	h, m, s := clock.Clock()
	seed := int64(h * m * s)
	rGen := rand.New(rand.NewSource(seed))
	roll := rGen.Intn(500)
	if roll > pokemon.BaseExperience {
		pokedex[pokemon.Name] = pokemon
		return true
	}

	return false
}
