package utils

import (
	"math"
	"math/rand"

	"github.com/martinlaursen97/sand/config"
)

func RoundYCoordinate(y float64) float64 {
	if y-math.Floor(y) >= 0.5 {
		return math.Ceil(y)
	} else {
		return math.Floor(y)
	}
}

func CheckBounds(x, y float64) (float64, float64) {
	if x < 0 {
		x = 0
	}

	if x >= config.ScreenWidth {
		x = config.ScreenWidth - 1
	}

	if y < 0 {
		y = 0
	}

	if y >= config.ScreenHeight {
		y = config.ScreenHeight - 1
	}

	return x, y
}

func WithinBounds(x, y int) bool {
	if x < 0 {
		return false
	}

	if x >= config.ScreenWidth {
		return false
	}

	if y < 0 {
		return false
	}

	if y >= config.ScreenHeight {
		return false
	}

	return true
}

func RandomUnsignedByteInRange(min, max uint8) uint8 {
	return uint8(rand.Intn(int(max-min)) + int(min))
}
