/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 2:15 下午
 */

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。


输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	//先确定链表长度，确保在一个循环内结束
	val := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		val++
	}
	tail.Next = head
	k = k % val
	//右移k位，即tail走n-k次。tail.Next为新的头节点
	for i := 0; i < val-k; i++ {
		tail = tail.Next
	}
	head, tail.Next = tail.Next, nil
	return head
}

func main() {
	head := &ListNode{Val: 1}

	//printLink(head)

	printLink(rotateRight(head, 2))
}

func printLink(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
