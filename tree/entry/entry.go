package main

import (
	"fmt"

	"imooc.com/ccmouse/learngo/tree"
)

//使用组合 (内嵌，语法糖，省下许多代码)
type myTreeNode struct {
	node *tree.Node

	//*tree.Node
}

func (myNode *myTreeNode) Traverse() {

}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	//nodes := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)
	//root.Left.Right.SetValue(5)
	//root.Left.Right.Print()
	//
	//var pRoot *tree.Node
	//pRoot.SetValue(200)
	//fmt.Println(pRoot)
	//pRoot = &root
	//fmt.Println(root)
	//pRoot.SetValue(300)
	//fmt.Println(pRoot)

	root.Traverse()
	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)

}
