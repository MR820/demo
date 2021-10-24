package main

import (
	"fmt"
	"os"
)

func readBfs(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	bfs := make([][]int, row)
	for i := range bfs {
		bfs[i] = make([]int, col)
		for j := range bfs[i] {
			fmt.Fscanf(file, "%d", &bfs[i][j])
		}
	}
	return bfs
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(bfs [][]int, start, end point) [][]int {
	steps := make([][]int, len(bfs))
	for i := range bfs {
		steps[i] = make([]int, len(bfs[i]))
	}

	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			// next at bfs is 0
			// and next as steps is 0
			// and next != start
			val, ok := next.at(bfs)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			queue = append(queue, next)
			//fmt.Println(curStep)
		}
	}
	return steps
}

func main() {
	bfs := readBfs("bfs/bfs.in")
	//for _, v := range bfs {
	//	fmt.Println(v)
	//}
	//fmt.Println()
	steps := walk(bfs, point{0, 0}, point{len(bfs) - 1, len(bfs[0]) - 1})
	for _, row := range steps {
		fmt.Println(row)
	}
}
