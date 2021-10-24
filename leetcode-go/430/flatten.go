/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/3
 * Time 11:42 上午
 */

/**
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//func flatten(root *Node) *Node {
//
//}

func main() {
	node1 := &Node{Val: 1}
	node1.Next = &Node{Val: 2}
	node1.Child = &Node{Val: 3}
	node1.Next.Prev = node1
	node1.Next.Next = &Node{Val: 4}
	node1.Next.Next.Prev = node1.Next

	//fmt.Println(node1)

	printNode(node1)

	//fmt.Println(flatten(node1))
}

func printNode(root *Node) {
	//可以想像成二叉树的深度优先遍历，利用栈实现
	for root != nil {
		fmt.Println(root.Val)
		root = root.Child
	}
}
