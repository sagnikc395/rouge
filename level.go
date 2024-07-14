package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// a struct to hold all of our static global game data
type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// load game data into the constructor
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}

	return g
}

//structure to hold our individual tiles or squares on our squares on our map

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

//Optimization -> rather than stoer the tile in our array of arrays to indicate the 2d
//array of our map, we will store all tiles in 1 slice and utilize a helper function to determine
// which tile to return

//GetIndexFromXY -> gets the index of the map array from a given X,Y tile coordinates
// this coordinate is logical tiles , not pixels

func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// map title is our initial map. This will be a simple one which consist of all floor tiles except the outer
// bounds of the window , which will be walls.

func CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {

				wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
				if err != nil {
					log.Fatal(err)
				}

				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: true,
					Image:   wall,
				}

				tiles = append(tiles, tile)
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: false,
					Image:   floor,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}
