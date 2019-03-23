package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var input = handleArguments()

	fmt.Println(input)
}

func handleArguments() int {

	var inputSeconds int

	// first arg is the program's path itself
	if len(os.Args) == 1 {

		inputSeconds = 0

		// extra arguments are ignored
	} else {

		userInput, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("You must specify an integer number of seconds")
			log.Fatalf("Error parsing input: %v", err)
		}

		inputSeconds = userInput

	}

	return inputSeconds

}
