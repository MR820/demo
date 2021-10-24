/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/1
 * Time 4:59 下午
 */

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/

package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	nums := []int{1, 3, 5, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
