package easy

func floodFill(image [][]int, sr, sc, newColor int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == newColor {
		return image
	}
	queue := [][]int{{sr, sc}}
	for len(queue) > 0 {
		node := queue[0]
		x, y := node[0], node[1]

		queue = queue[1:]

		if image[x][y] == originalColor {
			image[x][y] = newColor
			neighbors := findNeighbors(x, y, image)
			for _, neighbor := range neighbors {
				queue = append(queue, neighbor)
			}

		}
	}
	return image
}

func findNeighbors(x, y int, matrix [][]int) [][]int {
	result := [][]int{}
	if x+1 < len(matrix) {
		result = append(result, []int{x + 1, y})
	}
	if x-1 >= 0 {
		result = append(result, []int{x - 1, y})
	}
	if y+1 < len(matrix[x]) {
		result = append(result, []int{x, y + 1})
	}
	if y-1 >= 0 {
		result = append(result, []int{x, y - 1})
	}
	return result
}
