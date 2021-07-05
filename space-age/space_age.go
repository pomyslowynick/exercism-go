// space has functions used to calculate age at different planets
package space

import (
	"math"
)

type Planet string

// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years

// PrecisionRound is a helper function rounding a number to given precision point
func PrecisionRound(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// GetEarthYears calculates the number of human years given seconds lived on another planet and the unti of orbital period
func GetEarthYears(secs float64, unit float64) float64 {
	earthSeconds := (secs / unit) / 31557600.0
	return PrecisionRound(earthSeconds, 0.01)
}

// Age outputs a value of earth years for a given planet
func Age(secs float64, planet Planet) (actual float64) {
	switch planet {
	case "Mercury":
		return GetEarthYears(secs, 0.2408467)
	case "Venus":
		return GetEarthYears(secs, 0.61519726)
	case "Earth":
		return GetEarthYears(secs, 1.0)
	case "Mars":
		return GetEarthYears(secs, 1.8808158)
	case "Jupiter":
		return GetEarthYears(secs, 11.862615)
	case "Saturn":
		return GetEarthYears(secs, 29.447498)
	case "Uranus":
		return GetEarthYears(secs, 84.016846)
	case "Neptune":
		return GetEarthYears(secs, 164.79132)
	default:
		return 0.0
	}
}
