/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/9/15
 * Time 下午5:27
 */

/**
连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。



示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:

输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。


提示：

1 <= nums.length <= 105
nums[i] 不是 0 就是 1
*/
package main

import "fmt"

func main() {
	nums := []int{1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0}
	i := findMaxLength(nums)
	fmt.Println(i)
}

func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	counter := 0
	maxLength := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if preVIndex, ok := mp[counter]; ok {
			maxLength = max(maxLength, i-preVIndex)
		} else {
			mp[counter] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
