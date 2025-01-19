package main

import (
	"fmt"
	"sorting/algorithms"
)

func main() {
	var result = algorithms.BinarySearch([]int{1, 2, 32, 35, 46, 55, 233}, 46)
	fmt.Println(result)
}
