package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/martinlaursen97/sand/maths"
)

type BaseParticle struct {
	Position   maths.Vector
	Size       uint32
	Color      color.RGBA
	HasUpdated bool
}

func (p *BaseParticle) Update(world *World, dt float64) {}

func (p *BaseParticle) Draw(screen *ebiten.Image) {}

func (p *BaseParticle) Reset() {}

func (p *BaseParticle) GetPosition() maths.Vector {
	return p.Position
}

func (p *BaseParticle) SetPosition(v maths.Vector) {
	p.Position = v
}
