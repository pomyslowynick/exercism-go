package raindrops

import "fmt"

func Convert(num int) (sound string) {
	if num%3 == 0 {
		sound += "Pling"
	}
	if num%5 == 0 {
		sound += "Plang"
	}
	if num%7 == 0 {
		sound += "Plong"
	}

	if len(sound) == 0 {
		return fmt.Sprint(num)
	}

	return sound

}
