package compiler

type node struct {
	Prev *node
	Data token
	Next *node
}

func newNode(t token) node {
	newNode := node{}
	newNode.Data = t

	return newNode
}

func (n *node) insertToken(t token) {
	newNode := node{}
	newNode.Data = t
	n.Next = &newNode
}
