package main

import (
	"os"
	"testing"
)

func TestHandleArguments(t *testing.T) {
	expected := 1
	os.Args = []string{"cmd", "1"}

	actual := handleArguments()

	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
	}
}

func TestSecondsToMinutes(t *testing.T) {
	expected := "1 minutes, 1 seconds"
	actual := secondsToMinutes(61)

	if actual != expected {
		t.Errorf("Test failed, expected: %s, got: %s", expected, actual)
	}
}

// tests

// is input command line entry set
// is input an integer
// is input integer not negative
// is input a very large number

// 0 = none
// 1 = 1 second
// 2 = 2 seconds
// 60 = 1 minute
// 61 = 1 minute and 1 second
// 120 = 2 minutes
// 121 = 2 minutes and 1 second
// 122 = 2 minutes and 2 seconds
// 3600 = 1 hour
// 3600 = 1 hour and 1 second
// 3722 = 1 hour, 2 minutes and 2 seconds
// 86400 = 1 day
// 273660 =  3 days, 4 hours and 1 minute
// 31536000 = 1 year
// 94609440 = 3 years and 24 minutes
