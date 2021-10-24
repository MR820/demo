/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/15
 * Time 1:55 下午
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

func main() {
	s := []int{-10, -5, 0, 1, 3}
	fmt.Println(sortedArrayToBST(s))
}
