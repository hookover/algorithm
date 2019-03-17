package main

import "fmt"

var visited [][]bool
var d = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
var m, n int

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func dfs(board [][]byte, x, y int) {
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + d[i][0]
		newY := y + d[i][1]

		if inArea(newX, newY) && !visited[newX][newY] && board[newX][newY] == 1 {
			dfs(board, newX, newY)
		}
	}
}

func numIslands(board [][]byte) int {
	if len(board) == 0 {
		return 0
	}
	m = len(board)
	n = len(board[0])

	visited = make([][]bool, len(board))
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 && !visited[i][j] {
				res ++
				dfs(board, i, j)
			}
		}
	}

	return res
}

func main() {
	graid := [][]byte{
		{0, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0},
		{1, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 1},
		{1, 0, 1, 0, 1, 1},
	}

	fmt.Println(numIslands(graid))
}
