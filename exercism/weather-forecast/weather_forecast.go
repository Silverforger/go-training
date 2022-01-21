//Package weather determines the weather condition of a location input.
package weather

//CurrentCondition stores the forecast for the location.
var CurrentCondition string

//CurrentLocation stores the location input.
var CurrentLocation string

//Forecast function returns the current weather condition of the specified input location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
