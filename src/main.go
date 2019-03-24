package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TimeValues struct {
	Years   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func main() {

	output := FormatTime()

	// print to screen, but could be wrapped in alternatives
	fmt.Println(output)

}

// FormatTime receives an integer number of seconds and returns a string in human-readable format
func FormatTime() string {

	var output string
	var seconds int
	var minutes int
	var hours int
	var days int
	var years int

	totalSeconds := handleArguments()

	if totalSeconds == 0 {

		output = "None"

	} else {

		years, days, hours, minutes, seconds = convertSecondsToForecast(totalSeconds)

		output = formatAndDisplayOutput(years, days, hours, minutes, seconds)

	}

	return output

}

// converts integer values to compliant string and prints to screen
func formatAndDisplayOutput(years int, days int, hours int, minutes int, seconds int) string {

	// if one number, just print it
	// if we have two numbers, print with an 'and'
	// if we have more than two numbers, print with ',' until the last which is 'and'

	var output []string

	calculatedValues := TimeValues{
		Years:   years,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}

	if calculatedValues.Years > 0 {
		output = append(output, fmt.Sprintf("%d years", calculatedValues.Years))
	}
	if calculatedValues.Days > 0 {
		output = append(output, fmt.Sprintf("%d days", calculatedValues.Days))
	}
	if calculatedValues.Hours > 0 {
		output = append(output, fmt.Sprintf("%d hours", calculatedValues.Hours))
	}
	if calculatedValues.Minutes > 0 {
		output = append(output, fmt.Sprintf("%d minutes", calculatedValues.Minutes))
	}
	if calculatedValues.Seconds > 0 {
		output = append(output, fmt.Sprintf("%d seconds", calculatedValues.Seconds))
	}

	// join with commas then replace last comma with 'and'
	out := fmt.Sprint(strings.Join(output, ", "))
	i := strings.LastIndex(out, ",")
	if i > 0 {
		out = out[:i] + strings.Replace(out[i:], ",", " and", 1)
	}

	return out

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
