/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/10
 * Time 9:58 上午
 */

package main

import "math/rand"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
深度优先遍历
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	root := &TreeNode{Val: 1}
	p := root
	for i := 0; i < 10; i++ {
		createTreeNode(p.Left)
		createTreeNode(p.Right)

	}

	hasPathSum(root, 12)

}

func createTreeNode(root *TreeNode) {
	root.Left = &TreeNode{Val: rand.Int()}
	root.Right = &TreeNode{Val: rand.Int()}
}

/**
广度优先遍历
*/
/**
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == sum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}
*/
