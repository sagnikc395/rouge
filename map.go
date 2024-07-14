package main

//gamemap will hold all the levels and aggregate information for the entire world

type GameMap struct {
	Dungeons []Dungeon
}

//gamemap constructor to create the gamemap

func NewGameMap() GameMap {
	//return a new game map of a single level for now
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	d := Dungeon{Name: "default", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	gm := GameMap{Dungeons: dungeons}
	return gm
}
