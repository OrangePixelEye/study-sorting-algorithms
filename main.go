package main

import (
	"fmt"
	"sorting/algorithms"
	"sorting/basics"
	"sorting/structures"
)

func main() {
	list := structures.NewLinkedList()
	list.Push(2)
	list.Push(4)
	list.Push(5)
	list.RemoveFirst()
	list.ReverseTransverse()

	var str = basics.RemoveAllSymbols("dsdsaf asdf @122")
	fmt.Println(str)

	var result = algorithms.BinarySearchRecursiveHelper([]int{1, 2, 32, 35, 46, 55, 233}, 233)
	fmt.Println(result)

	var comb = algorithms.BasicCombinations([]string{"a", "b", "c"})
	fmt.Println(comb)

	var appendArr = basics.AppendTwoArrays([]string{"DEFAULT", "text"}, []string{"loren", "ipsun"})
	fmt.Println(appendArr)

	var elements [][]int
	elements = [][]int{{
		1, 5, 4, 3,
	}, {
		12, 45, 212,
	}}
	fmt.Println(basics.RemoveFirstElement(elements[0]))
	fmt.Println(basics.RemoveLastElement(elements[1]))

	fmt.Println(basics.FlipArray(elements[0]))
}
