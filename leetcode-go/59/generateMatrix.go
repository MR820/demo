/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/28
 * Time 4:38 下午
 */

/**
59. 螺旋矩阵 II
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
通过次数55,450提交次数70,645
*/
package main

import "fmt"

type point struct {
	i, j int
}

var dirs = [4]point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(s [][]int) bool {
	if p.i < 0 || p.i >= len(s) || p.j < 0 || p.j >= len(s[p.i]) || s[p.i][p.j] > 0 {
		return false
	}
	return true
}

func generateMatrix(n int) [][]int {
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
	}
	cur := 1
	dir := 0
	s[0][0] = cur
	p := point{
		i: 0,
		j: 0,
	}
	for cur < n*n {

		next := p.add(dirs[dir])
		ok := next.at(s)
		if ok {
			p = next
			cur++
			s[p.i][p.j] = cur
		} else {
			dir = (dir + 1) % 4
			continue
		}

	}
	return s
}

func pirntS(s [][]int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%d  ", s[i][j])
		}
		fmt.Println()
	}
}

func main() {
	pirntS(generateMatrix(3))
}
