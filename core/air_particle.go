package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/maths"
)

type AirParticle struct {
	BaseParticle
}

func (ap *AirParticle) Update(world *World, dt float64) {
	ap.HasUpdated = true
}

func (ap *AirParticle) Draw(screen *ebiten.Image) {
	screen.Set(int(ap.Position.X), int(ap.Position.Y), ap.Color)
}

func (ap *AirParticle) Reset() {
	ap.HasUpdated = false
}

func CreateAirParticle(x, y int) *AirParticle {
	p := &AirParticle{
		BaseParticle: BaseParticle{
			Position: maths.Vector{X: float64(x), Y: float64(y)},
			Size:     config.ParticleSize,
			Color:    color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
	}

	return p
}
