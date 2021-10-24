/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/18
 * Time 4:13 下午
 */

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	jw := 0
	sum := 0
	val := 0
	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := root
	for l1 != nil || l2 != nil || jw != 0 {
		sum = 0
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = v1 + v2 + jw
		jw = sum / 10
		val = sum % 10
		root.Next = Constructor(val)
		root = root.Next
	}
	return res.Next
}

func printListNode(root *ListNode) {
	node := root
	for node != nil {
		fmt.Print(node.Val)
		fmt.Print("->")
		node = node.Next
	}
}

func main() {
	l1 := Constructor(2)
	l1.Next = Constructor(4)
	l1.Next.Next = Constructor(3)
	l2 := Constructor(5)
	l2.Next = Constructor(6)
	l2.Next.Next = Constructor(4)
	res := addTwoNumbers(l1, l2)
	printListNode(res)
}
