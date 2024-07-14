package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game struct willl hold all the data the entire game will ned.
type Game struct {
	Map GameMap
}

//create a new game object and initialize the data

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g
}

// callbacks for update, draw and layout
func (g *Game) Update() error {
	return nil
}

// draw similarly is called draw cycle and is where we will blit
func (g *Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

//layout will return the screen dimensions

func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	//main loop to render
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Rouge")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
