package main

import "fmt"

//https://leetcode.com/problems/number-of-islands/

func numIslands(grid [][]byte) int {
	//DFS with as long as we get one for both horizontal and vertical
	count := 0
	for i, row := range grid {
		for j, el := range row {
			if el == '1' {
				count++
				DFS(grid, i, j)
			}
		}
	}
	return count
}
func DFS(grid [][]byte, i, j int) {
	if i < 0 || j < 0 {
		return
	}
	if i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	DFS(grid, i+1, j)
	DFS(grid, i, j+1)
	DFS(grid, i-1, j)
	DFS(grid, i, j-1)
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))

}
