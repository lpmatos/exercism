package space

/*
Given an age in seconds, calculate how old someone would be on:

* Mercury: orbital period 0.2408467 Earth years
* Venus: orbital period 0.61519726 Earth years
* Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
* Mars: orbital period 1.8808158 Earth years
* Jupiter: orbital period 11.862615 Earth years
* Saturn: orbital period 29.447498 Earth years
* Uranus: orbital period 84.016846 Earth years
* Neptune: orbital period 164.79132 Earth years

So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.
*/

const earthSeconds = 31557600

// Planet struct representation.
type Planet string

func calculeEarthYears(value float64) float64 {
	return value / earthSeconds
}

// Age function.
func Age(seconds float64, planet Planet) float64 {
	orbitalPeriod := map[Planet]float64{
		"Mercury": 0.2408467,  // Earth years
		"Venus":   0.61519726, // Earth years
		"Earth":   1,          // Earth years - 365.25 Earth days - 31557600 second
		"Mars":    1.8808158,  // Earth years
		"Jupiter": 11.8626215, // Earth years
		"Saturn":  29.447498,  // Earth years
		"Uranus":  84.016846,  // Earth years
		"Neptune": 164.79132,  // Earth years - 1821023456 second
	}
	orbitalPeriodEarthYearPorcent := orbitalPeriod[planet]
	return calculeEarthYears(seconds) / orbitalPeriodEarthYearPorcent
}
