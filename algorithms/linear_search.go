package algorithms

func LinearSearch(elements []int, target int) int {
	for index, element := range elements {
		if element == target {
			return index
		}
	}
	return -1

}
