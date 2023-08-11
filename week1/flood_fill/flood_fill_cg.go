package floodfill

func floodFillCG(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	rows, cols := len(image), len(image[0])
	startingColor := image[sr][sc]

	var recurser func(r, c int)
	recurser = func(r, c int) {
		if r < 0 || r > rows-1 ||
			c < 0 || c > cols-1 ||
			image[r][c] != startingColor {
			return
		}

		image[r][c] = color
		recurser(r-1, c)
		recurser(r, c+1)
		recurser(r+1, c)
		recurser(r, c-1)
	}

	recurser(sr, sc)

	return image
}
