// Package weather provides functionality to track and report
// the current weather conditions for a given location.
// It stores the latest weather state and formats forecasts
// for display.
package weather

// CurrentCondition stores the latest recorded weather condition
// (e.g., "Sunny", "Rainy") for the current location.
var CurrentCondition string

// CurrentLocation stores the name of the location (city)
// for which the current weather condition is recorded.
var CurrentLocation string

// Forecast updates the current weather information for a given location
// and returns a formatted string describing the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
