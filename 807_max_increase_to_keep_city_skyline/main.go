package main

func main() {
	grid := [][]int{
		[]int{3, 0, 8, 4},
		[]int{2, 4, 5, 7},
		[]int{9, 2, 6, 3},
		[]int{0, 3, 1, 0},
	}

	maxIncreaseKeepingSkyline(grid)
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	leftVal := make([]int, len(grid[0]), len(grid[0]))
	topVal := make([]int, len(grid[0]), len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if leftVal[i] < grid[i][j] {
				leftVal[i] = grid[i][j]
			}

			if topVal[i] < grid[j][i] {
				topVal[i] = grid[j][i]
			}
		}
	}

	sum := 0
	var val int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if leftVal[i] > topVal[j] {
				val = topVal[j]
			} else {
				val = leftVal[i]
			}
			if v := val - grid[i][j]; v > 0 {
				sum += v
			}
		}
	}

	return sum
}
