package internettime

import (
	"math"
	"testing"
	"time"
)

func TestDurationToInternetTime(t *testing.T) {
	testCases := []struct {
		durationStr string
		expected    float64
		expectError bool
	}{
		{"1h", 41.67, false},
		{"15m", 10.42, false},
		{"invalid", 0, true},
	}

	for _, tc := range testCases {
		got, err := DurationToInternetTime(tc.durationStr)
		if tc.expectError {
			if err == nil {
				t.Errorf("DurationToInternetTime(%s) expected an error, but got none", tc.durationStr)
			}
		} else {
			if err != nil {
				t.Errorf("DurationToInternetTime(%s) returned an unexpected error: %v", tc.durationStr, err)
			}
			if got != tc.expected {
				t.Errorf("DurationToInternetTime(%s) = %v, want %v", tc.durationStr, got, tc.expected)
			}
		}
	}
}

func TestInternetToStandardTime(t *testing.T) {
	now := time.Now()
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.FixedZone("Biel", 3600))
	tomorrowMidnight := todayMidnight.Add(24 * time.Hour)
	dayAfterTomorrowMidnight := tomorrowMidnight.Add(24 * time.Hour)

	secondsFor123_45Beats := math.Floor(123.45 * 86.4)
	durationFor123_45Beats := time.Duration(secondsFor123_45Beats) * time.Second
	secondsFor001_864Beats := math.Floor(0.01 * 86.4)

	testCases := []struct {
		internetTime float64
		expected     time.Time
	}{
		{0, todayMidnight},
		{1000, tomorrowMidnight},
		{250, todayMidnight.Add(6 * time.Hour)},
		{500, todayMidnight.Add(12 * time.Hour)},
		{750, todayMidnight.Add(18 * time.Hour)},
		{123.45, todayMidnight.Add(durationFor123_45Beats)},
		{999.99, tomorrowMidnight.Add(-time.Duration(secondsFor001_864Beats) * time.Second)},
		{-100, todayMidnight.Add(-time.Duration(100*86.4) * time.Second)},
		{2000, dayAfterTomorrowMidnight},
	}

	for _, tc := range testCases {
		got := InternetToStandardTime(tc.internetTime)

		if !timesAreEqualWithMargin(got, tc.expected, time.Second) {
			t.Errorf("InternetToStandardTime(%v) = %v, want %v", tc.internetTime, got, tc.expected)
		}
	}
}

func TestStandardToInternetTime(t *testing.T) {
	// Define test cases with known standard times and expected Internet times
	testCases := []struct {
		standardTime         time.Time
		expectedInternetTime float64
	}{
		// Test case example: assuming the correct conversion for a specific time
		{time.Date(2023, 4, 1, 0, 0, 0, 0, time.FixedZone("Biel", 3600)), 0},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		gotInternetTime := StandardToInternetTime(tc.standardTime)

		// Allow a small margin of error due to float precision
		if !floatsAreEqualWithinMargin(gotInternetTime, tc.expectedInternetTime, 0.001) {
			t.Errorf("StandardToInternetTime(%v) = %v, want %v", tc.standardTime, gotInternetTime, tc.expectedInternetTime)
		}
	}
}

// floatsAreEqualWithinMargin checks if two float64 values are equal within a specified margin.
func floatsAreEqualWithinMargin(a, b, margin float64) bool {
	return a > b-margin && a < b+margin
}

// timesAreEqualWithMargin checks if two times are equal within a certain margin.
func timesAreEqualWithMargin(a, b time.Time, margin time.Duration) bool {
	diff := a.Sub(b)
	if diff < 0 {
		diff = -diff
	}
	return diff <= margin
}
