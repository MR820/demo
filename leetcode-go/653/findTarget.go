/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/31
 * Time 4:14 下午
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	s := make([]*TreeNode, 0)
	s = append(s, root)
	m := make(map[int]int)
	for len(s) != 0 {
		node := s[0]
		s = s[1:]
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
		if _, ok := m[k-node.Val]; ok {
			return true
		}
		m[node.Val] = 1
	}
	return false
}

func main() {
	root := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(findTarget(root, 9))
}
