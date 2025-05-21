package heap

func intCompare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func HeapSort(arr []int) error {
	heap, err := Init(len(arr), false, intCompare)
	if err != nil {
		return err
	}

	// push to heap
	for _, element := range arr {
		err := heap.Add(element)
		if err != nil {
			return err
		}
	}

	// pop from heap, it will be ordered
	for i := 0; i < len(arr); i++ {
		val, err := heap.GetTheTop()
		if err != nil {
			return err
		}
		arr[i] = val.(int)
	}
	return nil
}
