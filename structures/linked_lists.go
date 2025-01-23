package structures

import "log"

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

func (list *LinkedList) Remove() {

}

func (list *LinkedList) Search(val int) int {
	return -1
}

func (list *LinkedList) Transverse() {
	element := list.head

	for element != nil {
		log.Println(element.data)
		element = element.next
	}
}

func (list *LinkedList) ReverseTransverse() {

}
