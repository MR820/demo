/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/19
 * Time 5:10 下午
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

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var s []int
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			s = append(s, node.Val)
			root = node.Right
		}
	}
	return s
}

func main() {
	root := Constructor(1)
	root.Right = Constructor(2)
	root.Right.Left = Constructor(3)
	fmt.Println(inorderTraversal(root))
}
