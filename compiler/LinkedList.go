package compiler

type linkedList struct {
	Beginning *node
	End       *node
}

func newLinkedList() linkedList {
	l := linkedList{}
	l.Beginning = &node{}
	l.End = &node{}

	l.Beginning.Next = l.End
	l.End.Prev = l.Beginning

	return l
}

func (l *linkedList) insertToken(t token) {
	l.End.Data = t

	newEnd := node{}
	newEnd.Prev = l.End

	l.End.Next = &newEnd

	l.End = &newEnd
}

func (l *linkedList) getFirst() node {
	return *l.Beginning.Next
}

func (l *linkedList) getLast() node {
	return *l.End.Prev
}
