package singlelist

func (l *SingleList) Reverse() {
	root := l.head
	// return if less than 2 nodes
	if root == nil || root.next == nil {
		return
	}

	// set curr.next = prev in every loop
	prev, curr, next := root, root.next, root.next.next
	for next != nil {
		curr.next = prev
		prev, curr, next = curr, next, next.next
	}
	curr.next = prev

	// begin from second node, the root node will be the last node, we should set it's next to nil
	root.next = nil
	// head, tail exchange
	l.head, l.tail = l.tail, l.head
}
