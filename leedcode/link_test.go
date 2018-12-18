package leedcode

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	var node, node2, node3, node4, node5, node6, ret *ListNode
	node = new(ListNode)
	node2 = new(ListNode)
	node3 = new(ListNode)
	node4 = new(ListNode)
	node5 = new(ListNode)
	node6 = new(ListNode)
	node.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node6.Val = 6
	node.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = nil
	ret = ReverseKGroup(node, 2)
	for i := ret; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
