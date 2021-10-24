/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/24
 * Time 9:53 上午
 */

package main

import "fmt"

/**
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
*/

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

/*
type Info struct {
	count int
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			count: 0,
		}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)
	count := leftInfo.count + rightInfo.count + 1
	return &Info{
		count: count,
	}
}

func countNodes(root *TreeNode) int {
	info := process(root)
	return info.count
}


*/

func countNodes(root *TreeNode) int {
	//利用完全二叉树特性
	if root == nil {
		return 0
	}
	level := 0
	node := root
	for node.Left != nil {
		level++
		node = node.Left
	}
	min := 1 << level
	max := 1<<(level+1) - 1
	for min < max {
		mid := min + (max-min+1)/2
		if exists(root, level, mid) {
			min = mid
		} else {
			max = mid - 1
		}
	}
	return min
}

func exists(root *TreeNode, level int, k int) bool {
	bits := 1 << (level - 1)
	node := root
	for node != nil && bits > 0 {
		if bits&k == 0 {
			node = node.Left
		} else {
			node = node.Right
		}
		bits >>= 1
	}
	return node != nil

}

func main() {
	root := Constructor(1)
	root.Left = Constructor(2)
	root.Left.Left = Constructor(4)
	root.Left.Right = Constructor(5)
	root.Right = Constructor(3)
	root.Right.Left = Constructor(6)
	fmt.Println(countNodes(root))
}
