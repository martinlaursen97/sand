package core

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Particle interface {
	Update(world *World, dt float64)
	Draw(screen *ebiten.Image)
	Reset()
	GetPosition() Vector
	SetPosition(Vector)
}

type SolidMoveable interface {
	Particle
	GetNextPosition(world *World, nextPos Vector) Vector
	ResetVelocity()
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
	sandInitialVelocityX = 0.5
	sandInitialVelocityY = 0.15
)

type SandParticle struct {
	BaseParticle
	Velocity Vector
}

func (sp *SandParticle) Update(world *World, dt float64) {
	if sp.HasUpdated || sp.Position.Y+1 >= worldHeight {
		return
	}

	sp.Velocity.Y += gravity
	sp.Velocity.Y = math.Min(sp.Velocity.Y, maxVelocity)

	nextPosition, collided := sp.CheckCollisionsBelowAndSides(world)

	if !collided {
		nextPosition = sp.GetNextPosition(world)
	}

	nextParticle := world.GetParticleAt(uint32(nextPosition.X), uint32(nextPosition.Y))

	world.SwapPosition(sp, nextParticle)
	sp.Position = nextPosition

	sp.HasUpdated = true
}

func (sp *SandParticle) CheckCollisionsBelowAndSides(world *World) (Vector, bool) {
	if !world.IsEmpty(uint32(sp.Position.X), uint32(sp.Position.Y+1)) {
		sp.ResetVelocity()
		if withinBounds(uint32(sp.Position.X-1), uint32(sp.Position.Y+1)) &&
			world.IsEmpty(uint32(sp.Position.X-1), uint32(sp.Position.Y+1)) {

			return Vector{X: sp.Position.X - 1, Y: sp.Position.Y + 1}, true
		} else if withinBounds(uint32(sp.Position.X+1), uint32(sp.Position.Y+1)) &&
			world.IsEmpty(uint32(sp.Position.X+1), uint32(sp.Position.Y+1)) {

			return Vector{X: sp.Position.X + 1, Y: sp.Position.Y + 1}, true
		}
	}

	return Vector{}, false
}

func (sp *SandParticle) Draw(screen *ebiten.Image) {
	screen.Set(int(sp.Position.X), int(sp.Position.Y), sp.Color)
}

func (sp *SandParticle) Reset() {
	sp.HasUpdated = false
}

func (sp *SandParticle) GetNextPosition(world *World) Vector {

	nextPos := Vector{
		X: sp.Position.X + sp.Velocity.X,
		Y: sp.Position.Y + sp.Velocity.Y,
	}

	vx := nextPos.X - sp.Position.X
	vy := nextPos.Y - sp.Position.Y

	dist := math.Sqrt(vx*vx + vy*vy)

	xIncrement, yIncrement := vx/dist, vy/dist

	numPoints := int(dist)

	prevX, prevY := sp.Position.X, sp.Position.Y

	for i := 1; i <= numPoints; i++ {
		x := sp.Position.X + xIncrement*float64(i)
		y := sp.Position.Y + yIncrement*float64(i)

		y = roundYCoordinate(y)
		x, y = checkBounds(x, y)

		if _, ok := world.GetParticleAt(uint32(x), uint32(y)).(*AirParticle); !ok {
			// Hit something, return the position before the collision
			sp.ResetVelocity()

			return Vector{X: prevX, Y: prevY}
		}

		prevX, prevY = x, y

	}

	// Did not hit anything, return the next position
	nextPos.X, nextPos.Y = checkBounds(nextPos.X, nextPos.Y)

	return nextPos
}

func (sp *SandParticle) ResetVelocity() {
	sp.Velocity.X = 0
	sp.Velocity.Y = 0
}

func NewSandParticle(x, y float64) *SandParticle {
	p := &SandParticle{
		BaseParticle: BaseParticle{
			Position: Vector{X: x, Y: y},
			Size:     1,
			Color:    color.RGBA{R: 255, G: 255, B: 0, A: 255},
		},
		Velocity: Vector{X: -sandInitialVelocityX + rand.Float64(), Y: sandInitialVelocityY},
	}

	return p
}
