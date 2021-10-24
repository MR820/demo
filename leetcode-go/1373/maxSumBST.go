/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/18
 * Time 4:20 下午
 */

/**
1373. 二叉搜索子树的最大键值和
给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。

二叉搜索树的定义如下：

任意节点的左子树中的键值都 小于 此节点的键值。
任意节点的右子树中的键值都 大于 此节点的键值。
任意节点的左子树和右子树都是二叉搜索树。


示例 1：



输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
输出：20
解释：键值为 3 的子树是和最大的二叉搜索树。
示例 2：



输入：root = [4,3,null,1,2]
输出：2
解释：键值为 2 的单节点子树是和最大的二叉搜索树。
示例 3：

输入：root = [-4,-2,-5]
输出：0
解释：所有节点键值都为负数，和最大的二叉搜索树为空。
示例 4：

输入：root = [2,1,3]
输出：6
示例 5：

输入：root = [5,4,8,3,null,6,3]
输出：7


提示：

每棵树最多有 40000 个节点。
每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。

*/

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if process(root).count < 0 {
		return 0
	} else {
		return process(root).maxCount
	}
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			isBST:    true,
			count:    0,
			maxCount: 0,
			min:      math.MaxInt64,
			max:      math.MinInt64,
		}
	}

	leftInfo := process(root.Left)
	rightInfo := process(root.Right)

	isBST := true
	count := root.Val
	maxCount := root.Val
	min := root.Val
	max := root.Val

	//经过X节点

	if leftInfo.isBST && rightInfo.isBST && root.Val > leftInfo.max && root.Val < rightInfo.min {
		isBST = true
		count = leftInfo.count + rightInfo.count + root.Val
		tmp := 0
		if leftInfo.maxCount > rightInfo.maxCount {
			tmp = leftInfo.maxCount
		} else {
			tmp = rightInfo.maxCount
		}
		if tmp > count {
			maxCount = tmp
		} else {
			maxCount = count
		}
		if leftInfo.min == math.MaxInt64 {
			min = root.Val
		} else {
			min = leftInfo.min
		}
		if rightInfo.max == math.MinInt64 {
			max = root.Val
		} else {
			max = rightInfo.max
		}

	} else {
		//不经过X节点
		isBST = false
		if leftInfo.maxCount > rightInfo.maxCount {
			count = leftInfo.count
			maxCount = leftInfo.maxCount
			min = leftInfo.min
			max = leftInfo.max
		} else {
			count = rightInfo.count
			maxCount = rightInfo.maxCount
			min = rightInfo.min
			max = rightInfo.max
		}
	}

	return &Info{isBST, count, maxCount, min, max}
}

type Info struct {
	isBST    bool
	count    int
	maxCount int
	min      int
	max      int
}

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right.Left = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right.Right = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(maxSumBST(root))
}
