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

func (list *LinkedList) Push(val int) {
	if list.Size == 0 {
		list.head = &LinkedListNode{data: val, next: nil, prev: nil}
	}
	currentTail := list.tail

	newNode := &LinkedListNode{data: val, prev: currentTail}

	newNode.prev = currentTail
	currentTail.next = newNode
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
