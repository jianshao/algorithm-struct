package singlelist

import "errors"

type singleListNode struct {
	val  interface{}
	next *singleListNode
}

type Compare func(a, b interface{}) int
type SingleListHead struct {
	head    *singleListNode
	tail    *singleListNode
	count   int
	compare Compare
}

func Init(compare Compare) *SingleListHead {
	return &SingleListHead{
		head:    nil,
		tail:    nil,
		compare: compare,
	}
}

func initNode(val interface{}) *singleListNode {
	return &singleListNode{
		val:  val,
		next: nil,
	}
}

// add a node to the last
func (h *SingleListHead) Append(val interface{}) {
	newNode := initNode(val)
	if h.head == nil {
		h.head = newNode
		h.tail = newNode
	} else {
		h.tail.next = newNode
		h.tail = newNode
	}
	h.count++
}

// add a node to be first node
func (h *SingleListHead) AddForword(val interface{}) {
	newNode := initNode(val)
	if h.head == nil {
		h.head = newNode
		h.tail = newNode
	} else {
		newNode.next = h.head
		h.head = newNode
	}
	h.count++
}

// remove a node, it's val = val
func (h *SingleListHead) Remove(val interface{}) {
	if h.head == nil {
		return
	}

	// remove the head node
	if h.compare(h.head.val, val) == 0 {
		h.head = h.head.next
		h.count--
		return
	}

	curr := h.head
	for curr.next != nil {
		if h.compare(curr.next.val, val) == 0 {
			curr.next = curr.next.next
			h.count--
			break
		}
		curr = curr.next
	}
	// removed the tail, update the tail node
	if curr.next == nil {
		h.tail = curr
	}
}

// check the val existed or not
func (h *SingleListHead) IsExist(val interface{}) bool {
	curr := h.head
	for curr != nil {
		if h.compare(curr.val, val) == 0 {
			return true
		}
		curr = curr.next
	}
	return false
}

// get the nth node in list
func (h *SingleListHead) GetTheNthNode(nth int) (interface{}, error) {
	curr := h.head
	for nth >= 0 {
		if curr == nil {
			return nil, errors.New("nth node not exist")
		}
		curr = curr.next
		nth--
	}
	return curr.val, nil
}

func (h *SingleListHead) Pop() (interface{}, error) {
	if h.head == nil {
		return nil, errors.New("no node")
	}
	curr := h.head
	h.head = h.head.next
	h.count--
	return curr.val, nil
}

func (h *SingleListHead) GetNodesCount() int {
	return h.count
}
