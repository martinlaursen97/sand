package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	worldWidth  = 320
	worldHeight = 240
	gravity     = 0.15
	maxVelocity = 2.5
)

type World struct {
	Particles [][]Particle
}

func GetWorldWidth() int {
	return worldWidth
}

func GetWorldHeight() int {
	return worldHeight
}

func NewWorld() *World {
	particleGrid := make([][]Particle, worldWidth)
	for i := range particleGrid {
		particleGrid[i] = make([]Particle, worldHeight)
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

	w.Particles[uint32(p1Pos.X)][uint32(p1Pos.Y)] = p2
	w.Particles[uint32(p2Pos.X)][uint32(p2Pos.Y)] = p1

	p1.SetPosition(p2Pos)
	p2.SetPosition(p1Pos)
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

func (w *World) PrintGridT() {
	print("SKIP\n")
	for i := 0; i < worldHeight; i++ {
		for j := 0; j < worldWidth; j++ {
			switch w.Particles[j][i].(type) {
			case *AirParticle:
				print("A ")
			case *SandParticle:
				print("S ")
			}
		}
		println()
	}
}

func (w *World) WithinBounds(x, y uint32) bool {
	return x < worldWidth && y < worldHeight && x >= 0 && y >= 0
}
