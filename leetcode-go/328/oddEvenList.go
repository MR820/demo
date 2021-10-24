/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/24
 * Time 1:40 下午
 */

/*
328. 奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
*/

package main

import (
	"fmt"
)

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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	evenHead := head.Next
	even := evenHead
	for odd.Next != nil && odd.Next.Next != nil {
		odd.Next = odd.Next.Next
		odd = odd.Next
		even.Next = even.Next.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
}

func main() {
	root := Constructor(1)
	root.Next = Constructor(2)
	root.Next.Next = Constructor(3)
	root.Next.Next.Next = Constructor(4)
	root.Next.Next.Next.Next = Constructor(5)
	printListNode(oddEvenList(root))
}
