package arraystring

// CanPlaceFlowers determines if n new flowers can be planted in a flowerbed without violating the no-adjacent-flowers rule.
//
// Args:
//
//	flowerbed: A slice of integers representing the flowerbed. 1 indicates a planted plot, 0 indicates an empty plot.
//	n: The number of flowers to plant.
//
// Returns:
//
//	true if n flowers can be planted, false otherwise.
func CanPlaceFlowers(flowerbed []int, n int) bool {
	prevPlotIsPlanted := false
	for _, plot := range flowerbed {
		if plot == 1 {
			if prevPlotIsPlanted {
				n++
			}
			prevPlotIsPlanted = true
		} else if !prevPlotIsPlanted {
			n--
			prevPlotIsPlanted = true
		} else {
			prevPlotIsPlanted = false
		}
		if n < 0 {
			return true
		}
	}
	return n <= 0
}
