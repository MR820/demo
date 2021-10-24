/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 1:53 下午
 */

/**
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
给定的 k 保证是有效的。
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	//倒数第k个，翻转链表即第k个
	node := reverseLink(head)
	val := 0
	for i := 0; i < k; i++ {
		val = node.Val
		node = node.Next
	}
	return val
}

func reverseLink(head *ListNode) *ListNode {
	pre, next := &ListNode{}, &ListNode{}
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	fmt.Println(kthToLast(head, 2))
}
