package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

func (node *Node) SetValue(Value int) {
	node.Value = Value
}
