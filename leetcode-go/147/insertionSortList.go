/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/20
 * Time 9:42 上午
 */

/**
147. 对链表进行插入排序
对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。



插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。


示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//借助数组
	s := make([]*ListNode, 0)
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	for i := 0; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j].Val < s[j-1].Val {
				swap(s, j, j-1)
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			s[i].Next = nil
		} else {
			s[i].Next = s[i+1]
		}
	}
	return s[0]
}

func swap(arr []*ListNode, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func printListNode(root *ListNode) {
	for root != nil {
		fmt.Println(root.Val)
		root = root.Next
	}
}

func main() {
	root := Constructor(4)
	root.Next = Constructor(2)
	root.Next.Next = Constructor(1)
	root.Next.Next.Next = Constructor(3)
	printListNode(insertionSortList(root))
}
