package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var input = handleArguments()

	var output = secondsToMinutes(input)

	fmt.Println(output)

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

func secondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	out := fmt.Sprintf("%d minutes, %d seconds", minutes, seconds)
	return out
}
