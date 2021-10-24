/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/15
 * Time 2:24 下午
 */

package main

import "fmt"

func trap(height []int) int {
	length := len(height)
	left, right, left_max, right_max, ans := 0, length-1, 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else {
				ans += left_max - height[left]
			}
			left += 1
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				ans += right_max - height[right]
			}
			right -= 1
		}
	}
	return ans

}

func main() {
	s := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(s))
}
