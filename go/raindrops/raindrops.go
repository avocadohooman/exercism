package raindrops

import "strconv"

func Convert(raindrop int) string {
	raindropSound := ""

	if raindrop%3 == 0 {
		raindropSound += "Pling"
	}
	if raindrop%5 == 0 {
		raindropSound += "Plang"
	}
	if raindrop%7 == 0 {
		raindropSound += "Plong"
	}
	if raindrop%3 != 0 && raindrop%5 != 0 && raindrop%7 != 0 {
		raindropSound = strconv.Itoa(raindrop)
	}
	return raindropSound
}
