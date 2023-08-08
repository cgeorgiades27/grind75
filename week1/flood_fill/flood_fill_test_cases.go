package floodfill

var TestCases = []struct {
	image  [][]int
	sr     int
	sc     int
	color  int
	output [][]int
}{
	{
		image: [][]int{{1,1,1},{1,1,0},{1,0,1}},
		sr: 1,
		sc: 1,
		color:  2,
		output: [][]int{{2,2,2},{2,2,0},{2,0,1}}

	},
	{
		image: [][]int{{0,0,0},{0,0,0}},
		sr: 0,
		sc: 0,
		color:  0,
		output: [][]int{{0,0,0},{0,0,0}}
	},
}