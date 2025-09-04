package day

import "math"

func findClosest(x int, y int, z int) int {

	distanceX := math.Abs(float64(x - z))
	distanceY := math.Abs(float64(y - z))

	if distanceX > distanceY {
		return 2
	} else if distanceX == distanceY {
		return 0
	} else {
		return 1
	}

}
