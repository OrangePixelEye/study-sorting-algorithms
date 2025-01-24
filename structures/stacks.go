package structures

import "errors"

// LIFO

type StackNode struct {
	data int
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (stack *Stack) Push(element int) {
	if stack.top == nil {
		stack.top = &StackNode{data: element, next: nil}
		return
	}
	topElement := stack.top
	newElement := &StackNode{data: element, next: topElement}
	stack.top = newElement
}

func (stack *Stack) Pop() (int, error) {
	if stack.top == nil {
		return -1, errors.New("Cant pop an empty stack")
	}
	topElement := stack.top.data
	stack.top = stack.top.next
	return topElement, nil
}

func (stack Stack) Peek() (int, error) {
	if stack.top == nil {
		return -1, errors.New("Cant peek an empty stack")
	}
	return stack.top.data, nil
}
