/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/19
 * Time 5:39 下午
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

func levelOrder(root *TreeNode) [][]int {
	var s [][]int
	if root == nil {
		return s
	}
	var cur []*TreeNode
	var next []*TreeNode
	cur = append(cur, root)
	for cur != nil {
		var arr []int
		for _, v := range cur {
			arr = append(arr, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		s = append(s, arr)
		cur = next
		next = nil
	}
	return s

}

func main() {
	root := Constructor(3)
	root.Left = Constructor(9)
	root.Right = Constructor(20)
	root.Right.Left = Constructor(15)
	root.Right.Right = Constructor(7)
	fmt.Println(levelOrder(root))
}
