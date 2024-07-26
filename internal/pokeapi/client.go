package pokeapi

import (
	"github.com/madraceee/pokedexcli/internal/pokecache"
	"time"
)

type PokeAPIClient struct {
	cache *pokecache.Cache
}

func GetNewClient(cacheInterval time.Duration) PokeAPIClient {
	cache := pokecache.NewCache(cacheInterval)
	return PokeAPIClient{
		cache: cache,
	}
}
