package main

import "fmt"

type DoubleNode struct {
	property     int
	previousNode *DoubleNode
	nextNode     *DoubleNode
}

type DoubleLinkedList struct {
	headNode *DoubleNode
}

func (list *DoubleLinkedList) AddToHead(property int) {
	var node = &DoubleNode{}
	node.property = property
	node.nextNode = nil

	if node.nextNode != nil {
		node.nextNode = list.headNode
		list.headNode.previousNode = node
	}
	list.headNode = node
}

func (list *DoubleLinkedList) IterateList() {
	var node *DoubleNode
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (list *DoubleLinkedList) LastNode() *DoubleNode {
	var node *DoubleNode
	var lastNode *DoubleNode

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}

	return lastNode
}

func (list *DoubleLinkedList) AddToEnd(property int) {
	var node = &DoubleNode{}
	node.property = property
	node.nextNode = nil

	var lastNode *DoubleNode
	lastNode = list.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (list *DoubleLinkedList) NodeWithValue(property int) *DoubleNode {
	var node *DoubleNode
	var nodeWith *DoubleNode

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}

	return nodeWith
}

func (list *DoubleLinkedList) AddAfter(nodeProperty int, property int) {
	var node = &DoubleNode{}
	node.property = property
	node.nextNode = nil

	var nodeWith *DoubleNode
	nodeWith = list.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (list *DoubleLinkedList) NodeBetweenValue(firstProperty int, secondProperty int) *DoubleNode {
	var node *DoubleNode
	var nodeWith *DoubleNode

	for node = list.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}

	return nodeWith
}
