package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (list *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = list.headNode
	}
	list.headNode = &node
}

func (list *LinkedList) IterateList() {
	var node *Node
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (list *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

func (list *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node
	lastNode = list.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (list *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

func (list *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node
	nodeWith = list.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}


