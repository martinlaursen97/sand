package core

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

type WaterParticle struct {
	Liquid
}

type Water interface {
	CreateWaterParticle(x, y int) *WaterParticle
}

const (
	waterDensity          = 0.5
	waterInitialVelocityX = 0.5
	waterInitialVelocityY = 0
	dispersionRate        = 7
)

func randomWaterColor() color.RGBA {
	return color.RGBA{
		R: utils.RandomUnsignedByteInRange(10, 30),
		G: utils.RandomUnsignedByteInRange(60, 110),
		B: utils.RandomUnsignedByteInRange(140, 170),
		A: 0,
	}
}

func CreateWaterParticle(x, y int) *WaterParticle {
	return &WaterParticle{
		Liquid{
			Moveable: Moveable{
				BaseParticle: BaseParticle{
					Position: maths.Vector{
						X: float64(x),
						Y: float64(y),
					},
					Size:       1,
					Color:      randomWaterColor(),
					HasUpdated: false,
				},
				Velocity: maths.Vector{
					X: utils.RandomFloatInRange(-waterInitialVelocityX, waterInitialVelocityX),
					Y: waterInitialVelocityY,
				},
				IsFalling: true,
				Density:   waterDensity,
			},
		},
	}
}

func (wp *WaterParticle) checkCollisionsAndGetNextPosition(world *World) (maths.Vector, bool) {

	if wp.Velocity.Y < 0 || wp.Position.Y+1 >= config.ScreenHeight {
		return wp.Position, false
	}

	belowIsEmpty := world.IsEmpty(wp.Position.X, wp.Position.Y+1)

	if !belowIsEmpty {

		if wp.IsFalling {
			wp.IsFalling = false
			return wp.Position, false
		}

		bottomLeftPos := maths.Vector{X: wp.Position.X - 1, Y: wp.Position.Y + 1}
		bottmRightPos := maths.Vector{X: wp.Position.X + 1, Y: wp.Position.Y + 1}

		leftPos := maths.Vector{X: wp.Position.X - 1, Y: wp.Position.Y}
		rightPos := maths.Vector{X: wp.Position.X + 1, Y: wp.Position.Y}

		bottomLeftIsEmpty := utils.WithinBounds(bottomLeftPos.X, bottmRightPos.Y) &&
			world.IsEmpty(bottomLeftPos.X, bottmRightPos.Y)
		bottomRightIsEmpty := utils.WithinBounds(bottomLeftPos.X, bottmRightPos.Y) &&
			world.IsEmpty(bottomLeftPos.X, bottmRightPos.Y)

		leftIsEmpty := utils.WithinBounds(leftPos.X, leftPos.Y) &&
			world.IsEmpty(leftPos.X, leftPos.Y)
		rightIsEmpty := utils.WithinBounds(rightPos.X, rightPos.Y) &&
			world.IsEmpty(rightPos.X, rightPos.Y)

		order := rand.Intn(2)

		if order == 0 {
			if bottomLeftIsEmpty {
				return bottomLeftPos, true
			} else if bottomRightIsEmpty {
				return bottmRightPos, true
			} else if leftIsEmpty {
				return leftPos, true
			} else if rightIsEmpty {
				return rightPos, true
			}
		} else {
			if bottomRightIsEmpty {
				return bottmRightPos, true
			} else if bottomLeftIsEmpty {
				return bottomLeftPos, true
			} else if rightIsEmpty {
				return rightPos, true
			} else if leftIsEmpty {
				return leftPos, true
			}

		}
	}

	wp.IsFalling = true

	return wp.Position, false
}

func (wp *WaterParticle) ResetVelocity() {
	wp.Velocity.X = utils.RandomFloatInRange(-sandInitialVelocityX, sandInitialVelocityX)
	wp.Velocity.Y = sandInitialVelocityY
}

func (wp *WaterParticle) GetDensity() float64 {
	return wp.Density
}

func (wp *WaterParticle) Update(world *World, dt float64) {
	if wp.HasUpdated {
		return
	}

	if rand.Intn(50) == 0 {
		wp.Color = randomWaterColor()
	}

	wp.Velocity.Y += config.Gravity
	wp.Velocity.Y = math.Min(wp.Velocity.Y, config.MaxVelocity)

	nextPos, collided := wp.checkCollisionsAndGetNextPosition(world)

	for i := 0; i < dispersionRate; i++ {
		nextPos, collided = wp.checkCollisionsAndGetNextPosition(world)

		world.MoveParticle(wp, nextPos)
	}

	if !collided {
		nextPos, collided = wp.GetPointBeforeCollision(world)
	}

	if collided {
		wp.ResetVelocity()
	}

	wp.IsFalling = (nextPos.X != wp.Position.X || nextPos.Y != wp.Position.Y)

	if wp.IsFalling {
		world.MoveParticle(wp, nextPos)
	}

	wp.HasUpdated = true
}
