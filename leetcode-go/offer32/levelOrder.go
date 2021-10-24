/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/1/19
 * Time 4:06 下午
 */

/**
剑指 Offer 32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]


提示：

节点总数 <= 1000
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	s := make([]*Info, 0)
	s = append(s, &Info{
		val:   root,
		level: 0,
	})

	m := make([][]int, 0)
	for len(s) > 0 {
		info := s[0]
		if len(m) < info.level+1 {
			m = append(m, nil)
		}
		m[info.level] = append(m[info.level], info.val.Val)

		s = s[1:]
		if info.val.Left != nil {
			s = append(s, &Info{
				val:   info.val.Left,
				level: info.level + 1,
			})
		}
		if info.val.Right != nil {
			s = append(s, &Info{
				val:   info.val.Right,
				level: info.level + 1,
			})
		}
	}

	return m
}

type Info struct {
	val   *TreeNode
	level int
}

func main() {
	root := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	root.Right.Right = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(levelOrder(root))
}
