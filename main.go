package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//game struct -> holds all the data we need globally for the game
// and will also be the structure we will meet the interface
// to bootstrap Ebiten.

type Game struct{}

// constructor for the game Object and init the data
func NewGame() *Game {
	g := &Game{}
	return g
}

// callbacks for the Updating systems , drawing and layout
// update will be called with each tic
func (g *Game) Update() error {
	return nil
}

// drawing is calling each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {

}

// the layout will return the screen dimensions
func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}

func main() {
	g := NewGame()
	ebiten.SetWindowResizingMode(1)
	ebiten.SetWindowTitle("Rouge")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
