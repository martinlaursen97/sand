package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	worldWidth  = 320
	worldHeight = 240
	gravity     = 0.15
	maxVelocity = 10
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

func (w *World) IsEmpty(x, y uint32) bool {
	if _, ok := w.Particles[x][y].(*AirParticle); ok {
		return true
	}

	return false
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

func (w *World) DrawWithBrush(size int, x, y int) {

	for i := -size / 2; i <= size/2; i += 2 {
		if withinBounds(uint32(x+i), uint32(y)) {
			w.InsertParticle(NewSandParticle(float64(x+i), float64(y)))
		}
	}
}
