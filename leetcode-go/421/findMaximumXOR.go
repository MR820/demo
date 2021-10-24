/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/11
 * Time 11:28 上午
 */

package main

import (
	"fmt"
)

/**
421. 数组中两个数的最大异或值
给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 231 。

找到 ai 和aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。

你能在O(n)的时间解决这个问题吗？

示例:

输入: [3, 10, 5, 25, 2, 8]

输出: 28

解释: 最大的结果是 5 ^ 25 = 28.
*/

/**
贪心算法+位运算
*/

func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask = mask | (1 << i)
		s := make(map[int]bool)
		for _, num := range nums {
			s[num&mask] = true
		}
		tmp := (1 << i) | res
		for v := range s {
			if _, ok := s[tmp^v]; ok {
				res = tmp
				break
			}
		}
	}
	return res
}

func main() {
	arr := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(arr))
}
