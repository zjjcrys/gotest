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
//快速排序
func quickSort(arr *[]int,left int,rig int) {
	if left<rig {
		mid:=partition(arr,left,rig)
		quickSort(arr,left,mid-1)
		quickSort(arr,mid+1,rig)
	}
}
func partition(arr *[]int,left int,rig int) int{
	key:=(*arr)[left]
	for left<rig {
		for left<rig&&(*arr)[rig]>=key {
			rig--
		}
		(*arr)[left]=(*arr)[rig]
		for left<rig&&(*arr)[left]<=key {
			left++
		}
		(*arr)[rig]=(*arr)[left]
	}
	(*arr)[left]=key
	return left
}
