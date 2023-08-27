package floodfill

var (
	queue  [][]int
	xb, yb int
)

func floodFillCR(img [][]int, sr, sc, color int) [][]int {

	if img[sr][sc] == color {
		return img
	}

	queue = append(queue, []int{sr, sc, img[sr][sc]})
	xb = (len(img[0]) - 1)
	yb = (len(img) - 1)

	for len(queue) > 0 {
		p := queue[0]
		oc := p[2]
		x, y := p[0], p[1]

		if x > 0 {
			scan(x-1, y, oc, img[x-1][y])
		}

		if x < xb {
			scan(x+1, y, oc, img[x+1][y])
		}

		if y > 0 {
			scan(x, y-1, oc, img[x][y-1])
		}

		if y < yb {
			scan(x, y+1, oc, img[x][y+1])
		}

		fill(x, y, color, img)
		queue = queue[1:]
	}

	return img
}

func fill(x, y, color int, img [][]int) {
	img[x][y] = color
}

func scan(x, y, oc, pc int) {
	if pc == oc {
		queue = append(queue, []int{x, y, pc})
	}
}
