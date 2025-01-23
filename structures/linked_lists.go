package structures

import (
	"errors"
	"log"
)

type LinkedListNode struct {
	data int
	prev *LinkedListNode
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	Size int
}

func NewLinkedList() LinkedList {
	return LinkedList{Size: 0, head: nil, tail: nil}
}

func (list *LinkedList) Push(val int) {
	if list.Size == 0 {
		newNode := &LinkedListNode{data: val, next: nil, prev: nil}
		list.head = newNode
		list.tail = newNode
		list.Size++
		return
	}
	list.Size++
	currentTail := list.tail
	newNode := &LinkedListNode{data: val, prev: currentTail}
	currentTail.next = newNode
	newNode.prev = currentTail
	list.tail = newNode
}

func (list *LinkedList) Pop() (int, error) {
	if list.Size == 0 {
		return 0, errors.New("Cant pop empty list")
	}

	value := list.tail.data
	list.Size--

	if list.Size == 0 {
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}

	return value, nil
}

func (list *LinkedList) RemoveFirst() (int, error) {
	if list.Size == 0 {
		return 0, errors.New("Cant pop empty list")
	}
	list.Size--
	value := list.head.data
	if list.Size == 0 {
		list.head = nil
		list.tail = nil
		return value, nil
	} else {
		list.head = list.head.next
		list.head.prev = nil
		return value, nil
	}
}

func (list *LinkedList) Has(val int) bool {
	element := list.head
	for element != nil {
		if element.data == val {
			return true
		}
		element = element.next
	}

	return false
}

func (list *LinkedList) Transverse() {
	element := list.head

	for element != nil {
		log.Println(element.data)
		element = element.next
	}
}

func (list *LinkedList) ReverseTransverse() {
	element := list.tail

	for element != nil {
		log.Println(element.data)
		element = element.prev
	}
}
