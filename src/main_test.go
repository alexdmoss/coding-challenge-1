package main

import (
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestFormatAndDisplayOutput(t *testing.T) {

	thisFunc := getFunction()
	expected := "1 years, 2 days, 3 hours, 4 minutes, 5 seconds"
	actual := formatAndDisplayOutput(1, 2, 3, 4, 5)

	if actual != expected {
		t.Errorf("%s (Minutes): Expected  %s - Actual %s", thisFunc, expected, actual)
	}

}

func TestConvertSecondsToForecast(t *testing.T) {

	thisFunc := getFunction()

	// format: input seconds => output years, days, hours, minutes, seconds
	tables := []struct {
		inputSeconds    int
		expectedYears   int
		expectedDays    int
		expectedHours   int
		expectedMinutes int
		expectedSeconds int
	}{
		// 1 or 2 Y/D/H/M/S
		{1, 0, 0, 0, 0, 1},
		{2, 0, 0, 0, 0, 2},
		{60, 0, 0, 0, 1, 0},
		{120, 0, 0, 0, 2, 0},
		{3600, 0, 0, 1, 0, 0},
		{7200, 0, 0, 2, 0, 0},
		{86400, 0, 1, 0, 0, 0},
		{172800, 0, 2, 0, 0, 0},
		{31536000, 1, 0, 0, 0, 0},
		{63072000, 2, 0, 0, 0, 0},

		// 2 of Y/D/H/M/S
		{61, 0, 0, 0, 1, 1},
		{3601, 0, 0, 1, 0, 1},
		{86401, 0, 1, 0, 0, 1},
		{31536001, 1, 0, 0, 0, 1},
		{3660, 0, 0, 1, 1, 0},
		{86460, 0, 1, 0, 1, 0},
		{31536060, 1, 0, 0, 1, 0},
		{90000, 0, 1, 1, 0, 0},
		{31539600, 1, 0, 1, 0, 0},
		{31622400, 1, 1, 0, 0, 0},
	}

	for _, table := range tables {
		actualYears, actualDays, actualHours, actualMinutes, actualSeconds := convertSecondsToForecast(table.inputSeconds)
		if actualYears != table.expectedYears {
			t.Errorf("%s(%d): Years   - Expected %d, Actual %d", thisFunc, table.inputSeconds, table.expectedYears, actualYears)
		}
		if actualDays != table.expectedDays {
			t.Errorf("%s(%d): Days    - Expected %d, Actual %d", thisFunc, table.inputSeconds, table.expectedDays, actualDays)
		}
		if actualHours != table.expectedHours {
			t.Errorf("%s(%d): Hours   - Expected %d, Actual %d", thisFunc, table.inputSeconds, table.expectedHours, actualHours)
		}
		if actualMinutes != table.expectedMinutes {
			t.Errorf("%s(%d): Minutes - Expected %d, Actual %d", thisFunc, table.inputSeconds, table.expectedMinutes, actualMinutes)
		}
		if actualSeconds != table.expectedSeconds {
			t.Errorf("%s(%d): Seconds - Expected %d, Actual %d", thisFunc, table.inputSeconds, table.expectedSeconds, actualSeconds)
		}
	}

}

func TestHandleArguments(t *testing.T) {

	os.Args = []string{"cmd", "1"}
	expected := 1
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

func getFunction() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	functionParts := strings.Split(frame.Function, ".")
	function := functionParts[len(functionParts)-1]
	return function
}

// tests

// is input command line entry set
// is input an integer
// is input integer not negative
// is input a very large number

// giant number

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

// formatting tests

// [done]
// S - 1, 2
// M - 60, 120
// H - 3600, 7200
// D - 86400, 86400*2
// Y - 3153600, 3153600*2

// [done]
// M S - 61
// H S - 3601
// D S - 86401
// Y S - 3153601
// H M - 3660
// D M - 86460
// Y M - 3153660
// D H - 90000
// Y H - 3157200
// Y D - 3240000

// H M S
// D M S
// Y M S
// D H M
// Y H M
// Y D H

// D H M S
// Y H M S
// Y D H M

// Y D H M S
