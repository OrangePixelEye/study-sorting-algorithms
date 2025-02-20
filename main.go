package main

import (
	"fmt"
	"sorting/algorithms"
	"sorting/basics"
	"sorting/structures"
  "sorting/exercises/sliding_technique"
)

func main() {
  fmt.Println(slidingtechnique.LongestSubstring("abcjkabcbb"))

  fmt.Println(slidingtechnique.ContainsDuplicate([]int{1, 2,  3, 4, 3}, 2))

  maxSum := slidingtechnique.MaxSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
  fmt.Println(maxSum)

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

	myStack := structures.NewStack()
	myStack.Push(2)
	myStack.Push(4)
	fmt.Println(myStack.Peek())
	myStack.Pop()

	fmt.Println(myStack.Peek())
	myStack.Pop()
	fmt.Println(myStack.Peek())
}
