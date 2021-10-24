/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/22
 * Time 5:38 下午
 */

/**
103. 二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
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

func zigzagLevelOrder(root *TreeNode) [][]int {

}

func main() {
	root := Constructor(3)
	root.Left = Constructor(9)
	root.Right = Constructor(20)
	root.Right.Left = Constructor(15)
	root.Right.Right = Constructor(7)
	fmt.Println(zigzagLevelOrder(root))
}
