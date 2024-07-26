package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

//game holds all the data the entire game will need

type Game struct{}

//constructor for the new game

func NewGame() *Game {
	g := &Game{}
	return g
}

// update the game on each of the tic
func (g *Game) Update() error {
	return nil
}

// draw will be called on each draw cycle and is where we will blip
func (g *Game) Draw(screen *ebiten.Image) {

}

// layout will return the screen dimensions
func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
}

// main render loop
func main() {
	g := NewGame()
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Tower Defense")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
