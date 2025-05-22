package singlelist

// merge all the nodes to dest with given ordered lists. it must be same type list
func Merge(list1 *SingleListHead, list2 *SingleListHead) *SingleListHead {
	if list1 == nil && list2 == nil {
		return nil
	}

	var dest *SingleListHead
	if list1 != nil {
		dest = Init(list1.compare)
	} else {
		dest = Init(list2.compare)
	}

	currNode1, _ := list1.Pop()
	currNode2, _ := list2.Pop()
	for currNode1 != nil || currNode2 != nil {
		if currNode1 == nil {
			dest.Append(currNode2)
			currNode2, _ = list2.Pop()
		} else if currNode2 == nil {
			dest.Append(currNode1)
			currNode1, _ = list1.Pop()
		} else {
			result := dest.compare(currNode1, currNode2)
			if result > 0 {
				dest.Append(currNode1)
				currNode1, _ = list1.Pop()
			} else {
				dest.Append(currNode2)
				currNode2, _ = list2.Pop()
			}
		}
	}
	return dest
}
