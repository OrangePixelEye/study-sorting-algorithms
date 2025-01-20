package basics

import "slices"

func AppendTwoArrays(first, second []string) []string {
	return append(first, second...)
}

func RemoveFirstElement(arr []int) []int {
	return arr[1:]
}

func RemoveLastElement(arr []int) []int {
	return arr[0 : len(arr)-1]
}

func FlipArray(arr []int) []int {
	slices.Reverse(arr)
	return arr
}
