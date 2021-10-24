/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/21
 * Time 4:29 下午
 */

/**
给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:

// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	len  int
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	solution := Solution{
		len:  0,
		head: head,
	}
	for head != nil {
		solution.len++
		head = head.Next
	}
	return solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(this.len)
	node := this.head
	for i > 0 {
		node = node.Next
		i--
	}
	return node.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	solution := Constructor(head)
	fmt.Println(solution.GetRandom())
}
