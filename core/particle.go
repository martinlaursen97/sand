package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/maths"
)

type Particle interface {
	Update(world *World, dt float64)
	Draw(screen *ebiten.Image)
	Reset()
	GetPosition() maths.Vector
	SetPosition(maths.Vector)
}

type SolidMoveable interface {
	Particle
	GetNextPosition(world *World, nextPos maths.Vector) maths.Vector
	ResetVelocity()
}
