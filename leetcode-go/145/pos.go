/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/20
 * Time 1:32 下午
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

func postorderTraversal(root *TreeNode) []int {
	s := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return s
	}
	h := root
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node.Left != nil && node.Left != h && node.Right != h {
			stack = append(stack, node.Left)
		} else if node.Right != nil && node.Right != h {
			stack = append(stack, node.Right)
		} else {
			s = append(s, node.Val)
			stack = stack[:len(stack)-1]
			h = node
		}
	}

	return s
}

/**
var s []int

func postorderTraversal(root *TreeNode) []int {
	s = []int{}
	recursion(root)
	return s
}

//递归
func recursion(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		recursion(root.Left)
	}
	if root.Right != nil {
		recursion(root.Right)
	}
	s = append(s, root.Val)
}

*/

func main() {
	root := Constructor(1)
	root.Right = Constructor(2)
	root.Right.Left = Constructor(3)
	fmt.Println(postorderTraversal(root))
}
