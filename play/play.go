package play

// No error handling logic as of yet, returning nil for now

// Play ..
func Play() error {
	if _, err := GenerateTilePool(); err != nil {
		return err
	}
	return nil
}

// GenerateTilePool ..
func GenerateTilePool() (map[string][]int, error) {
	// General logic for this. Might want to optimize
	tileNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	tilePool := map[string][]int{
		"black":  tileNumbers,
		"blue":   tileNumbers,
		"orange": tileNumbers,
		"red":    tileNumbers,
	}
	return tilePool, nil
}
