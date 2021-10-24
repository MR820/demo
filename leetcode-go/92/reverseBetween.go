/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/21
 * Time 10:50 上午
 */

/**
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode = nil
	cur := head
	for m > 1 {
		pre = cur
		cur = cur.Next
		m--
		n--
	}
	con, tail := pre, cur
	var next *ListNode = nil
	for n > 0 {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		n--
	}
	if con != nil {
		con.Next = pre
	} else {
		head = pre
	}
	tail.Next = cur
	return head
}

func reverseListNode(root *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := root
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func printListNode(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val)
		root = root.Next
		fmt.Print("->")
	}
}

func main() {
	root := &ListNode{
		Val:  7,
		Next: nil,
	}
	root.Next = &ListNode{
		Val:  9,
		Next: nil,
	}
	root.Next.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	root.Next.Next.Next = &ListNode{
		Val:  10,
		Next: nil,
	}
	root.Next.Next.Next.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	root.Next.Next.Next.Next.Next = &ListNode{
		Val:  8,
		Next: nil,
	}
	root.Next.Next.Next.Next.Next.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	printListNode(reverseBetween(root, 2, 4))
}
