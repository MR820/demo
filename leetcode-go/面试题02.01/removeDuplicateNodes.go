/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/18
 * Time 3:33 下午
 */

/**
面试题 02.01. 移除重复节点
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor(n int) *ListNode {
	return &ListNode{
		Val:  n,
		Next: nil,
	}
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	pre, cur := head, head
	m := make(map[int]int)
	for cur != nil {
		if _, ok := m[cur.Val]; !ok {
			m[cur.Val] = cur.Val
			pre = cur
			cur = cur.Next
		} else {
			//删除此节点
			pre.Next = cur.Next
			cur = cur.Next
		}

	}
	return head
}

func printListNode(head *ListNode) {
	node := head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}

func main() {
	node := Constructor(1)
	node.Next = Constructor(2)
	node.Next.Next = Constructor(3)
	node.Next.Next.Next = Constructor(2)
	node.Next.Next.Next.Next = Constructor(1)
	head := removeDuplicateNodes(node)
	printListNode(head)
}
