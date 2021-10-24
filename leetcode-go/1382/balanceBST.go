/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/19
 * Time 4:14 下午
 */

/**
1382. 将二叉搜索树变平衡
给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。

如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。

如果有多种构造方法，请你返回任意一种。



示例：



输入：root = [1,null,2,null,3,null,4,null,null]
输出：[2,1,3,null,null,null,4]
解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。


提示：

树节点的数目在 1 到 10^4 之间。
树节点的值互不相同，且在 1 到 10^5 之间。
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s []*TreeNode

func Constructor(n int) *TreeNode {
	return &TreeNode{
		Val:   n,
		Left:  nil,
		Right: nil,
	}
}

//先序遍历
func pre(root *TreeNode) {
	if root == nil {
		return
	}
	pre(root.Left)
	s = append(s, root)
	pre(root.Right)
}

func balanceBST(root *TreeNode) *TreeNode {
	pre(root)

	length := len(s)
	if length == 0 {
		return nil
	}
	mid := length / 2
	head := s[mid]
	s[mid].Left = nil
	s[mid].Right = nil

	left, right := mid-1, mid+1
	l, r := head, head
	for left >= 0 {
		l.Left = s[left]
		s[left].Left = nil
		s[left].Right = nil
		l = l.Left
		left--
	}
	for right < length {
		r.Right = s[right]
		s[right].Left = nil
		s[right].Right = nil
		r = r.Right
		right++
	}
	return head
}

func printTree(head *TreeNode) {
	if head == nil {
		return
	}
	if head != nil {
		fmt.Print(head.Val)
	}
	if head.Left != nil {
		printTree(head.Left)
	}

	if head.Right != nil {
		printTree(head.Right)
	}
}

func main() {
	root := Constructor(1)
	root.Right = Constructor(2)
	root.Right.Right = Constructor(3)
	root.Right.Right.Right = Constructor(4)

	head := balanceBST(root)

	printTree(head)

	fmt.Println()

	for _, v := range s {
		fmt.Println(v.Val)
	}

}
