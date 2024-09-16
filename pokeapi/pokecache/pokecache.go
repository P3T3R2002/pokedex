package pokecache

import (
	"sync"
	"errors"
	"fmt"
	//"log"
)

type Cache struct {
	cache 	map[string]interface{}
	mu 		*sync.Mutex
}
	//*********

type Area struct{
	_			int 	`json:"count"`
	Next		string 	`json:"next"`
	Previous	string 	`json:"previous"`
	Results		[]Place `json:"results"`
}

type Place struct {
	Name	string `json:"name"`
	_		string `json:"url"`
}
	//*********

type Location_Area struct {
	_ 					any 		`json:"encounter_method_rates"`
	_ 					any 		`json:"location"`
	_ 					any 		`json:"names"`
	Pokemon_encounters	[]Encounter 	`json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon 	Pokemon	`json:"pokemon"`
	_ 			any			`json:"version_details"`

}

type Pokemon struct {
	Name 	string	`json:"name"`
	_		string 	`json:"url"`
}

	   //\\
	  //**\\
	 //****\\
	//******\\

func Create_cache() *Cache {
	cache := make(map[string]interface{})
	pokemon := make(map[string][]Encounter)
	location := make(map[string][]Place)
	cache["pokemon"] = pokemon
	cache["location"] = location

	return &Cache{
		cache:	cache,
		mu:		&sync.Mutex{},
	}
}

	//*********//

func Get_location(url string) Area {
	return Area{}
}

	//*********//

func Write_place_cache(cache *Cache, url string, area *Area) {
	var places []Place
	for _, place := range area.Results {
		places = append(places, Place{Name:place.Name})
	}
	places = append(places, Place{Name:area.Next})
	places = append(places, Place{Name:area.Previous})
	cache.mu.Lock()
	assert := cache.cache["location"].(map[string][]Place)
	assert[url] = places
	cache.mu.Unlock()
	return

}

	//*********//

func Read_place_cache(cache *Cache, next string, area *Area) (error) {
	cache.mu.Lock()
	assert, ok := cache.cache["location"].(map[string][]Place)
	if !ok {
		cache.mu.Unlock()
		return errors.New("Type problem!")
	}
	places, ok := assert[next]
	cache.mu.Unlock()
	if !ok {
		return errors.New("Not in cache!")
	}
	area.Results = places[0:len(places)-3]
	area.Next = places[len(places)-2].Name
	area.Previous = places[len(places)-1].Name
	return nil
}

//-----------------------------------------------------------------------

func Write_pokemon_cache(cache *Cache, url string, area *Location_Area) {
	var encount []Encounter
	for _, encounter := range area.Pokemon_encounters {
		encount = append(encount, Encounter{Pokemon: encounter.Pokemon})
	}
	cache.mu.Lock()
	assert := cache.cache["pokemon"].(map[string][]Encounter)
	assert[url] = encount
	cache.mu.Unlock()
	return
}

	//*********//

func Read_pokemon_cache(cache *Cache, explore string, area *Location_Area) (error) {
	cache.mu.Lock()
	assert := cache.cache["pokemon"].(map[string][]Encounter)
	encount, ok := assert[explore]
	if !ok {
		cache.mu.Unlock()
		return errors.New("Not in cache!")
	}
	area.Pokemon_encounters = encount
	cache.mu.Unlock()
	return nil
}
