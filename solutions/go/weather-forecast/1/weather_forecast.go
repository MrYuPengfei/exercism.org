// Package weather provides tools to forecast the current weather condition of various cities in Goblinocus. 
package weather

var (
    // CurrentCondition represents a certain current condition.
	CurrentCondition string
    // CurrentLocation  represents a certain current location.
	CurrentLocation  string
)
// Forecast returns an string to describ current weather condition. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
