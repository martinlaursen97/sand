package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle interface {
	Update(world *World, dt float64)
	Draw(screen *ebiten.Image)
	Reset()
	GetPosition() Vector
	SetPosition(Vector)
}

//// BaseParticle ////

type BaseParticle struct {
	Position   Vector
	Size       uint32
	Color      color.RGBA
	HasUpdated bool
}

func (p *BaseParticle) Update(world *World, dt float64) {}

func (p *BaseParticle) Draw(screen *ebiten.Image) {}

func (p *BaseParticle) Reset() {}

func (p *BaseParticle) GetPosition() Vector {
	return p.Position
}

func (p *BaseParticle) SetPosition(v Vector) {
	p.Position = v
}

//// AirParticle ////

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

func NewAirParticle(x, y int) *AirParticle {
	p := &AirParticle{
		BaseParticle: BaseParticle{
			Position: Vector{X: float64(x), Y: float64(y)},
			Size:     1,
			Color:    color.RGBA{R: 0, G: 0, B: 0, A: 0},
		},
	}

	return p
}

//// SandParticle ////

const (
	sandInitialVelocity = 0.5
)

type SandParticle struct {
	BaseParticle
	Velocity float32
}

func (sp *SandParticle) Update(world *World, dt float64) {
	if sp.HasUpdated || sp.Position.Y+1 >= screenHeight {
		return
	}

	nextPosition := Vector{
		X: sp.Position.X,
		Y: sp.Position.Y + float64(sp.Velocity),
	}

	if nextPosition.Y >= screenHeight {
		nextPosition = Vector{
			X: sp.Position.X,
			Y: screenHeight - 1,
		}
	}

	nextParticle := world.Particles[uint32(nextPosition.X)][uint32(nextPosition.Y)]

	world.SwapPosition(sp, nextParticle)

	sp.Velocity += gravity
	sp.HasUpdated = true

}

func (sp *SandParticle) Draw(screen *ebiten.Image) {
	screen.Set(int(sp.Position.X), int(sp.Position.Y), sp.Color)
}

func (sp *SandParticle) Reset() {
	sp.HasUpdated = false
}

func NewSandParticle(x, y float64) *SandParticle {
	p := &SandParticle{
		BaseParticle: BaseParticle{
			Position: Vector{X: x, Y: y},
			Size:     1,
			Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
		},
		Velocity: sandInitialVelocity,
	}

	return p
}
