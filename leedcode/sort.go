package leedcode

//pid 21指针 怎么把最后的节点del掉,特殊case，输入[],先生成一个空节点，最后返回head.next
//归并算法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	res := new(ListNode)
	list := res
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				list.Val = l1.Val
				l1 = l1.Next
			} else {
				list.Val = l2.Val
				l2 = l2.Next
			}
		} else if l1 != nil {
			list.Val = l1.Val
			l1 = l1.Next
		} else {
			list.Val = l2.Val
			l2 = l2.Next
		}

		if l1 != nil || l2 != nil {
			list.Next = new(ListNode)
			list = list.Next
		}
	}
	return res
}
