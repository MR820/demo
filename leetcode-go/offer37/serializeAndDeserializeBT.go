/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/11/20
 * Time 3:50 下午
 */

package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func newNode(n int) *Node {
	return &Node{
		val:   n,
		left:  nil,
		right: nil,
	}
}

func serialize(root *Node) string {
	ans := make([]string, 0)
	s := make([]*Node, 0)
	s = append(s, root)
	for len(s) != 0 {
		node := s[0]
		s = s[1:]
		if node == nil {
			ans = append(ans, "nil")
			continue
		} else {
			ans = append(ans, strconv.Itoa(node.val))
		}
		s = append(s, node.left)
		s = append(s, node.right)
	}
	str := "["
	for _, v := range ans {
		str += v + ","
	}
	str = str[:len(str)-1]
	str += "]"

	return str
}

//func deserialize(data string) *Node {
//	if data == "[]" {
//		return nil
//	}
//	data = data[1 : len(data)-1]
//	arr := strings.Split(data, ",")
//	s := make([]*Node, 0)
//	head := &Node{
//		val:   arr[0],
//		left:  nil,
//		right: nil,
//	}
//	s = append(s, head)
//	for len(s) != 0 {
//		node := s[0]
//		s = s[1:]
//		node.left = &Node{
//			val:   arr[],
//			left:  nil,
//			right: nil,
//		}
//	}
//}

func main() {
	root := newNode(1)
	root.left = newNode(2)
	root.right = newNode(3)
	root.right.left = newNode(4)
	root.right.right = newNode(5)
	str := serialize(root)
	fmt.Println(str)
	//node := deserialize(str)
	//fmt.Println(node)
}
