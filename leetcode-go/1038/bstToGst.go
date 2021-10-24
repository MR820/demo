/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/18
 * Time 3:36 下午
 */

/**
1038. 把二叉搜索树转换为累加树
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：该题目与 538: https://leetcode-cn.com/problems/convert-bst-to-greater-tree/  相同



示例 1：



输入：[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：

输入：root = [0,null,1]
输出：[1,null,1]
示例 3：

输入：root = [1,0,2]
输出：[3,3,2]
示例 4：

输入：root = [3,2,4,1]
输出：[7,9,4,10]


提示：

树中的节点数介于 1 和 100 之间。
每个节点的值介于 0 和 100 之间。
树中的所有值 互不相同 。
给定的树为二叉搜索树。
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s []int = make([]int, 0)

func pre(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		pre(root.Left)
	}
	s = append(s, root.Val)
	if root.Right != nil {
		pre(root.Right)
	}
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre(root)
	length := len(s)
	m := make(map[int]int, 0)
	m[s[length-1]] = s[length-1]
	for i := length - 2; i >= 0; i-- {
		m[s[i]] = s[i] + s[i+1]
		s[i] = s[i] + s[i+1]
	}
	fmt.Println(s)
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Val = m[node.Val]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}

func printTree(head *TreeNode) {
	if head == nil {
		return
	}
	var s []*TreeNode
	s = append(s, head)
	for len(s) > 0 {
		node := s[0]
		fmt.Println(node.Val)
		s = s[1:]
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}

	}

}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	//root.Left.Right = &TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Left.Right.Right = &TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right = &TreeNode{
	//	Val:   6,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right.Left = &TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right.Right = &TreeNode{
	//	Val:   7,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right.Right.Right = &TreeNode{
	//	Val:   8,
	//	Left:  nil,
	//	Right: nil,
	//}
	node := bstToGst(root)
	printTree(node)
}
