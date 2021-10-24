/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/18
 * Time 4:50 下午
 */

/**
142. 环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

进阶：

你是否可以使用 O(1) 空间解决此题？


示例 1：



输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：



输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：



输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。


提示：

链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
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

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	node := head
//	m := make(map[*ListNode]int)
//	for node != nil {
//		if _, ok := m[node]; !ok {
//			m[node] = node.Val
//		} else {
//			return node
//		}
//		node = node.Next
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func main() {
	head := Constructor(1)
	head.Next = Constructor(2)
	head.Next.Next = Constructor(0)
	head.Next.Next.Next = Constructor(-4)
	head.Next.Next.Next.Next = head.Next.Next
	fmt.Println(detectCycle(head))
}
