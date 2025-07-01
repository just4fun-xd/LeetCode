package easy

func canPlaceFlowers(flowerbed []int, n int) bool {
	i := 0
	for i < len(flowerbed) && n > 0 {
		if (flowerbed[i] == 0) && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		i++
	}
	if n != 0 {
		return false
	}
	return true
}
