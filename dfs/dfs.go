/**
1是陆地 0是水
陆地通过水平与垂直方向连接
默认网格四周为水
求岛屿的数量
*/

package main

import (
	"fmt"
	"os"
)

//读取网格
func readDfs(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	dfs := make([][]int, row)
	for i := range dfs {
		dfs[i] = make([]int, col)
		for j := range dfs[i] {
			fmt.Fscanf(file, "%d", &dfs[i][j])
		}
	}
	return dfs
}

//获取岛屿数量
func numIslands(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	num := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

//深度遍历
func dfs(grid [][]int, i, j int) {
	row := len(grid)
	col := len(grid[0])
	if i < 0 || i >= row || j < 0 || j >= col || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i, j-1)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i+1, j)

}

func main() {
	dfs := readDfs("dfs/dfs.in")
	//for i := range dfs {
	//	for j := range dfs[i] {
	//		fmt.Print(dfs[i][j])
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}
	num := numIslands(dfs)
	fmt.Println(num)
}
