// Package internettime package provides functions to convert between Swatch Internet Time (.beats) and standard time.
package internettime

import (
	"math"
	"time"
)

// DurationToInternetTime converts a given duration string to Swatch Internet Time (.beats)
func DurationToInternetTime(durationStr string) (float64, error) {
	// Parse the duration string
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, err
	}

	// Convert duration to seconds
	totalSeconds := duration.Seconds()

	beats := totalSeconds / 86.4

	// Round to 2 decimal places
	roundedBeats := math.Round(beats*100) / 100

	return roundedBeats, nil
}

// InternetToStandardTime converts .beats to standard time in UTC+1
func InternetToStandardTime(beats float64) time.Time {
	// One beat is 86.4 seconds
	seconds := beats * 86.4

	// Get today's date in UTC+1
	now := time.Now().In(time.FixedZone("Biel", 3600))
	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Add the seconds to today's midnight
	standardTime := todayMidnight.Add(time.Second * time.Duration(seconds))

	return standardTime
}

// StandardToInternetTime converts the current time to Swatch Internet Time (.beats)
func StandardToInternetTime(t time.Time) float64 {
	// Convert current time to UTC+1
	bielTime := t.In(time.FixedZone("Biel", 3600))

	// Calculate time elapsed since midnight in seconds
	elapsedSinceMidnight := bielTime.Sub(time.Date(bielTime.Year(), bielTime.Month(), bielTime.Day(), 0, 0, 0, 0, bielTime.Location())).Seconds()

	// Convert to .beats (1 .beat = 86.4 seconds)
	beats := elapsedSinceMidnight / 86.4

	// Round to 2 decimal places
	roundedBeats := math.Round(beats*100) / 100

	return roundedBeats
}
