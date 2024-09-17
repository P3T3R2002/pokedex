package main

import (
	"fmt"
	//"log"
	"math/rand"
	"errors"
    "github.com/P3T3R2002/pokedex/pokeapi"
	"github.com/P3T3R2002/pokedex/pokeapi/pokecache"
)

const start_url string = "https://pokeapi.co/api/v2/location/?offset=0&limit=5"
const next_compare_url string = "https://pokeapi.co/api/v2/location/?offset=5&limit=5"

type Pokedex struct {
	pokemons		map[string]pokecache.Pokemon
	current_area	*pokecache.Area
	commands 		map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Pokedex, *pokecache.Cache, string) error
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
		pokemons:		map[string]pokecache.Pokemon{},
		current_area: 	&area,
		commands:		map[string]cliCommand{
							"help": {
								name:        "help",
								description: "Displays help messages.",
								callback:    commandHelp,
							},
							"exit": {
								name:        "exit",
								description: "Exit the Pokedex.",
								callback:    commandExit,
							},
							"map": {
								name:        "map",
								description: "Goes forwards on the map.",
								callback:    commandMap,
							},
							"mapb": {
								name:        "mapb",
								description: "Goes backwards on the map.",
								callback:    commandMap_back,
							},
							"explore": {
								name:        "explore",
								description: "Explores a location and displays the located pokemons.",
								callback:    commandExplore,
							},
							"catch": {
								name:        "catch",
								description: "Tries to catch a Pokemon.",
								callback:    commandCatch,
							},
							"inspect": {
								name:        "inspect",
								description: "Shows a cought Pokemons stats.",
								callback:    commandInspect,
							},
							"pokedex": {
								name:        "pokedex",
								description: "Shows all cought pokemon.",
								callback:    commandPokedex,
							},
		},
	}, cache, nil
}

//-----------------------------------------------------------------------

func commandExit(_ *Pokedex, _ *pokecache.Cache, _ string) error {
	return errors.New("Exit")
}

//-----------------------------------------------------------------------

func commandHelp(pokedex *Pokedex, _ *pokecache.Cache, _ string) error {
	fmt.Println("")
	for _, val := range pokedex.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	fmt.Println("")
	return nil
}

//-----------------------------------------------------------------------

func commandMap(pokedex *Pokedex, cache *pokecache.Cache, _ string) error {
	print_map(pokedex.current_area, true)
	step(pokedex, cache, true)
	return nil
}

//-----------------------------------------------------------------------

func commandMap_back(pokedex *Pokedex, cache *pokecache.Cache, _ string) error {
	if pokedex.current_area.Next == next_compare_url {
		fmt.Println("Cant go back!")
		return nil
	}
	err := step(pokedex, cache, false)
	if err != nil {
		return err
	}
	print_map(pokedex.current_area, false)
	return nil
}

//-----------------------------------------------------------------------

func commandExplore(_ *Pokedex, cache *pokecache.Cache, place string) error {
	url := "https://pokeapi.co/api/v2/location-area/"+place+"-area/"
	fmt.Println("Pokemon found:")
	err := pokeapi.Get_pockemon(url, cache)
	if err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------------------------------

func commandCatch(pokedex *Pokedex, cache *pokecache.Cache, pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	poke, err := pokeapi.Catch_pockemon(pokemon, cache)
	if err != nil {
		return err
	}
	num := rand.Intn(250)
	if num < poke.Base_experience {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon)
	pokedex.pokemons[pokemon] = poke
	return nil
}

//-----------------------------------------------------------------------

func commandInspect(pokedex *Pokedex, cache *pokecache.Cache, pokemon string) error {
	poke, ok := pokedex.pokemons[pokemon]
	if !ok {
		fmt.Println("You have yet to catch that pokemon!")
		return nil
	}
	pokeapi.Inspect_pokemon(poke)
	return nil
}

//-----------------------------------------------------------------------

func commandPokedex(pokedex *Pokedex, cache *pokecache.Cache, pokemon string) error {
	for _, val := range pokedex.pokemons {
		fmt.Printf("  -%s\n", val.Name)
	}
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

