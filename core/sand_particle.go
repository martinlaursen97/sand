package core

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

const (
	sandInitialVelocityX = 0.5
	sandInitialVelocityY = 0.0
)

type SandParticle struct {
	Moveable
}

type Sand interface {
	CreateSandParticle(x, y int) *SandParticle
}

func (sp *SandParticle) Update(world *World, dt float64) {
	if sp.HasUpdated || sp.Position.Y+1 >= config.ScreenHeight {
		return
	}

	sp.Velocity.Y += config.Gravity
	sp.Velocity.Y = math.Min(sp.Velocity.Y, config.MaxVelocity)

	nextPos, collided := sp.checkCollisionsAndGetNextPosition(world)

	if !collided {
		nextPos = sp.getNextPosition(world)
	}

	sp.IsFalling = nextPos.X != sp.Position.X || nextPos.Y != sp.Position.Y

	if sp.IsFalling {
		world.MoveParticle(sp, nextPos)
	}

	sp.HasUpdated = true
}

func (sp *SandParticle) checkCollisionsAndGetNextPosition(world *World) (maths.Vector, bool) {

	if sp.Velocity.Y < 0 {
		return sp.Position, false
	}

	belowIsEmpty := world.IsEmpty(sp.Position.X, sp.Position.Y+1)

	if !belowIsEmpty {

		//sp.ResetVelocity()

		if sp.IsFalling {
			return sp.Position, false
		}

		leftPos := maths.Vector{X: sp.Position.X - 1, Y: sp.Position.Y + 1}
		rightPos := maths.Vector{X: sp.Position.X + 1, Y: sp.Position.Y + 1}

		leftIsEmpty := utils.WithinBounds(leftPos.X, leftPos.Y) &&
			world.IsEmpty(leftPos.X, leftPos.Y)
		rightIsEmpty := utils.WithinBounds(rightPos.X, rightPos.Y) &&
			world.IsEmpty(rightPos.X, rightPos.Y)

		order := rand.Intn(2)

		if order == 0 {
			if leftIsEmpty {
				return leftPos, true
			} else if rightIsEmpty {
				return rightPos, true
			}
		} else {
			if rightIsEmpty {
				return rightPos, true
			} else if leftIsEmpty {
				return leftPos, true
			}
		}
	}

	sp.IsFalling = belowIsEmpty

	return sp.Position, false
}

func (sp *SandParticle) getNextPosition(world *World) maths.Vector {

	nextPos := maths.Vector{
		X: sp.Position.X + sp.Velocity.X,
		Y: sp.Position.Y + sp.Velocity.Y,
	}

	vx := nextPos.X - sp.Position.X
	vy := nextPos.Y - sp.Position.Y

	length := math.Sqrt(vx*vx + vy*vy)

	xIncrement, yIncrement := vx/length, vy/length

	numPoints := int(length)

	prevX, prevY := sp.Position.X, sp.Position.Y

	for i := 1; i <= numPoints; i++ {
		dx := sp.Position.X + xIncrement*float64(i)
		dy := sp.Position.Y + yIncrement*float64(i)

		dy = utils.RoundYCoordinate(dy)
		dx, dy = utils.CheckBounds(dx, dy)

		// Hit something, return the position before the collision
		if !world.IsEmpty(dx, dy) {

			sp.ResetVelocity()
			return maths.Vector{X: prevX, Y: prevY}
		}

		prevX, prevY = dx, dy

	}

	nextPos.X, nextPos.Y = utils.CheckBounds(nextPos.X, nextPos.Y)

	// Did not hit anything, return the next position
	return nextPos
}

func (sp *SandParticle) ResetVelocity() {
	sp.Velocity.X = getRandomVelocityX()
	sp.Velocity.Y = sandInitialVelocityY
}

func CreateSandParticle(x, y int) *SandParticle {

	p := &SandParticle{
		Moveable: Moveable{
			BaseParticle: BaseParticle{
				Position: maths.Vector{X: float64(x), Y: float64(y)},
				Size:     config.ParticleSize,
				Color:    randomSandColor(),
			},
			Velocity: maths.Vector{
				X: getRandomVelocityX(),
				Y: sandInitialVelocityY,
			},
			IsFalling: true,
		},
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

func getRandomVelocityX() float64 {
	return utils.RandomFloatInRange(-sandInitialVelocityX, sandInitialVelocityX)
}

// func randomSandColor() color.RGBA {
// 	return color.RGBA{
// 		R: utils.RandomUnsignedByteInRange(50, 80),
// 		G: utils.RandomUnsignedByteInRange(45, 70),
// 		B: utils.RandomUnsignedByteInRange(40, 55),
// 		A: 0,
// 	}
// }
