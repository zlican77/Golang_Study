package main

import (
	"fmt"
	"tree"
)

func main() {
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{1, nil, nil}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	fmt.Println(root)
	root.Print()
	root.Left.Right.Print()

	root.Right.Left.SetValue(4)

	//中序遍历
	root.Overview()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	fmt.Println(nodes[2].Right.Left.Value)
}
