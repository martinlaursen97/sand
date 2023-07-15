package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/martinlaursen97/sand/core"
)

type Game struct {
	world *core.World
}

func (g *Game) Update() error {
	dt := 1.0 / ebiten.ActualTPS()

	g.world.Update(dt)

	cursorPositionX, cursorPositionY := ebiten.CursorPosition()
	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	withinBounds := g.world.WithinBounds(
		uint32(cursorPositionX),
		uint32(cursorPositionY),
	)

	if mouseClicked && withinBounds {
		sp := core.NewSandParticle(
			float64(cursorPositionX),
			float64(cursorPositionY),
		)

		g.world.InsertParticle(sp)

		sp1 := core.NewSandParticle(
			float64(cursorPositionX-2),
			float64(cursorPositionY),
		)

		g.world.InsertParticle(sp1)

		sp2 := core.NewSandParticle(
			float64(cursorPositionX+4),
			float64(cursorPositionY),
		)

		g.world.InsertParticle(sp2)
	}

	g.world.Reset()

	// For debugging
	if ebiten.IsKeyPressed(ebiten.Key1) {
		fmt.Println(g.world.GetAirParticleCount())
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		fmt.Println(g.world.GetSandParticleCount())
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return core.GetWorldWidth(), core.GetWorldHeight()
}

func main() {
	ebiten.SetWindowSize(core.GetWorldWidth()*2, core.GetWorldHeight()*2)
	ebiten.SetWindowTitle("Sand")

	game := &Game{}
	game.world = core.NewWorld()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
