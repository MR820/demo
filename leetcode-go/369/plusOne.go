/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 4:20 下午
 */

/**
用一个 非空 单链表来表示一个非负整数，然后将这个整数加一。

你可以假设这个整数除了 0 本身，没有任何前导的 0。

这个整数的各个数位按照 高位在链表头部、低位在链表尾部 的顺序排列。

示例:

输入: [1,2,3]
输出: [1,2,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func plusOne(head *ListNode) *ListNode {
	//+1 有可能产生进位，要能找到上一位
	//翻转链表，第一位+1，有进位往后。最后一位进位有则+1，无则创建
	//再次翻转
	head = reverseLink(head)
	printLink(head)
	cur := head
	for cur != nil {
		cur.Val = (cur.Val + 1) % 10
		if cur.Val == 0 {
			//进位,继续循环
			//如果进位到最后Next为nil
			if cur.Next == nil {
				cur.Next = &ListNode{Val: 1}
				break
			}
		} else {
			//无进位，终止循环
			break
		}
		cur = cur.Next
	}
	return reverseLink(head)
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

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 8}
	printLink(head)
	printLink(plusOne(head))
}

func printLink(head *ListNode) {
	fmt.Println()
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}
