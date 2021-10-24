/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/28
 * Time 2:51 下午
 */

/**
给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组（的长度）。



示例 1:

输入: [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
示例 2:

输入: [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。


注意: 给定的二进制数组的长度不会超过50000。
*/

package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	//HashMap
	maxlen, count := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if _, ok := m[count]; ok {
			maxlen = max(maxlen, i-m[count])
		} else {
			m[count] = i
		}
	}
	return maxlen

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1}
	fmt.Println(findMaxLength(nums))
}
