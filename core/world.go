package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

type World struct {
	Particles [][]Particle
}

func NewWorld() *World {
	particleGrid := make([][]Particle, config.ScreenWidth)
	for i := range particleGrid {
		particleGrid[i] = make([]Particle, config.ScreenHeight)
	}

	return &World{
		Particles: particleGrid,
	}
}

func (w *World) Clear() {
	for i := 0; i < config.ScreenWidth; i++ {
		for j := 0; j < config.ScreenHeight; j++ {
			w.Particles[i][j] = nil
		}
	}
}

func (w *World) Update(dt float64) {
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				particle.Update(w, dt)
			}
		}
	}
}

func (w *World) Draw(screen *ebiten.Image) {
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				particle.Draw(screen)
			}
		}
	}
}

func (w *World) Reset() {
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				particle.Reset()
			}
		}
	}
}

func (w *World) IsEmpty(x, y float64) bool {
	return w.Particles[int(x)][int(y)] == nil
}

func (w *World) SwapPosition(p1, p2 MoveableParticle) {
	p1Pos := p1.GetPosition()
	p2Pos := p2.GetPosition()

	w.Particles[int(p1Pos.X)][int(p1Pos.Y)] = p2
	w.Particles[int(p2Pos.X)][int(p2Pos.Y)] = p1

	p1.SetPosition(p2Pos)
	p2.SetPosition(p1Pos)
}

func (w *World) GetParticleAt(x, y float64) Particle {
	return w.Particles[int(x)][int(y)]
}

func (w *World) InsertParticle(p Particle) {
	position := p.GetPosition()
	x, y := int(position.X), int(position.Y)
	if utils.WithinBoundsInt(x, y) {
		w.Particles[x][y] = p
	}
}

// To be used with TraversePathAndApplyFunc
func (w *World) DrawWithBrush(args ...any) any {
	fx := args[0].(float64)
	fy := args[1].(float64)

	x, y := int(fx), int(fy)

	size := args[2].(int)
	particleNum := args[3].(int)

	for i := -size / 2; i <= size/2; i++ {
		if utils.WithinBoundsInt(x+i, y) {
			if particleNum == 1 {
				w.InsertParticle(CreateSandParticle(x+i, y))
			} else if particleNum == 2 {
				w.InsertParticle(CreateWallParticle(x+i, y))
			} else if particleNum == 3 {
				w.Particles[x+i][y] = nil
			}
		}
	}

	return nil
}

func (w *World) MoveParticle(p MoveableParticle, nextPosition maths.Vector) {
	// Ensure that the particle is not overwritten with nil
	if int(nextPosition.X) == int(p.GetPosition().X) &&
		int(nextPosition.Y) == int(p.GetPosition().Y) {
		return
	}

	w.Particles[int(nextPosition.X)][int(nextPosition.Y)] = p
	w.Particles[int(p.GetPosition().X)][int(p.GetPosition().Y)] = nil
	p.SetPosition(nextPosition)

}

func (w *World) GetParticleCount() int {
	count := 0
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				count++
			}
		}
	}
	return count
}

// Get sand particle count
func (w *World) GetSandParticleCount() int {
	count := 0
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				if _, ok := particle.(*SandParticle); ok {
					count++
				}
			}
		}
	}
	return count
}

// Get wall particle count
func (w *World) GetWallParticleCount() int {
	count := 0
	for _, row := range w.Particles {
		for _, particle := range row {
			if particle != nil {
				if _, ok := particle.(*WallParticle); ok {
					count++
				}
			}
		}
	}
	return count
}

func (w *World) PrintTransposed() {
	for i := 0; i < config.ScreenHeight; i++ {
		for j := 0; j < config.ScreenWidth; j++ {
			if w.Particles[j][i] != nil {
				print("1 ")
			} else {
				print("0 ")
			}
		}
		print("\n")
	}
	print("\n")
}
