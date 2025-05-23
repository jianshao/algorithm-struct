package singlelist

import "errors"

type singleListNode struct {
	val  interface{}
	next *singleListNode
}

type Compare func(a, b interface{}) int
type SingleList struct {
	head    *singleListNode
	tail    *singleListNode
	count   int
	compare Compare
}

func Init(compare Compare) *SingleList {
	return &SingleList{
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

func (l *SingleList) isValid() error {
	if l == nil {
		return errors.New("list should be init first")
	}
	return nil
}

// add a node to the last
func (l *SingleList) Append(val interface{}) error {
	if val == nil {
		return errors.New("val can not be nil")
	}
	if err := l.isValid(); err != nil {
		return err
	}
	newNode := initNode(val)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.count++
	return nil
}

// add a node to be first node
func (l *SingleList) AddForword(val interface{}) error {
	if val == nil {
		return errors.New("val can not be nil")
	}
	if err := l.isValid(); err != nil {
		return err
	}
	newNode := initNode(val)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.count++
	return nil
}

// remove a node, it's val = val
func (l *SingleList) Remove(val interface{}) error {
	if val == nil {
		return errors.New("val can not be nil")
	}
	if err := l.isValid(); err != nil {
		return err
	}

	if l.head == nil {
		return nil
	}

	// remove the head node
	if l.compare(l.head.val, val) == 0 {
		l.head = l.head.next
		l.count--
		return nil
	}

	curr := l.head
	for curr.next != nil {
		if l.compare(curr.next.val, val) == 0 {
			curr.next = curr.next.next
			l.count--
			break
		}
		curr = curr.next
	}
	// removed the tail, update the tail node
	if curr.next == nil {
		l.tail = curr
	}
	return nil
}

// check the val existed or not
func (l *SingleList) IsExist(val interface{}) (bool, error) {
	if val == nil {
		return false, errors.New("val can not be nil")
	}
	if err := l.isValid(); err != nil {
		return false, err
	}
	curr := l.head
	for curr != nil {
		if l.compare(curr.val, val) == 0 {
			return true, nil
		}
		curr = curr.next
	}
	return false, nil
}

// get the nth node in list
func (l *SingleList) GetTheNthNode(nth int) (interface{}, error) {
	if err := l.isValid(); err != nil {
		return nil, err
	}
	// support -1 ...
	if nth < 0 {
		nth = l.count + nth + 1
	}
	curr := l.head
	for nth > 1 {
		if curr == nil {
			return nil, errors.New("nth node not exist")
		}
		curr = curr.next
		nth--
	}
	return curr.val, nil
}

func (l *SingleList) Pop() (interface{}, error) {
	if err := l.isValid(); err != nil {
		return nil, err
	}
	if l.head == nil {
		return nil, errors.New("no node")
	}
	curr := l.head
	l.head = l.head.next
	l.count--
	return curr.val, nil
}

func (l *SingleList) GetNodesCount() (int, error) {
	if err := l.isValid(); err != nil {
		return 0, err
	}
	return l.count, nil
}
