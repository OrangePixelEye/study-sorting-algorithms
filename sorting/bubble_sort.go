package sorting

func BubbleSort(arr []int) {
	ordered := false
	n := len(arr)
	// Traverse through all array elements
	for i := 0; i < n-1; i++ {
		// Last i elements are already in place
		for j := 0; j < n-i-1; j++ {
			// Traverse the array from 0 to n-i-1
			// Swap if the element found is greater
			// than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				ordered = true
			}
		}

		// If no two elements were swapped by inner loop, then break
		if !ordered {
			break
		}
	}

}
