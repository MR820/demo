/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/13
 * Time 下午3:37
 */

/**
给你两个大小为 n x n 的二进制矩阵 mat 和 target 。现 以 90 度顺时针轮转 矩阵 mat 中的元素 若干次 ，
如果能够使 mat 与 target 一致，返回 true ；否则，返回 false 。
示例 1：


输入：mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
输出：true
解释：顺时针轮转 90 度一次可以使 mat 和 target 一致。
示例 2：


输入：mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
输出：false
解释：无法通过轮转矩阵中的元素使 equal 与 target 一致。
示例 3：


输入：mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
输出：true
解释：顺时针轮转 90 度两次可以使 mat 和 target 一致。


提示：

n == mat.length == target.length
n == mat[i].length == target[i].length
1 <= n <= 10
mat[i][j] 和 target[i][j] 不是 0 就是 1


*/

package main

import "fmt"

func main() {
	var mat = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	var target = [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(findRotation(mat, target))
}

func findRotation(mat [][]int, target [][]int) bool {
	// 90 行 为 列 倒排
	// 旋转 3 次后，仍然false 则false

	if equal(mat, target) {
		return true
	}

	for x := 0; x < 3; x++ {

		a := transform(mat)
		mat = a
		if equal(mat, target) {
			return true
		}
	}
	return false
}

func transform(mat [][]int) [][]int {
	l := len(mat)
	a := make([][]int, l)
	for k, _ := range a {
		a[k] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			a[i][j] = mat[l-1-j][i]
		}
	}
	return a
}

func equal(a, b [][]int) bool {
	for i, val := range a {
		for j, v := range val {
			if b[i][j] != v {
				return false
			}
		}
	}
	return true
}
