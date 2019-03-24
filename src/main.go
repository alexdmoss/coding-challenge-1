package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var totalSeconds = handleArguments()
	var seconds int
	var minutes int
	var hours int
	var days int
	var years int
	var output string

	years, days, hours, minutes, seconds = convertSecondsToForecast(totalSeconds)

	output = formatAndDisplayOutput(years, days, hours, minutes, seconds)

	// print to screen, but could be wrapped as alternate output
	fmt.Println(output)

}

// converts integer values to compliant string and prints to screen
func formatAndDisplayOutput(years int, days int, hours int, minutes int, seconds int) string {
	output := fmt.Sprintf("%d years, %d days, %d hours, %d minutes, %d seconds", years, days, hours, minutes, seconds)
	return output
}

func convertSecondsToForecast(inputSeconds int) (int, int, int, int, int) {

	minutes, seconds := minutesFromSeconds(inputSeconds)
	hours, minutes := hoursFromMinutes(minutes)
	days, hours := daysFromHours(hours)
	years, days := yearsFromDays(days)

	return years, days, hours, minutes, seconds

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

func minutesFromSeconds(inSeconds int) (int, int) {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	return minutes, seconds
}

func hoursFromMinutes(inSeconds int) (int, int) {
	hours := inSeconds / 60
	minutes := inSeconds % 60
	return hours, minutes
}

func daysFromHours(inSeconds int) (int, int) {
	days := inSeconds / 24
	hours := inSeconds % 24
	return days, hours
}

func yearsFromDays(inSeconds int) (int, int) {
	years := inSeconds / 365
	days := inSeconds % 365
	return years, days
}
