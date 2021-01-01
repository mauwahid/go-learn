package main

import (
	"container/heap"
	"fmt"
)

func main() {
	runLinkedList()
}

func runHeapSample() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum : %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}

func runTupleSample() {
	square, cube := powerSeries(4)
	fmt.Println("Square ", square, "Cube ", cube)
}

func runLinkedList() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	//linkedList.AddToHead(3)

	fmt.Println("property: ", linkedList.headNode.property)

}
