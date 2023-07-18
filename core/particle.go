package core

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

type Particle interface {
	Update(world *World, dt float64)
	Draw(screen *ebiten.Image)
	Reset()
	GetPosition() maths.Vector
}

type ImmovableSolid interface {
	Particle
}

type MoveableSolid interface {
	MoveableParticle
}

type Moveable struct {
	BaseParticle
	Velocity  maths.Vector
	IsFalling bool
	Density   float64
}

type Immovable struct {
	BaseParticle
}
type BaseParticle struct {
	Position   maths.Vector
	Size       uint32
	Color      color.RGBA
	HasUpdated bool
}

type LiquidSolid interface {
	MoveableParticle
}

type Liquid struct {
	Moveable
}

func (p *BaseParticle) GetPosition() maths.Vector {
	return p.Position
}

func (p *BaseParticle) SetPosition(v maths.Vector) {
	p.Position = v
}

func (p *BaseParticle) Reset() {
	p.HasUpdated = false
}

func (p *BaseParticle) Draw(screen *ebiten.Image) {
	screen.Set(int(p.Position.X), int(p.Position.Y), p.Color)
}

type MoveableParticle interface {
	Particle
	GetPointBeforeCollision(world *World) (maths.Vector, bool)
	checkCollisionsAndGetNextPosition(world *World) (maths.Vector, bool)
	ResetVelocity()
	SetPosition(maths.Vector)
	GetIsFalling() bool
	GetDensity() float64
}

func (m *Moveable) GetIsFalling() bool {
	return m.IsFalling
}

func (m *Moveable) GetPointBeforeCollision(world *World) (maths.Vector, bool) {
	nextPos := maths.Vector{
		X: m.Position.X + m.Velocity.X,
		Y: m.Position.Y + m.Velocity.Y,
	}

	vx := nextPos.X - m.Position.X
	vy := nextPos.Y - m.Position.Y
	length := math.Sqrt(vx*vx + vy*vy)

	xIncrement, yIncrement := vx/length, vy/length
	numPoints := int(length)

	if numPoints == 0 {
		return nextPos, false
	}

	prevX, prevY := m.Position.X, m.Position.Y

	for i := 1; i <= numPoints; i++ {
		dx := m.Position.X + xIncrement*float64(i)
		dy := m.Position.Y + yIncrement*float64(i)

		dy = utils.RoundYCoordinate(dy)
		dx, dy = utils.CheckBounds(dx, dy)

		// Hit something, return the position before the collision
		if !world.IsEmpty(dx, dy) {
			return maths.Vector{X: prevX, Y: prevY}, true
		}

		prevX, prevY = dx, dy
	}

	nextPos.X, nextPos.Y = utils.CheckBounds(nextPos.X, nextPos.Y)

	// Did not hit anything, return the next position
	return nextPos, false
}
