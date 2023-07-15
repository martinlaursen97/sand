package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/core"
)

type Game struct {
	world     *core.World
	brushSize int
}

func (g *Game) Update() error {
	dt := 1.0 / ebiten.ActualTPS()

	g.world.Update(dt)

	cursorPositionX, cursorPositionY := ebiten.CursorPosition()

	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	spacePressed := ebiten.IsKeyPressed(ebiten.KeySpace)
	increaseBrushSize := ebiten.IsKeyPressed(ebiten.Key1)
	decreaseBrushSize := ebiten.IsKeyPressed(ebiten.Key2)

	if mouseClicked {
		g.world.DrawWithBrush(g.brushSize, cursorPositionX, cursorPositionY)
	}

	if spacePressed {
		g.world.Clear()
	}

	if increaseBrushSize {
		g.brushSize += 2
	}

	if decreaseBrushSize {
		if g.brushSize > 1 {
			g.brushSize -= 2
		}
	}

	g.world.Reset()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.ScreenWidth, config.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(config.ScreenWidth*2, config.ScreenHeight*2)
	ebiten.SetWindowTitle("Sand")

	game := &Game{}
	game.world = core.NewWorld()
	game.brushSize = 5

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
