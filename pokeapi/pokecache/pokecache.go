package pokecache

import (
	"sync"
	"errors"
	//	"fmt"
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
	Pokemon 	PokeName	`json:"pokemon"`
	_ 			any			`json:"version_details"`

}

type PokeName struct {
	Name 	string	`json:"name"`
	_		string 	`json:"url"`
}
	//*********

type Pokemon struct {
	_					any
  	Name 				string
  	Base_experience 	int
  	Height				int
  	_					any
  	Order				int
  	Weight				int
  	Abilities			[]Ability
	_ 					any
	_ 					any
	_ 					any
	_ 					any
	Moves				[]Move
	_ 					any
	_ 					any
	_ 					any
	Stats				[]Stat
	Types				[]Type
	_ 					any
}
//----------------//

type Ability struct {
	_		any
	_		any
	Ability	AbiName
}

type AbiName struct {
	Name	string
	_		any
}
//----------------//

type Move struct {
	Move	MoveName
	_		any
}

type MoveName struct {
	Name	string
	_		any
}
//----------------//

type Stat struct {
	Base_stat	int
	_			any
	Stat		StatName
}

type StatName struct {
	Name	string
	_		any
}
//----------------//

type Type struct {
	_		any
	Type	TypeName
}

type TypeName struct {
	Name	string
	_		any
}
	   //\\
	  //**\\
	 //****\\
	//******\\

func Create_cache() *Cache {
	cache := make(map[string]interface{})
	encounter := make(map[string][]Encounter)
	location := make(map[string][]Place)
	pokemon := make(map[string]Pokemon)
	cache["encounter"] = encounter
	cache["location"] = location
	cache["pokemon"] = pokemon

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

func Write_encounter_cache(cache *Cache, url string, area *Location_Area) {
	var encount []Encounter
	for _, encounter := range area.Pokemon_encounters {
		encount = append(encount, Encounter{Pokemon: encounter.Pokemon})
	}
	cache.mu.Lock()
	assert := cache.cache["encounter"].(map[string][]Encounter)
	assert[url] = encount
	cache.mu.Unlock()
	return
}

	//*********//

func Read_encounter_cache(cache *Cache, explore string, area *Location_Area) (error) {
	cache.mu.Lock()
	assert := cache.cache["encounter"].(map[string][]Encounter)
	encount, ok := assert[explore]
	if !ok {
		cache.mu.Unlock()
		return errors.New("Not in cache!")
	}
	area.Pokemon_encounters = encount
	cache.mu.Unlock()
	return nil
}

//-----------------------------------------------------------------------

func Write_pokemon_cache(cache *Cache, _ string, pokemon *Pokemon) {
	cache.mu.Lock()
	assert := cache.cache["pokemon"].(map[string]Pokemon)
	assert[pokemon.Name] = *pokemon
	cache.mu.Unlock()
	return
}

	//*********//

func Read_pokemon_cache(cache *Cache, poke string, pokemon *Pokemon) (error) {
	cache.mu.Lock()
	assert := cache.cache["pokemon"].(map[string]Pokemon)
	mon, ok := assert[poke]
	if !ok {
		cache.mu.Unlock()
		return errors.New("Not in cache!")
	}
	
	pokemon.Base_experience = mon.Base_experience
	pokemon.Abilities = mon.Abilities	
	pokemon.Height = mon.Height			
	pokemon.Weight = mon.Weight				
	pokemon.Order = mon.Order				
	pokemon.Moves = mon.Moves			
	pokemon.Types = mon.Types	
	pokemon.Name = mon.Name 	
	cache.mu.Unlock()
	return nil
}