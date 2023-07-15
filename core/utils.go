package core

import "math"

func roundYCoordinate(y float64) float64 {
	if y-math.Floor(y) >= 0.5 {
		return math.Ceil(y)
	} else {
		return math.Floor(y)
	}
}

func checkBounds(x, y float64) (float64, float64) {
	if x < 0 {
		x = 0
	}

	if x >= worldWidth {
		x = worldWidth - 1
	}

	if y < 0 {
		y = 0
	}

	if y >= worldHeight {
		y = worldHeight - 1
	}

	return x, y
}