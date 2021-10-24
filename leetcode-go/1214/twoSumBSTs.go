/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/18
 * Time 10:46 上午
 */

/**
1214. 查找两棵二叉搜索树之和
给出两棵二叉搜索树，请你从两棵树中各找出一个节点，使得这两个节点的值之和等于目标值 Target。

如果可以找到返回 True，否则返回 False。



示例 1：



输入：root1 = [2,1,4], root2 = [1,0,3], target = 5
输出：true
解释：2 加 3 和为 5 。
示例 2：



输入：root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
输出：false


提示：

每棵树上最多有 5000 个节点。
-10^9 <= target, node.val <= 10^9
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pre(head *TreeNode, s *[]int) {
	if head == nil {
		return
	}
	if head.Left != nil {
		pre(head.Left, s)
	}
	*s = append(*s, head.Val)
	if head.Right != nil {
		pre(head.Right, s)
	}
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	s1 := make([]int, 0)
	pre(root1, &s1)
	s2 := make([]int, 0)
	pre(root2, &s2)
	m := make(map[int]int, 0)
	for _, v := range s2 {
		m[v] = 1
	}
	for _, v := range s1 {
		n := target - v
		if _, ok := m[n]; ok {
			return true
		}
	}
	return false
}

func main() {
	root1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root1.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root1.Right = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root2.Left = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	root2.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	fmt.Print(twoSumBSTs(root1, root2, 5))
}
