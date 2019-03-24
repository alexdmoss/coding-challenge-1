package main

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func getFunction() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	functionParts := strings.Split(frame.Function, ".")
	function := functionParts[len(functionParts)-1]
	return function
}

func TestHandleArguments(t *testing.T) {
	expected := 1
	os.Args = []string{"cmd", "1"}

	actual := handleArguments()

	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got:  '%d'", expected, actual)
	}
}

func TestMinutesFromSeconds(t *testing.T) {

	thisFunc := getFunction()
	expectedMinutes := 1
	expectedSeconds := 1
	actualMinutes, actualSeconds := minutesFromSeconds(61)

	if actualMinutes != expectedMinutes {
		t.Errorf("%s (Minutes): Expected  %d - Actual %d", thisFunc, expectedMinutes, actualMinutes)
	}
	if actualSeconds != expectedSeconds {
		t.Errorf("%s (Seconds): Expected %d - Actual %d", thisFunc, expectedSeconds, actualSeconds)
	}
}

func TestHoursFromMinutes(t *testing.T) {

	thisFunc := getFunction()
	expectedHours := 1
	expectedMinutes := 1
	actualHours, actualMinutes := hoursFromMinutes(61)

	if actualHours != expectedHours {
		t.Errorf("%s (Hours): Expected %d - Actual %d", thisFunc, expectedHours, actualHours)
	}
	if actualMinutes != expectedMinutes {
		t.Errorf("%s (Minutes): Expected %d - Actual %d", thisFunc, expectedMinutes, actualMinutes)
	}

}

func TestDaysFromHours(t *testing.T) {

	thisFunc := getFunction()
	expectedHours := 1
	expectedDays := 1
	actualDays, actualHours := daysFromHours(25)

	if actualDays != expectedDays {
		t.Errorf("%s (Days): Expected %d - Actual %d", thisFunc, expectedDays, actualDays)
	}
	if actualHours != expectedHours {
		t.Errorf("%s (Hours): Expected %d - Actual %d", thisFunc, expectedHours, actualHours)
	}

}

func TestYearsFromDays(t *testing.T) {

	thisFunc := getFunction()
	expectedYears := 1
	expectedDays := 1
	actualYears, actualDays := yearsFromDays(366)

	if actualYears != expectedYears {
		t.Errorf("%s (Years): Expected %d - Actual %d", thisFunc, expectedYears, actualYears)
	}
	if actualDays != expectedDays {
		t.Errorf("%s (Days): Expected %d - Actual %d", thisFunc, expectedDays, actualDays)
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
