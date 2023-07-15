package core

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

const (
	sandInitialVelocityX = 0.5
	sandInitialVelocityY = 0.15
)

type SandParticle struct {
	BaseParticle
	Velocity  maths.Vector
	IsFalling bool
}

func (sp *SandParticle) Update(world *World, dt float64) {
	if sp.HasUpdated || sp.Position.Y+1 >= config.ScreenHeight {
		return
	}

	sp.Velocity.Y += config.Gravity
	sp.Velocity.Y = math.Min(sp.Velocity.Y, config.MaxVelocity)

	nextPosition, collided := sp.CheckCollisionsBelowAndSides(world)

	if !collided {
		nextPosition = sp.GetNextPosition(world)
	}

	nextPosition.X, nextPosition.Y = utils.CheckBounds(nextPosition.X, nextPosition.Y)

	nextParticle := world.GetParticleAt(int(nextPosition.X), int(nextPosition.Y))

	world.SwapPosition(sp, nextParticle)
	sp.Position = nextPosition

	sp.HasUpdated = true
}

func (sp *SandParticle) CheckCollisionsBelowAndSides(world *World) (maths.Vector, bool) {
	if !world.IsEmpty(int(sp.Position.X), int(sp.Position.Y+1)) {
		sp.ResetVelocity()

		order := rand.Intn(2)

		if order == 0 {
			if utils.WithinBounds(int(sp.Position.X-1), int(sp.Position.Y+1)) &&
				world.IsEmpty(int(sp.Position.X-1), int(sp.Position.Y+1)) {

				return maths.Vector{X: sp.Position.X - 1, Y: sp.Position.Y + 1}, true
			} else if utils.WithinBounds(int(sp.Position.X+1), int(sp.Position.Y+1)) &&
				world.IsEmpty(int(sp.Position.X+1), int(sp.Position.Y+1)) {

				return maths.Vector{X: sp.Position.X + 1, Y: sp.Position.Y + 1}, true
			}
		} else {
			if utils.WithinBounds(int(sp.Position.X+1), int(sp.Position.Y+1)) &&
				world.IsEmpty(int(sp.Position.X+1), int(sp.Position.Y+1)) {

				return maths.Vector{X: sp.Position.X + 1, Y: sp.Position.Y + 1}, true
			} else if utils.WithinBounds(int(sp.Position.X-1), int(sp.Position.Y+1)) &&
				world.IsEmpty(int(sp.Position.X-1), int(sp.Position.Y+1)) {

				return maths.Vector{X: sp.Position.X - 1, Y: sp.Position.Y + 1}, true
			}
		}
	}

	return maths.Vector{}, false
}

func (sp *SandParticle) Draw(screen *ebiten.Image) {
	screen.Set(int(sp.Position.X), int(sp.Position.Y), sp.Color)
}

func (sp *SandParticle) Reset() {
	sp.HasUpdated = false
}

func (sp *SandParticle) GetNextPosition(world *World) maths.Vector {

	nextPos := maths.Vector{
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

		y = utils.RoundYCoordinate(y)
		x, y = utils.CheckBounds(x, y)

		if _, ok := world.GetParticleAt(int(x), int(y)).(*AirParticle); !ok {
			// Hit something, return the position before the collision

			sp.ResetVelocity()

			return maths.Vector{X: prevX, Y: prevY}
		}

		prevX, prevY = x, y

	}

	// Did not hit anything, return the next position
	nextPos.X, nextPos.Y = utils.CheckBounds(nextPos.X, nextPos.Y)

	return nextPos
}

func (sp *SandParticle) ResetVelocity() {
	sp.Velocity.X = 0
	sp.Velocity.Y = 0
}

func CreateSandParticle(x, y int) *SandParticle {
	p := &SandParticle{
		BaseParticle: BaseParticle{
			Position: maths.Vector{X: float64(x), Y: float64(y)},
			Size:     config.ParticleSize,
			Color:    randomSandColor(),
		},
		Velocity: maths.Vector{X: -sandInitialVelocityX + rand.Float64(), Y: sandInitialVelocityY},
	}

	return p
}

func randomSandColor() color.RGBA {
	return color.RGBA{
		R: utils.RandomUnsignedByteInRange(180, 210),
		G: utils.RandomUnsignedByteInRange(160, 190),
		B: utils.RandomUnsignedByteInRange(110, 140),
		A: 0,
	}
}
