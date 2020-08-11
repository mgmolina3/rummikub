package rummikub

// basic types needed for the game

// table contains the each player's board of the game
type table struct {
	boards []board
}

// board contains the set of tiles each player has
type board struct {
	tiles []tile // your set of tiles at any given time in the board

}

// tile is the individual tile, contains the color and number
// where number ranges from 0-14 (0 can only be used twice as it
// represents the free tile) and the colors are limited to:
// red, black, blue, and orange.
type tile struct {
	color string
	number int
}
