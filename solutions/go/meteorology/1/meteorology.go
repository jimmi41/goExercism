package meteorology
import "strconv"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)
func (sc TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[sc]
}
// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}
func (t Temperature) String() string {
	return strconv.Itoa(t.degree)+" " + t.unit.String()
}
// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)
func (s SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[s]
}
// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (t Speed) String() string {
	return strconv.Itoa(t.magnitude) +" "+ t.unit.String()
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}
func (m MeteorologyData) String() string {
	return m.location +
		": " +
		m.temperature.String() +
		", Wind " +
		m.windDirection +
		" at " +
		m.windSpeed.String() +
		", " +
		strconv.Itoa(m.humidity) +
		"% Humidity"
}
// Add a String method to MeteorologyData
