package tree

import "fmt"

func (node *Node) Overview() {
	i := 0
	node.OverviewFunc(func(n *Node) {
		n.Print()
		i++
	})
	fmt.Println("All counts : ", i)
}

func (node *Node) OverviewFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.OverviewFunc(f)
	f(node)
	node.Right.OverviewFunc(f)
}
