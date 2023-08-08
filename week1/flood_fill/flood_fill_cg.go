package floodfill

/*
i, j

					    [max(i - 1, 0), j]
					         	N
								|
	   i, max(j - 1, 0)  W -- [i,j] --  E  i, min(j + 1, len())
								|
	 	  					    S
						min(i + 1, len() - 1), j
*/

func FloodFillCG(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	n := Max(sr-1, 0)
	e := Min(j+1, len(image)-1)
	s := Min(sc+1, len(image)-1)
	w := Max(sr-1, 0)

	image[up][sc] = color
	image[sr][e] = color
	image[s][sc] = color
	image[sr][w] = color

	FloodFillCG(image, up, sc, color)
	FloodFillCG(image, sr, e, color)
	FloodFillCG(image, s, sc, color)
	FloodFillCG(image, sr, w, color)

	return image
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
