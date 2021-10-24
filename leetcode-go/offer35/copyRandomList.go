/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 3:18 下午
 */

/**
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。



示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	//复制节点，插入原节点和下一个节点之间
	//遍历原节点，找Random。复制节点Random
	cur, next := head, head
	//Next方向已经串起来了
	for cur != nil {
		next = cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}
	//Random方向
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		} else {
			cur.Next.Random = nil
		}
		cur = cur.Next.Next
	}
	//分离原节点和复制节点 split
	pre := head.Next
	cur = head
	for cur != nil {
		next = cur.Next.Next
		if next != nil {
			cur.Next.Next = next.Next
		} else {
			cur.Next.Next = nil
		}
		cur.Next = next
		cur = next
	}
	return pre
}

func main() {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Next.Next = &Node{Val: 3}
	head.Next.Next.Next = &Node{Val: 4}

	head.Random = head.Next.Next.Next           //1->4
	head.Next.Random = head                     //2->1
	head.Next.Next.Random = head.Next.Next.Next //3->4
	head.Next.Next.Next.Random = head.Next      //4->2

	printLink(head)

	fmt.Println("复制节点：")
	printLink(copyRandomList(head))
	fmt.Println("原节点")
	printLink(head)
}

func printLink(head *Node) {
	cur := head
	fmt.Print("Next:")
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
	cur = head
	fmt.Print("Random:")
	for cur != nil {
		fmt.Printf("%d ", cur.Random.Val)
		cur = cur.Next
	}
	fmt.Println()
}
