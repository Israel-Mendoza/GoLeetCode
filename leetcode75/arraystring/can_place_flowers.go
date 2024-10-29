package arraystring

func canPlaceFlowers(flowerbed []int, n int) bool {
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
