/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 4:50 下午
 */

/**
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。



示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//先翻转两个链表，从第一个节点依次相加，有进位问题
	l1 = reverseLink(l1)
	l2 = reverseLink(l2)

	h1, h2 := l1, l2
	m, n := 1, 1
	//计算链表长度
	for h1.Next != nil {
		m++
		h1 = h1.Next
	}
	for h2.Next != nil {
		n++
		h2 = h2.Next
	}
	//补全链表
	if m > n {
		for i := 0; i < m-n; i++ {
			h2.Next = &ListNode{Val: 0}
		}
	}
	if m < n {
		for i := 0; i < n-m; i++ {
			h1.Next = &ListNode{Val: 0}
		}
	}
	printLink(l1)
	printLink(l2)

	cur := l1
	carry := 0

	for l1 != nil {
		l1.Val = l1.Val + l2.Val + carry
		carry = l1.Val / 10
		l1.Val = l1.Val % 10
		l1 = l1.Next
		l2 = l2.Next
	}
	pre := cur
	for cur.Next != nil {
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return reverseLink(pre)

}

func main() {
	head := &ListNode{Val: 9}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 6}

	head1 := &ListNode{Val: 0}

	printLink(addTwoNumbers(head, head1))
}

func reverseLink(head *ListNode) *ListNode {
	pre, next := (*ListNode)(nil), (*ListNode)(nil)
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func printLink(head *ListNode) {
	fmt.Println()
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}
