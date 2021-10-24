/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/19
 * Time 5:23 下午
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

func preorderTraversal(root *TreeNode) []int {

	var s []int
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
		for len(stack) != 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			s = append(s, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return s

}

func main() {
	root := Constructor(1)
	root.Right = Constructor(2)
	root.Right.Left = Constructor(3)
	fmt.Println(preorderTraversal(root))
}
