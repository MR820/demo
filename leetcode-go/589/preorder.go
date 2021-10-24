/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/20
 * Time 11:54 上午
 */

package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

/*
//迭代
func preorder(root *Node) []int {
	s := make([]int, 0)
	if root == nil {
		return s
	}
	arr := make([]*Node, 0)
	arr = append(arr, root)
	for len(arr) != 0 {
		node := arr[len(arr)-1]
		s = append(s, node.Val)
		arr = arr[:len(arr)-1]
		for i := len(node.Children) - 1; i >= 0; i-- {
			arr = append(arr, node.Children[i])
		}
	}
	return s
}

*/

var s []int

//递归
func preorder(root *Node) []int {
	s = []int{}
	recursion(root)
	return s

}

func recursion(root *Node) {
	if root != nil {
		s = append(s, root.Val)
		for _, v := range root.Children {
			recursion(v)
		}
	}
}

func main() {

	root := &Node{
		Val:      1,
		Children: nil,
	}
	node1 := &Node{
		Val:      3,
		Children: nil,
	}
	node2 := &Node{
		Val:      2,
		Children: nil,
	}
	node3 := &Node{
		Val:      4,
		Children: nil,
	}
	root.Children = []*Node{node1, node2, node3}
	node5 := &Node{
		Val:      5,
		Children: nil,
	}
	node6 := &Node{
		Val:      6,
		Children: nil,
	}
	node1.Children = []*Node{node5, node6}
	fmt.Println(preorder(root))

}
