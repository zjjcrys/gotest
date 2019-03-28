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

//2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	head := tmp
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			tmp.Val = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			tmp.Val = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
		} else {
			tmp.Val = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		}
	}
	if carry > 0 {
		tmp.Next = new(ListNode)
		tmp.Next.Val = carry
	}
	return head
}
