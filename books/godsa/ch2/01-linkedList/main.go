package main

import "fmt"

// Node
type Node struct {
	property int
	nextNode *Node
}

// LinkedList
type LinkedList struct {
	headNode *Node
}

// AddToHead
func (l *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if l.headNode == nil {
		l.headNode = &node
	} else {
		node.nextNode = l.headNode
		l.headNode = &node
	}
	/*if node.nextNode != nil {
		node.nextNode = l.headNode
	}
	l.headNode = &node*/
}

// Iterate
func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}

// Last node
func (l *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// Add to the end
func (l *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	lastNode := l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	} else {
		// if lastNode is nil then list is empty
		l.headNode = node
	}
}

// NodeWithValue method returns Node given parameter property
func (l *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter
func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	// find node
	nodeWith := l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	myList := LinkedList{}
	myList.AddToHead(1)
	myList.AddToHead(2)
	myList.AddToHead(3)
	myList.IterateList()
	last := myList.LastNode()
	fmt.Println("last:", last)
	myList.AddToEnd(20)
	myList.IterateList()
	nodeValue := myList.NodeWithValue(3)
	fmt.Println("node with value:", nodeValue)
}
