package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("setting Value to nil node.ignored")
		return
	}
	node.Value = Value
}
