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

	if mouseClicked {
		g.world.DrawWithBrush(5, cursorPositionX, cursorPositionY)
	}

	g.world.Reset()

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
