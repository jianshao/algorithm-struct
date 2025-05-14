package sort

// find the right place for the pivot
func findTheRightPlace(arr []int, start, end int) int {
	if start >= end {
		return -1
	}

	pivot := arr[start]
	left, right := start, end
	for left < right {
		for left < right && pivot <= arr[right] {
			right--
		}
		// left, right is the margin of array, just need set the pivot to the margin everytime
		arr[left] = arr[right]
		for left < right && pivot >= arr[left] {
			left++
		}
		arr[left] = arr[right]
	}
	// left is equal to right, set the pivot to the final position
	arr[left] = pivot
	return left
}

func TopK(arr []int, k int) int {
	start, end := 0, len(arr)-1
	pos := findTheRightPlace(arr, start, end)
	for pos != k-1 {
		if pos < k {
			start = pos + 1
		} else {
			end = pos - 1
		}
		pos = findTheRightPlace(arr, start, end)
	}
	return arr[pos]
}
