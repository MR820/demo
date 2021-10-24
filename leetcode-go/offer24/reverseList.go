/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/9/10
 * Time 11:45 上午
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	s := []*ListNode{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	for i := len(s) - 1; i > 0; i-- {
		s[i].Next = s[i-1]
	}
	s[0].Next = nil
	return s[len(s)-1]
}
