// Package weather provides functions for weather forcast.
package weather

var (
    // CurrentCondition represents current weather condition.
	CurrentCondition string
    // CurrentLocation represents current weather location.
	CurrentLocation  string
)

// Forecast returns string containing message about weather condition and its location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
