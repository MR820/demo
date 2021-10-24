/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/19
 * Time 2:41 下午
 */

/**

257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return process(root).path
}

func process(node *TreeNode) *Info {
	if node == nil {
		return &Info{path: []string{""}}
	}
	if node.Left == nil && node.Right == nil {
		return &Info{path: []string{strconv.Itoa(node.Val)}}
	}
	leftInfo := process(node.Left)
	rightInfo := process(node.Right)
	path := make([]string, 0)
	for _, v := range leftInfo.path {
		if v == "" {
			continue
		}
		p := strconv.Itoa(node.Val) + "->" + v
		path = append(path, p)
	}
	for _, v := range rightInfo.path {
		if v == "" {
			continue
		}
		p := strconv.Itoa(node.Val) + "->" + v
		path = append(path, p)
	}
	return &Info{path: path}
}

type Info struct {
	path []string
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(binaryTreePaths(root))
}
