package main

import "fmt"

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// LinkedList
type LinkedList struct {
	headNode *Node
}

// NodeBetweenValues
func (l *LinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead
func (l *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.previousNode = node
	}
	l.headNode = node
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

// addAfter
func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	nodeWith := l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
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

// AddToEnd
func (l *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	lastNode := l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

// Iterate
func (l *LinkedList) IterateList() {
	var node *Node
	for node = l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}

func main() {
	myList := &LinkedList{}
	myList.AddToHead(1)
	myList.AddToHead(3)
	myList.AddToEnd(5)
	myList.AddAfter(1, 7)
	fmt.Println(myList.headNode.property)
	node := myList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
	myList.IterateList()
}
