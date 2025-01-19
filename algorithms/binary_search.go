package algorithms

func BinarySearch(arr []int, target int) int {
	var length = len(arr)
	var left = 0
	var right = length - 1
	for left <= right {
		var mid = (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive() {

}
