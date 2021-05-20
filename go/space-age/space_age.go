package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var benchmarkYear float64;
	benchmarkYear = 31557600;
	var age float64;
	age = 0;
	switch (planet) {
		case "Mercury": 
			benchmarkYear *= 0.2408467;
			age = seconds / benchmarkYear;
		case "Venus": 
			benchmarkYear *= 0.61519726;
			age = seconds / benchmarkYear;
		case "Earth": 
			benchmarkYear *= 1.0;
			age = seconds / benchmarkYear
		case "Mars": 
			benchmarkYear *= 1.8808158;
			age = seconds / benchmarkYear;
		case "Jupiter": 
			benchmarkYear *= 11.862615;
			age = seconds / benchmarkYear;
		case "Saturn": 
			benchmarkYear *= 29.447498;
			age = seconds / benchmarkYear;
		case "Uranus": 
			benchmarkYear *= 84.016846;
			age = seconds / benchmarkYear;
		case "Neptune": 
			benchmarkYear *= 164.79132;
			age = seconds / benchmarkYear;
	}
	return age;
}
