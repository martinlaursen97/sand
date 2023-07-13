package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	gravity      = 0.1
	maxVelocity  = 2.5
)

type World struct {
	Particles [][]Particle
}

func NewWorld() *World {
	particleGrid := make([][]Particle, screenWidth)
	for i := range particleGrid {
		particleGrid[i] = make([]Particle, screenHeight)
		for j := range particleGrid[i] {
			particleGrid[i][j] = NewAirParticle(i, j)
		}
	}

	return &World{
		Particles: particleGrid,
	}
}

func (w *World) Update(dt float64) {
	for _, row := range w.Particles {
		for _, particle := range row {
			particle.Update(w, dt)
		}
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, row := range w.Particles {
		for _, particle := range row {
			particle.Draw(screen)
		}
	}
}

func (w *World) Reset() {
	for _, row := range w.Particles {
		for _, particle := range row {
			particle.Reset()
		}
	}
}

func (w *World) SwapPosition(p1, p2 Particle) {
	p1Pos := p1.GetPosition()
	p2Pos := p2.GetPosition()

	p1.SetPosition(p2Pos)
	p2.SetPosition(p1Pos)

	w.Particles[uint32(p1Pos.X)][uint32(p1Pos.Y)] = p1
	w.Particles[uint32(p2Pos.X)][uint32(p2Pos.Y)] = p2
}

func (w *World) GetParticleAt(x, y uint32) Particle {
	return w.Particles[x][y]
}

func (w *World) InsertParticle(p Particle) {
	w.Particles[uint32(p.GetPosition().X)][uint32(p.GetPosition().Y)] = p
}

func (w *World) MoveParticle(p Particle, x, y uint32) {
	w.Particles[x][y] = p
}

// Debugging

func (w *World) GetAirParticleCount() int {
	count := 0

	for _, row := range w.Particles {
		for _, particle := range row {
			if _, ok := particle.(*AirParticle); ok {
				count++
			}
		}
	}

	return count
}

func (w *World) GetSandParticleCount() int {
	count := 0

	for _, row := range w.Particles {
		for _, particle := range row {
			if _, ok := particle.(*SandParticle); ok {
				count++
			}
		}
	}

	return count
}
