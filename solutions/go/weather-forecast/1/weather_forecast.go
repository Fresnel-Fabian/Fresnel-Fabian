// Package weather provides tools to predict current weather.
package weather
// CurrentCondition represents the current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current city.
var CurrentLocation string

// Forecast returns a string which gives the current condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
