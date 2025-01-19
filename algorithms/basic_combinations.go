package algorithms

func BasicCombinations(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{{}}
	}

	firstElement := arr[0]
	elementsWithoutFirst := arr[1:]

	combinationsWithoutFirst := BasicCombinations(elementsWithoutFirst)

	var combWithFirstElement [][]string
	for _, combination := range combinationsWithoutFirst {
		// needs to change order here to create a new array, cause combination can take more elements
		combWithFirst := append([]string{firstElement}, combination...)
		combWithFirstElement = append(combWithFirstElement, combWithFirst)
	}

	result := append(combWithFirstElement, combinationsWithoutFirst...)
	return result
}
