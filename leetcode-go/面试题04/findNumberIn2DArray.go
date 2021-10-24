/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/2
 * Time 8:06 下午
 */

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	flag := false
	h := len(matrix) //j
	if h == 0 {
		return false
	}
	l := len(matrix[0]) //i

	for i, j := l-1, 0; i >= 0 && i < l && j >= 0 && j < h; {
		if matrix[j][i] > target {
			i--
		} else if matrix[j][i] < target {
			j++
		} else {
			flag = true
			break
		}
	}

	return flag
}

func main() {
	l := make([][]int, 0)
	for j := 1; j < 8; j++ {
		s := make([]int, 0)
		for i := j + 1; i < j+10; i++ {
			s = append(s, i)
		}
		l = append(l, s)
	}

	fmt.Println(l)
	fmt.Println(findNumberIn2DArray(l, 15))
}
