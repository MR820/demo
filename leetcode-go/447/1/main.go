/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/13
 * Time 下午1:35
 */

/**
给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，
其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。

返回平面上所有回旋镖的数量。




示例 1：

输入：points = [[0,0],[1,0],[2,0]]
输出：2
解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
示例 2：

输入：points = [[1,1],[2,2],[3,3]]
输出：2
示例 3：

输入：points = [[1,1]]
输出：0


提示：

n == points.length
1 <= n <= 500
points[i].length == 2
-104 <= xi, yi <= 104
所有点都 互不相同


一定是双数，只需对比一边*2即可

循环遍历
时间复杂度 n^2，组合

*/

package main

import (
	"fmt"
)

func main() {
	p := make([][]int, 0)
	p = append(p, []int{0, 0})
	p = append(p, []int{1, 0})
	p = append(p, []int{-1, 0})
	p = append(p, []int{0, 1})
	p = append(p, []int{0, -1})
	fmt.Println(numberOfBoomerangs(p))
}

func numberOfBoomerangs(points [][]int) int {

	l := len(points)

	s1 := make([][]int, l)
	for x := 0; x < l; x++ {
		s1[x] = make([]int, l)
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			s1[i][j] = (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) +
				(points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			s1[j][i] = s1[i][j]
		}
	}

	s := make([]int, 0)

	for _, val := range s1 {
		m := make(map[int]int)
		for _, v := range val {
			if _, ok := m[v]; !ok {
				m[v] = 1
			} else {
				m[v] += 1
			}
		}
		for k, v := range m {
			if k != 0 && v > 1 {
				s = append(s, v)
			}
		}
	}

	num := 0
	for _, v := range s {
		num += v * (v - 1)

	}
	return num
}

func C(n, k int) int {
	return F(n) / F(n-k) / F(k)
}

func F(n int) int {
	if n <= 1 {
		return 1
	}
	return n * F(n-1)
}
