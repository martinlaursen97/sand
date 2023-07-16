package core

import (
	"image/color"

	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

type WallParticle struct {
	Immovable
}

type Wall interface {
	CreateWallParticle(x, y int) *WallParticle
}

func CreateWallParticle(x, y int) *WallParticle {
	return &WallParticle{
		Immovable{
			BaseParticle{
				Position: maths.Vector{X: float64(x), Y: float64(y)},
				Size:     1,
				Color:    randomWallColor(),
			},
		},
	}
}

func randomWallColor() color.RGBA {
	return color.RGBA{
		R: utils.RandomUnsignedByteInRange(160, 180),
		G: utils.RandomUnsignedByteInRange(160, 180),
		B: utils.RandomUnsignedByteInRange(160, 180),
		A: 0,
	}
}

func (wp *WallParticle) Update(world *World, dt float64) {
}
