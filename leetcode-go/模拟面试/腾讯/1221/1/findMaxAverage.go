/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/21
 * Time 3:19 下午
 */

/**
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

示例 1:

输入: [1,12,-5,-6,50,3], k = 4
输出: 12.75
解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


注意:

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。
*/

package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	l := len(nums)
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	sum := 0
	for i := 1; i <= l-k; i++ {
		sum = 0
		for j := i; j < k+i; j++ {
			sum += nums[j]
		}
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Println(findMaxAverage(nums, k))
}
