package linkedlist

import (
	"github.com/gsora/hid-compiler/types"
)

type Node struct {
	Prev *Node
	Data types.Token
	Next *Node
}

func NewNode(t types.Token) Node {
	newNode := Node{}
	newNode.Data = t

	return newNode
}

func (n *Node) InsertToken(t types.Token) {
	newNode := Node{}
	newNode.Data = t
	n.Next = &newNode
}
