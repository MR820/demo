/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/14
 * Time 4:15 下午
 */

package main

import "fmt"

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] = nums[i-1] + nums[i]
		}
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}
