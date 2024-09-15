package main

import (
	"fmt"
	"errors"
    "github.com/P3T3R2002/pokedex/pokeapi"
	"github.com/P3T3R2002/pokedex/pokeapi/pokecache"
)

const start_url string = "https://pokeapi.co/api/v2/location/?offset=0&limit=5"
const next_compare_url string = "https://pokeapi.co/api/v2/location/?offset=5&limit=5"

type Pokedex struct {
	current_area	*pokecache.Area
	commands 		map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(Pokedex, *pokecache.Cache) error
}
	 //\\
    //**\\
   //****\\
  //******\\

func pokedex_setup() (Pokedex, *pokecache.Cache, error) {
	area := pokecache.Get_location(start_url)
	cache := pokecache.Create_cache()
	err := pokeapi.Update_location(start_url, &area, cache)
	if err != nil {
		return Pokedex{}, &pokecache.Cache{}, err
	}
	return Pokedex{
		current_area: 	&area,
		commands:		map[string]cliCommand{
							"help": {
								name:        "help",
								description: "Displays a help message",
								callback:    commandHelp,
							},
							"exit": {
								name:        "exit",
								description: "Exit the Pokedex",
								callback:    commandExit,
							},
							"map": {
								name:        "map",
								description: "Goes forwards on the map",
								callback:    commandMap,
							},
							"mapb": {
								name:        "mapb",
								description: "Goes backwards on the map",
								callback:    commandMap_back,
							},
		},
	}, cache, nil
}

//-----------------------------------------------------------------------

func commandExit(_ Pokedex, _ *pokecache.Cache) error {
	return errors.New("Exit")
}

//-----------------------------------------------------------------------

func commandHelp(pokedex Pokedex, _ *pokecache.Cache) error {
	fmt.Println("")
	for _, val := range pokedex.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	fmt.Println("")
	return nil
}

//-----------------------------------------------------------------------

func commandMap(pokedex Pokedex, cache *pokecache.Cache) error {
	print_map(pokedex.current_area, true)
	step(&pokedex, cache, true)
	return nil
}

//-----------------------------------------------------------------------

func commandMap_back(pokedex Pokedex, cache *pokecache.Cache) error {
	if pokedex.current_area.Next == next_compare_url {
		fmt.Println("Cant go back!")
		return nil
	}
	err := step(&pokedex, cache, false)
	if err != nil {
		return err
	}
	print_map(pokedex.current_area, false)
	return nil
}

//-----------------------------------------------------------------------

func step(pokedex *Pokedex, cache *pokecache.Cache, dir bool) error {
	if dir {
		err := pokeapi.Update_location(pokedex.current_area.Next, pokedex.current_area, cache)
		if err != nil {
			return err
		}
	} else {
		err := pokeapi.Update_location(pokedex.current_area.Previous, pokedex.current_area, cache)
		if err != nil {
			return err
		}
	}
	return nil
}

//-----------------------------------------------------------------------

func print_map(area *pokecache.Area, dir bool) {
	if dir {
		for _, place := range area.Results {
			fmt.Println(place.Name)
		}
		return 
	}
	for i:=len(area.Results)-1; i >= 0; i-- {
		fmt.Println(area.Results[i].Name)
	}

	return
}

//-----------------------------------------------------------------------

