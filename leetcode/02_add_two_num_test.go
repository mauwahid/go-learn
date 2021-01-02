package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumber(t *testing.T) {

	l1_3 := &ListNode{Val: 3}
	l1_2 := &ListNode{Val: 4, Next: l1_3}
	l1_1 := &ListNode{Val: 2, Next: l1_2}

	//	l2_4 := &ListNode{Val: 9}
	l2_3 := &ListNode{Val: 4}
	l2_2 := &ListNode{Val: 6, Next: l2_3}
	l2_1 := &ListNode{Val: 5, Next: l2_2}

	listNode := addTwoNumbers(l1_1, l2_1)

	for {
		fmt.Println("listNode val ", listNode.Val)

		if listNode.Next == nil {
			break
		}

		listNode = listNode.Next
	}
}
