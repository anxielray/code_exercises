package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	grid := make([][]byte, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}

	if isSea(grid, N) {
		fmt.Println("Sea")
	} else {
		fmt.Println("Lake")
	}
}

func isSea(grid [][]byte, N int) bool {
	visited := make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, N)
	}

	// BFS/DFS to check if water is connected to the border
	queue := [][]int{}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 'o' && (i == 0 || j == 0 || i == N-1 || j == N-1) {
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			}
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		x, y := cell[0], cell[1]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && ny >= 0 && nx < N && ny < N && grid[nx][ny] == 'o' && !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, []int{nx, ny})
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 'o' && !visited[i][j] {
				return false
			}
		}
	}

	return true
}
