package pokecache

import (
	"sync"
	"errors"
)

type Cache struct {
	cache 	map[string][]map[string]any
	mu 		*sync.Mutex
}

type Pokemon struct {
	pokemon 	string
}

type Area struct{
	_		int 	`json:"count"`
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Results		[]Place `json:"results"`
}

type Place struct {
	Name	string `json:"name"`
	_		string `json:"url"`
}
	   //\\
	  //**\\
	 //****\\
	//******\\

func Create_cache() *Cache {
	return &Cache{
		cache:	{
			pokemon:	[map[string]Pokemon{}],
			location: 	[map[string][]Place{}, map[string]map[string]string{}],
		}
		mu:		&sync.Mutex{},
	}
}

	//*********//

func Get_location(url string) Area {
	return Area{}
}

	//*********//

func Write_cache(cache *Cache, url string, area *Area) {
	var places []Place
	for _, place := range area.Results {
		places = append(places, Place{Name:place.Name})
	}
	dir := map[string]string{}
	dir["next"] = area.Next
	dir["prev"] = area.Previous
	cache.mu.Lock()
	cache.places[url][0] = places
	cache.places[url][1] = dir
	cache.mu.Unlock()
	return

}

	//*********//

func Read_cache(cache *Cache, next string, area *Area) (error) {
	cache.mu.Lock()
	places, ok := cache.places[next]
	dir, ko := cache.dir[next]
	cache.mu.Unlock()
	if !ok && !ko {
		return errors.New("Not in cache!")
	} else if !ok || !ko {
		return errors.New("Unknown problem in cache!")
	}
	area.Results = places
	area.Next = dir["next"]
	area.Previous = dir["prev"]
	return nil
}