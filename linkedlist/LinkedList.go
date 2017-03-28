package linkedlist

import (
	"github.com/gsora/hid-compiler/types"
)

type LinkedList struct {
	Beginning *Node
	End       *Node
}

func NewLinkedList() LinkedList {
	l := LinkedList{}
	l.Beginning = &Node{}
	l.End = &Node{}

	l.Beginning.Next = l.End
	l.End.Prev = l.Beginning

	return l
}

func (l *LinkedList) InsertToken(t types.Token) {
	l.End.Data = t

	newEnd := Node{}
	newEnd.Prev = l.End

	l.End.Next = &newEnd

	l.End = &newEnd
}

func (l *LinkedList) GetFirst() Node {
	return *l.Beginning.Next
}

func (l *LinkedList) GetLast() Node {
	return *l.End.Prev
}
