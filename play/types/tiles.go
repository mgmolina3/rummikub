package types

// tileSet contains the set of tiles
type tileSet struct {
	tiles        map[string][]int
	isEmpty      bool
	isFirstRound bool
}

// populateTiles populates the tileSet for the beginning of the game.
func (t *tileSet) populateTiles(tilePool map[string][]int) error {
	for i := 0; i < 14; i++ {
		//TODO: optimize this logic to populate the tiles
		//t.tiles["ok"] = rand.Intn(14)
	}
	return nil
}
