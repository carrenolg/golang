package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// second act as a aux
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	/*nodeToDelete := l.head
	if nodeToDelete.data == 30 {
		fmt.Println("skip 30")
		nodeToDelete = nodeToDelete.next
	}
	if nodeToDelete.data == 20 {
		fmt.Println("skip 20")
		nodeToDelete = nodeToDelete.next
	}
	if nodeToDelete.data == 10 {
		fmt.Println("skip 10")
		nodeToDelete = nodeToDelete.next
	}
	if nodeToDelete.data == 5 {
		fmt.Println("skip 5")
		//nodeToDelete = nodeToDelete.next
	}*/
	if l.head.data == value {
		node := l.head
		node.data = node.next.data
		node.next = node.next.next
		return
	}

	for node, nodeNext := l.head, l.head.next; node != nil; node, nodeNext = node.next, node.next.next {
		if nodeNext.data == value {
			node.next = node.next.next
			break
		}
	}
}

// Show list
func (l linkedList) show() {
	for node := l.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 5}
	node2 := &node{data: 10}
	node3 := &node{data: 20}
	node4 := &node{data: 30}
	// prepend
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	fmt.Println(myList)
	myList.show()
	// delete
	fmt.Println("delete")
	myList.deleteWithValue(30)
	myList.show()
}
