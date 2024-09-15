package main

import (
	"fmt"
	"errors"
)

const start_url string = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
const next_compare_url string = "https://pokeapi.co/api/v2/location/?offset=20&limit=20"

type Pokedex struct {
	current_area	*Area
	commands 		map[string]cliCommand
}

type cliCommand struct {
	name        string
	description string
	callback    func(Pokedex) error
}
	 //\\
    //**\\
   //****\\
  //******\\

func pokedex_setup() (Pokedex, error) {
	area := get_location(start_url)
	err := update_location(start_url, &area)
	if err != nil {
		return Pokedex{}, err
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
	}, nil
}

//-----------------------------------------------------------------------

func commandExit(_ Pokedex) error {
	return errors.New("Exit")
}

//-----------------------------------------------------------------------

func commandHelp(pokedex Pokedex) error {
	fmt.Println("")
	for _, val := range pokedex.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	fmt.Println("")
	return nil
}

//-----------------------------------------------------------------------

func commandMap(pokedex Pokedex) error {
	print_map(pokedex.current_area, true)
	step(&pokedex, true)
	return nil
}

//-----------------------------------------------------------------------

func commandMap_back(pokedex Pokedex) error {
	if pokedex.current_area.Next == next_compare_url {
		fmt.Println("Cant go back!")
		return nil
	}
	err := step(&pokedex, false)
	if err != nil {
		return err
	}
	print_map(pokedex.current_area, false)
	return nil
}

//-----------------------------------------------------------------------

func step(pokedex *Pokedex, dir bool) error {
	if dir {
		err := update_location(pokedex.current_area.Next, pokedex.current_area)
		if err != nil {
			return err
		}
	} else {
		err := update_location(pokedex.current_area.Previous, pokedex.current_area)
		if err != nil {
			return err
		}
	}
	return nil
}

//-----------------------------------------------------------------------

func print_map(area *Area, dir bool) {
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

