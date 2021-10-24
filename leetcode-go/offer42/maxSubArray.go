/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/11
 * Time 10:47 上午
 */

package main

import "fmt"

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}

func main() {
	s := []int{1, -1, 2, 3, 4, -3, 2, -1, 3, 6, -3}
	fmt.Println(maxSubArray(s))
}
