/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/21
 * Time 5:10 下午
 */

package main

import "fmt"

func numIsLands(grid [][]string) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	num := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "1" {
				num++
				dfs(grid, i, j)
			}

		}
	}
	return num
}

func dfs(grid [][]string, i, j int) {
	w, h := len(grid), len(grid[0])
	if i < 0 || i >= w || j < 0 || j >= h || grid[i][j] == "0" {
		return
	}
	grid[i][j] = "0"
	dfs(grid, i, j-1)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i+1, j)
}

func main() {
	grid := [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "1"},
	}
	fmt.Println(numIsLands(grid))

}
