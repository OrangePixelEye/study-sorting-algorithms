package algorithms

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursiveHelper(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target int, left int, right int) int {
	mid := (left + right) / 2
	for left <= right {
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			return binarySearchRecursive(arr, target, left-1, mid)
		} else {
			return binarySearchRecursive(arr, target, mid+1, right)
		}
	}
	return -1
}
