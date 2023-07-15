package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/config"

	"github.com/martinlaursen97/sand/utils"
)

type World struct {
	Particles [][]Particle
}

func NewWorld() *World {
	particleGrid := make([][]Particle, config.ScreenWidth)
	for i := range particleGrid {
		particleGrid[i] = make([]Particle, config.ScreenHeight)
		for j := range particleGrid[i] {
			particleGrid[i][j] = CreateAirParticle(i, j)
		}
	}

	return &World{
		Particles: particleGrid,
	}
}

func (w *World) Clear() {
	for i := 0; i < config.ScreenWidth; i++ {
		for j := 0; j < config.ScreenHeight; j++ {
			w.Particles[i][j] = CreateAirParticle(i, j)
		}
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

func (w *World) IsEmpty(x, y int) bool {
	if _, ok := w.Particles[x][y].(*AirParticle); ok {
		return true
	}

	return false
}

func (w *World) SwapPosition(p1, p2 Particle) {
	p1Pos := p1.GetPosition()
	p2Pos := p2.GetPosition()

	w.Particles[int(p1Pos.X)][int(p1Pos.Y)] = p2
	w.Particles[int(p2Pos.X)][int(p2Pos.Y)] = p1

	p1.SetPosition(p2Pos)
	p2.SetPosition(p1Pos)
}

func (w *World) GetParticleAt(x, y int) Particle {
	return w.Particles[x][y]
}

func (w *World) InsertParticle(p Particle) {
	w.Particles[int(p.GetPosition().X)][int(p.GetPosition().Y)] = p
}

func (w *World) DrawWithBrush(size, x, y int) {
	for i := -size / 2; i <= size/2; i += 2 {
		if utils.WithinBounds(x+i, y) {
			w.InsertParticle(
				CreateSandParticle(x+i, y),
			)
		}
	}
}
