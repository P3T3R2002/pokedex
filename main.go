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
	pokedex, cache, err := pokedex_setup()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pokedex.current_area.Next)
	fmt.Println(pokedex.current_area.Previous)
    scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if err != nil {
		log.Fatal(err)
	}
	//*********//
	for true{
		fmt.Println("==========")
		fmt.Println("Commands:")
		for key, _ := range pokedex.commands {
			fmt.Println("   "+key)
		}
		fmt.Printf("Pokedex > ")
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
			return 
		}
		text, arg := Get_command(scanner)
		_, ok := pokedex.commands[text]
		if !ok {
			fmt.Println("Wrong command!")
			continue
		}
		//**********//
		err = pokedex.commands[text].callback(&pokedex, cache, arg)
		if err != nil {
			log.Fatal(err)
			return 
		}
	}
	return 
}

//-----------------------------------------------------------------------

func Get_command(scanner *bufio.Scanner) (string, string) {
	var text string
	scanner.Scan()
	text = scanner.Text()
	switch text {
		case "catch":
			fmt.Printf("Pokemon to catch > ")
			scanner.Scan()
			return text, scanner.Text()
		case "explore":
			fmt.Printf("Area to explore > ")
			scanner.Scan()
			return text, scanner.Text()
		case "inspect":
			fmt.Printf("Pokemon to inspect > ")
			scanner.Scan()
			return text, scanner.Text()
		default:
			return text, ""
	}
}