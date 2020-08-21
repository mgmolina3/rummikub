package play

// Play ..
func Play() error {
	if _, err := GenerateTilePool(); err != nil {
		return err
	}
	return nil
}

// GenerateTilePool ..
func GenerateTilePool() (map[string][]int, error) {
	tileNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	tilePool := map[string][]int{
		"black":  tileNumbers,
		"blue":   tileNumbers,
		"orange": tileNumbers,
		"red":    tileNumbers,
	}
	return tilePool, nil
}
