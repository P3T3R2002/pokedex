package pokeapi

import (
	"errors"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
	//"log"
	"github.com/P3T3R2002/pokedex/pokeapi/pokecache"
)

func Update_location(url string, area *pokecache.Area, cache *pokecache.Cache) error {
	ok ,err := get_from_cache(cache, "location", url, area)
	if ok {
		return nil
	} else if err != nil {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//*********//
	if res.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("error status: %s", res.Status))
	}
	//*********//
	data, err := io.ReadAll(res.Body)
	if err != nil {
			return err
	}
	//*********//
	err = json.Unmarshal(data, area)
	if err != nil {
		return err
	}
	//*********//
	add_to_cache(cache, url, area, "location")
	return nil
}

	//*********//

func Get_pockemon(url string, cache *pokecache.Cache) error {
	var loc_area pokecache.Location_Area
	ok ,err := get_from_cache(cache, "pokemon", url, &loc_area)
	if ok {
		print_pokemon(loc_area)
		return nil
	} else if err != nil {
		return err
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//*********//
	if res.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("error status: %s", res.Status))
	}
	//*********//
	data, err := io.ReadAll(res.Body)
	if err != nil {
			return err
	}
	//*********//
	err = json.Unmarshal(data, &loc_area)
	if err != nil {
		return err
	}
	print_pokemon(loc_area)
	//*********//
	add_to_cache(cache, url, &loc_area, "pokemon")
	return nil
}

//-----------------------------------------------------------------------

func get_from_cache(cache *pokecache.Cache, from, url string, get_for interface{}) (bool, error) {
	var err error
	switch from {
	case "location": err = pokecache.Read_place_cache(cache, url, get_for.(*pokecache.Area))
	case "pokemon": err = pokecache.Read_pokemon_cache(cache, url, get_for.(*pokecache.Location_Area))
	default: return false, errors.New("Wrong case!")
	}
	if err != nil {
		if error.Error(err) == "Not in cache!" {
			return false, nil
		}
		return false, err
	} 
	return true, nil
}

	//*********//

func add_to_cache(cache *pokecache.Cache, url string, add_it interface{}, from string) {
	switch from {
	case "location": pokecache.Write_place_cache(cache, url, add_it.(*pokecache.Area))
	case "pokemon": pokecache.Write_pokemon_cache(cache, url, add_it.(*pokecache.Location_Area))
	default: return 
	}
	return 
}

//-----------------------------------------------------------------------

func print_pokemon(pokemons pokecache.Location_Area) {
	for _, encount := range pokemons.Pokemon_encounters {
		fmt.Println(encount.Pokemon.Name)
	}
	return
}