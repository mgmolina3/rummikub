package types

// Player ..
type Player struct {
	tiles tileSet
}

func (p *Player) hasNoTiles() bool {
	return p.tiles.isEmpty
}

func (p *Player) sumOfTiles() int {
	return 0
}
