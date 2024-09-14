package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

//-----------------------------------------------------------------------

func main() {
	//*********//
	pokedex, err := pokedex_setup()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pokedex.current_area.Next)
	fmt.Println(pokedex.current_area.Previous)
    scanner := bufio.NewScanner(os.Stdin)
	//*********//
	for true{
		fmt.Println("Commands:")
		for key, _ := range pokedex.commands {
			fmt.Println("   "+key)
		}
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
			return 
		}
		//**********//
		err = pokedex.commands[scanner.Text()].callback(pokedex)
		if err != nil {
			log.Fatal(err)
			return 
		}
	}
	return 
}

