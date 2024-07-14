package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

// Game struct willl hold all the data the entire game will ned.
type Game struct{}

//create a new game object and initialize the data

func NewGame() *Game {
	g := &Game{}
	return g
}

// callbacks for update, draw and layout
func (g *Game) Update() error {
	return nil
}

// draw similarly is called draw cycle and is where we will blit
func (g *Game) Draw(screen *ebiten.Image) {

}

//layout will return the screen dimensions

func (g *Game) Layout(w, h int) (int, int) {
	return 1280, 800
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
