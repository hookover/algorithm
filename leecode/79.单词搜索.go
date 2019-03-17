package main

import "fmt"

var visited [][]bool
var d = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
var m, n int

func searchWord(board [][]int, word string, index, startX, startY int) bool {
	if index == len(word)-1 {
		return board[startX][startY] == int(word[index])
	}

	if board[startX][startY] == int(word[index]) {
		visited[startX][startY] = true
		for i := 0; i < 4; i++ {
			newX := startX + d[i][0]
			newY := startY + d[i][1]

			if inArea(newX, newY) && !visited[newX][newY] {
				if searchWord(board, word, index+1, newX, newY) {
					return true
				}
			}
		}
		visited[startX][startY] = false
	}
	return false
}

func inArea(x, y int) bool {
	return x >= 0 && y >= 0 && x < m && y < n
}

func exists(board [][]int, search string) bool {
	if len(board) == 0 {
		panic("board is empty")
	}
	m, n = len(board), len(board[0])

	visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, search, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]int{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exists(board,"ABCCEDA"))
	fmt.Println(exists(board,"ABCESE"))
	fmt.Println(exists(board,"DFCSE"))
	fmt.Println(exists(board,"AFOPD"))
}
