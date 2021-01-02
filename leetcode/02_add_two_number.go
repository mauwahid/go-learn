//https://leetcode.com/problems/add-two-numbers/
package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 8 ms, faster than 89.47% of Go online submissions for Add Two Numbers.
Memory Usage: 6.1 MB, less than 5.09% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sliceNode1 []int
	var sliceNode2 []int
	var newSlice []int

	node1 := l1
	node2 := l2

	for {
		sliceNode1 = append(sliceNode1, node1.Val)

		if node1.Next == nil {
			break
		}

		node1 = node1.Next
	}

	for {
		sliceNode2 = append(sliceNode2, node2.Val)

		if node2.Next == nil {
			break
		}

		node2 = node2.Next
	}

	node1Len := len(sliceNode1)
	node2Len := len(sliceNode2)

	if node1Len >= node2Len {

		j := 0
		additional := 0

		for i := 0; i < node1Len; i++ {

			temp := sliceNode1[i] + additional

			additional = 0

			if j < node2Len {
				temp = temp + sliceNode2[j]
				j++
			}

			tempRes := temp

			if temp >= 10 {
				additional = 1
				tempRes = temp % 10
			}

			newSlice = append(newSlice, tempRes)
		}

		if additional > 0 {
			newSlice = append(newSlice, additional)
		}
	}

	if node2Len > node1Len {

		j := 0
		additional := 0

		for i := 0; i < node2Len; i++ {

			temp := sliceNode2[i] + additional

			additional = 0

			if j < node1Len {
				temp = temp + sliceNode1[j]
				j++
			}

			tempRes := temp

			if temp >= 10 {
				additional = 1
				tempRes = temp % 10
			}

			newSlice = append(newSlice, tempRes)
		}

		if additional > 0 {
			newSlice = append(newSlice, additional)
		}
	}

	var nextNode *ListNode

	for x := len(newSlice) - 1; x >= 0; x-- {

		temp := new(ListNode)
		temp.Val = newSlice[x]
		temp.Next = nextNode

		nextNode = temp
	}

	return nextNode
}

/*
Runtime: 12 ms, faster than 61.39% of Go online submissions for Add Two Numbers.
Memory Usage: 5 MB, less than 93.94% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	//	p, q := l1, l2
	carry := 0
	var curr, head *ListNode

	for l1 != nil || l2 != nil {

		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
		}

		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum / 10

		temp := &ListNode{Val: sum % 10}

		if curr == nil {
			head = temp
			curr = head
		} else {
			curr.Next = temp
			curr = temp
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		fmt.Print(curr.Val)

	}

	if carry > 0 {
		nodeCarry := &ListNode{Val: carry}
		curr.Next = nodeCarry
		fmt.Print(curr.Val)
	}

	return head
}
