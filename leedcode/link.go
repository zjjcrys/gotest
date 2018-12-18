package leedcode

//25 没有特殊烧脑的地方，最直白的变指针
//字符串反转
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//计算长度
	length := 0
	for index := head; index != nil; index = index.Next {
		length++
	}
	circle := length / k //反转的次数
	if circle < 1 {
		return head
	}
	count := 0 //和circle 比较
	index := 0 //循环位置和k比较
	var headNew, before, beforeLeft, left *ListNode
	left = nil
	for head != nil {
		index++
		tmp := head
		head = head.Next
		tmp.Next = before
		before = tmp
		if index == 1 {
			beforeLeft = left
			left = tmp
		}

		if index == k {
			index = 0
			count++
			before = nil
			if beforeLeft != nil {
				beforeLeft.Next = tmp
			}
			if count == 1 {
				headNew = tmp
			}
			if count == circle {
				left.Next = head
				break
			}
		}
	}
	return headNew
}

//44

//68

//76

//85

//87
