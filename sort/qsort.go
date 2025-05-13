package sort

type Less func(a, b interface{}) bool // return true if a < b

// find the right place for the pivot
func findThePlace(arr []interface{}, start, end int, less Less) int {
	if start >= end {
		return -1
	}

	pivot := arr[start]
	left, right := start, end
	for left < right {
		for left < right && less(pivot, arr[right]) {
			right--
		}
		// left, right is the margin of array, just need set the pivot to the margin everytime
		arr[left] = arr[right]
		for left < right && !less(pivot, arr[left]) {
			left++
		}
		arr[left] = arr[right]
	}
	// left is equal to right, set the pivot to the final position
	arr[left] = pivot
	findThePlace(arr, start, left-1, less)
	findThePlace(arr, left+1, end, less)
	return left
}

func QSort(arr []interface{}, less Less) {
	findThePlace(arr, 0, len(arr)-1, less)
}
