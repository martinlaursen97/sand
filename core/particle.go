package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/maths"
)

type Particle interface {
	Update(world *World, dt float64)
	Draw(screen *ebiten.Image)
	Reset()
	GetPosition() maths.Vector
}

type MoveableParticle interface {
	Particle
	getNextPosition(world *World) maths.Vector
	checkCollisionsAndGetNextPosition(world *World) (maths.Vector, bool)
	ResetVelocity()
	SetPosition(maths.Vector)
}

type MoveableSolid interface {
	MoveableParticle
}

type BaseParticle struct {
	Position   maths.Vector
	Size       uint32
	Color      color.RGBA
	HasUpdated bool
}

func (p *BaseParticle) GetPosition() maths.Vector {
	return p.Position
}

func (p *BaseParticle) SetPosition(v maths.Vector) {
	p.Position = v
}
