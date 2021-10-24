/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/19
 * Time 1:51 下午
 */

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func printTree(head *TreeNode) {
	fmt.Println("Binary:")
	printInOrder(head, 0, "H", 17)
	fmt.Println()
}

func printInOrder(node *TreeNode, height int, to string, length int) {
	if node == nil {
		return
	}

	printInOrder(node.Right, height+1, "V", length)

	val := to + strconv.Itoa(node.Val) + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*length) + val)

	printInOrder(node.Left, height+1, "^", length)
}

func getSpace(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	printTree(buildTree(preorder, inorder))
}
