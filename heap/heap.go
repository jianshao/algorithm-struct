package heap

import (
	"errors"
)

type Compare func(a, b interface{}) int
type Heap struct {
	count   int
	max     bool
	heap    []interface{}
	compare Compare
}

// tree style array: the children will be p*2+1 and p*2+2 if start with 0
// or p*2 and p*2+1 if start with 1
func Init(size int, max bool, compare Compare) (*Heap, error) {
	return &Heap{
		count:   0,
		max:     max,
		heap:    make([]interface{}, size),
		compare: compare,
	}, nil
}

func isValidHeap(h *Heap) error {
	if h == nil || h.heap == nil || h.compare == nil {
		return errors.New("invalid heap")
	}
	return nil
}

// from the bottom to top and replace if need
func (h *Heap) Add(val interface{}) (bool, error) {
	if err := isValidHeap(h); err != nil {
		return false, err
	}

	// need append if no place left
	if len(h.heap) <= h.count {
		h.heap = append(h.heap, val)
	} else {
		h.heap[h.count] = val
	}

	h.count++

	pos, parent := h.count-1, 0
	for pos > 0 {
		parent = (pos - 1) / 2
		if h.max {
			if h.compare(h.heap[pos], h.heap[parent]) > 0 {
				h.heap[pos], h.heap[parent] = h.heap[parent], h.heap[pos]
				pos = parent
			} else {
				break
			}
		} else {
			if h.compare(h.heap[pos], h.heap[parent]) < 0 {
				h.heap[pos], h.heap[parent] = h.heap[parent], h.heap[pos]
				pos = parent
			} else {
				break
			}
		}
	}
	// fmt.Println(h.heap...)
	return true, nil
}

// replace the last one to be first one, ensure that the array is fullsize(no leak)
// from top to bottom and replace if need
func (h *Heap) GetTheTop() (interface{}, error) {
	if err := isValidHeap(h); err != nil {
		return nil, err
	}

	if h.count == 0 {
		return nil, errors.New("no value")
	}

	// update the heap first
	result := h.heap[0]
	h.heap[0] = h.heap[h.count-1]
	h.count--

	// fixed
	pos, curr := 0, -1
	for pos < h.count {
		curr = -1
		// left child first
		if pos*2+1 < h.count {
			curr = pos*2 + 1
		}
		// check right child, find the bigger one
		if pos*2+2 < h.count && h.compare(h.heap[pos*2+2], h.heap[pos*2]) > 0 {
			curr = pos*2 + 2
		}
		// break if no child
		if curr == -1 {
			break
		}

		// no need to replace if the parent node is the greatest one
		if h.compare(h.heap[pos], h.heap[curr]) > 1 {
			break
		}
		// replace
		h.heap[pos], h.heap[curr] = h.heap[curr], h.heap[pos]
		pos = curr
	}
	// fmt.Println(result)
	// fmt.Println(h.heap...)
	return result, nil
}
