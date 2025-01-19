package main

import (
	"fmt"
	"sorting/algorithms"
)

func main() {
	var result = algorithms.BinarySearchRecursiveHelper([]int{1, 2, 32, 35, 46, 55, 233}, 233)
	fmt.Println(result)
}
