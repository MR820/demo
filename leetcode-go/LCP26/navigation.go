/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/24
 * Time 11:37 上午
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(n int) *TreeNode {
	return &TreeNode{
		Val:   n,
		Left:  nil,
		Right: nil,
	}
}

type Info struct {
	num int
}

func process(root *TreeNode) Info {
	if root == nil {
		return Info{num: 0}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)

	return Info{leftInfo.num + rightInfo.num + 1}
}

func navigation(root *TreeNode) int {
	info := process(root)
	return info.num
}

func main() {
	root := Constructor(1)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(3)
	root.Left.Right = Constructor(4)
	fmt.Println(navigation(root))
}
