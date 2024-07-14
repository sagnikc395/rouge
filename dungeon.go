//dungeon will consist of a name -> what the dungeon is
// and a slice of levels -> the generated levels for the dungeon

package main

// dungeon is a container for all the levels that make a particular dungeon

type Dungeon struct {
	Name   string
	Levels []Level
}
