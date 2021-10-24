/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/27
 * Time 1:32 下午
 */

package main

import (
	"fmt"
)

/**
454. 四数相加 II
给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

/*
//超时
func fourSumCount(A []int, B []int, C []int, D []int) int {
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	sort.Ints(D)
	count := 0
	lA := len(A)
	lB := len(B)
	lC := len(C)
	lD := len(D)
	for i := 0; i < lA; i++ {
		for j := 0; j < lB; j++ {
			for k := 0; k < lC; k++ {
				for m := 0; m < lD; m++ {
					if A[i]+B[j]+C[k]+D[m] > 0 {
						break
					} else if A[i]+B[j]+C[k]+D[m] == 0 {
						count++
					}
				}
			}
		}
	}
	return count

}

*/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	//分区 分治

	m := map[int]int{}
	count := 0
	for _, v := range A {
		for _, w := range B {
			m[v+w]++
		}
	}
	for _, v := range C {
		for _, w := range D {
			if _, ok := m[-v-w]; ok {
				count += m[-v-w]
			}
		}
	}
	return count
}

func main() {
	A := []int{-1, 1, 1, 1, -1}
	B := []int{0, -1, -1, 0, 1}
	C := []int{-1, -1, 1, -1, -1}
	D := []int{0, 1, 0, -1, -1}
	fmt.Println(fourSumCount(A, B, C, D))
}
