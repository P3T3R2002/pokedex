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
	ok ,err := get_from_cache(cache, "encounter", url, &loc_area)
	if ok {
		print_encounter(loc_area)
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
	print_encounter(loc_area)
	//*********//
	add_to_cache(cache, url, &loc_area, "encounter")
	return nil
}

	//*********//

func Catch_pockemon(pokemon string, cache *pokecache.Cache) (pokecache.Pokemon, error) {	
	url := "https://pokeapi.co/api/v2/pokemon/"+pokemon+"/"
	var poke pokecache.Pokemon
	ok ,err := get_from_cache(cache, "pokemon", pokemon, &poke)
	if ok {
		return poke, nil
	} else if err != nil {
		return pokecache.Pokemon{}, err
	}

	res, err := http.Get(url)
	if err != nil {
		return pokecache.Pokemon{}, err
	}
	defer res.Body.Close()

	//*********//
	if res.StatusCode >= 400 {
		return pokecache.Pokemon{}, errors.New(fmt.Sprintf("error status: %s", res.Status))
	}
	//*********//
	data, err := io.ReadAll(res.Body)
	if err != nil {
			return pokecache.Pokemon{}, err
	}
	//*********//
	err = json.Unmarshal(data, &poke)
	if err != nil {
		return pokecache.Pokemon{}, err
	}
	//*********//
	add_to_cache(cache, pokemon, &poke, "pokemon")
	return poke, nil
}

//-----------------------------------------------------------------------

func Inspect_pokemon(pokemon pokecache.Pokemon) {
	print_pokemon(pokemon)
	return
}

//-----------------------------------------------------------------------

func get_from_cache(cache *pokecache.Cache, from, url string, get_for interface{}) (bool, error) {
	var err error
	switch from {
	case "location": err = pokecache.Read_place_cache(cache, url, get_for.(*pokecache.Area))
	case "encounter": err = pokecache.Read_encounter_cache(cache, url, get_for.(*pokecache.Location_Area))
	case "pokemon": err = pokecache.Read_pokemon_cache(cache, url, get_for.(*pokecache.Pokemon))
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
	case "encounter": pokecache.Write_encounter_cache(cache, url, add_it.(*pokecache.Location_Area))
	case "pokemon": pokecache.Write_pokemon_cache(cache, url, add_it.(*pokecache.Pokemon))
	default: return 
	}
	return 
}

//-----------------------------------------------------------------------

func print_encounter(encounter pokecache.Location_Area) {
	for _, encount := range encounter.Pokemon_encounters {
		fmt.Println(encount.Pokemon.Name)
	}
	return
}

func print_pokemon(pokemon pokecache.Pokemon) {
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Printf("Stats:\n", )
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Base_stat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}
	return
} 