package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodeToNumber(head *ListNode) uint64 {
	var ret uint64 = 0
	cnt := 0
	for node := head; node != nil; node = node.Next {
		ret = uint64(node.Val)*uint64(math.Pow10(cnt)) + ret
		cnt++
	}
	return ret
}

func numberToNode(num uint64) *ListNode {
	if num == 0 {
		return &ListNode{0, nil}
	}

	var head, pre *ListNode = nil, nil

	for ; num > 0; num = num / 10 {
		now := &ListNode{int(num % 10), nil}
		if head == nil {
			head = now
		} else {
			pre.Next = now
		}
		pre = now
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, pre *ListNode = nil, nil
	flag := 0
	for l1 != nil || l2 != nil {
		tmp := 0
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		tmp += flag

		now := &ListNode{tmp % 10, nil}

		if head == nil {
			head = now
		} else {
			pre.Next = now
		}

		flag = tmp / 10
		pre = now
	}

	if flag > 0 {
		pre.Next = &ListNode{flag, nil}
	}

	return head
}

func main() {
	ret := addTwoNumbers(numberToNode(81), numberToNode(0))
	fmt.Println(nodeToNumber(ret))
	ret = addTwoNumbers(numberToNode(5), numberToNode(5))
	fmt.Println(nodeToNumber(ret))
}
