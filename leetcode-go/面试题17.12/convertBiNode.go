/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/13
 * Time 9:31 上午
 */

/**
面试题 17.12. BiNode
二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。

返回转换后的单向链表的头节点。

注意：本题相对原题稍作改动



示例：

输入： [4,2,5,1,3,null,6,0]
输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]
提示：

节点数量不会超过 100000。
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var s []*TreeNode

func in(root *TreeNode) {
	if root == nil {
		return
	}
	node := root
	in(node.Left)
	s = append(s, node)
	in(node.Right)
}

func buildByQueue() *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return build()
}

func build() *TreeNode {
	if len(s) == 0 {
		return nil
	}
	node := s[0]
	s = s[1:]
	node.Left = nil
	node.Right = build()
	return node
}

func convertBiNode(root *TreeNode) *TreeNode {
	//首先先中序遍历得到从小到大的有序队列
	//队列中第一个为头节点，所有节点做节点为空，右节点为下一个节点
	in(root)

	return buildByQueue()
}

func main() {
	node := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node.Left.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node.Left.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node.Right.Right = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	node.Left.Left.Left = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	convertBiNode(node)
}
