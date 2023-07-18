package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/martinlaursen97/sand/config"
	"github.com/martinlaursen97/sand/core"
	"github.com/martinlaursen97/sand/maths"
)

type Game struct {
	world         *core.World
	brushSize     int
	paused        bool
	particleNum   int
	mouseClicked  bool
	prevCursor    maths.Vector
	currentCursor maths.Vector
}

func (g *Game) Update() error {
	dt := 1.0 / ebiten.ActualTPS()

	g.world.Update(dt)

	cursorPositionX, cursorPositionY := ebiten.CursorPosition()

	mouseClicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	spacePressed := ebiten.IsKeyPressed(ebiten.KeySpace)
	pickSand := ebiten.IsKeyPressed(ebiten.Key1)
	pickWater := ebiten.IsKeyPressed(ebiten.Key2)
	pickWall := ebiten.IsKeyPressed(ebiten.Key3)
	pickEraser := ebiten.IsKeyPressed(ebiten.Key4)

	_, dy := ebiten.Wheel()

	if spacePressed {
		g.world.Clear()
	}

	if dy > 0 {
		g.brushSize += config.BrushSizeIncrement
	}

	if dy < 0 {
		if g.brushSize-config.BrushSizeIncrement > 0 {
			g.brushSize -= config.BrushSizeIncrement
		} else {
			g.brushSize = 1
		}
	}

	if pickSand {
		g.particleNum = 1
	}

	if pickWater {
		g.particleNum = 2
	}

	if pickWall {
		g.particleNum = 3
	}

	if pickEraser {
		g.particleNum = 4
	}

	if mouseClicked {
		g.prevCursor.X = g.currentCursor.X
		g.prevCursor.Y = g.currentCursor.Y

		g.currentCursor.X = float64(cursorPositionX)
		g.currentCursor.Y = float64(cursorPositionY)

		// Draw a line between the previous and current cursor position
		core.TraversePathAndApplyFunc(
			g.prevCursor,
			g.currentCursor,
			core.FunctionInput{
				Func: g.world.DrawWithBrush,
				Args: []interface{}{g.brushSize, g.particleNum},
			},
		)
	} else {
		g.currentCursor.X = float64(cursorPositionX)
		g.currentCursor.Y = float64(cursorPositionY)
	}

	g.world.Reset()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS())))

	particleCount := g.world.GetParticleCount()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Particles: %d", particleCount), 0, 15)

	sandCount := g.world.GetSandParticleCount()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sand: %d", sandCount), 0, 30)

	wallCount := g.world.GetWallParticleCount()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Walls: %d", wallCount), 0, 45)
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
	game.paused = false
	game.particleNum = 2

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
