package leedcode

import "sort"

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
func quickSort(arr *[]int, left int, rig int) {
	if left < rig {
		mid := partition(arr, left, rig)
		quickSort(arr, left, mid-1)
		quickSort(arr, mid+1, rig)
	}
}
func partition(arr *[]int, left int, rig int) int {
	key := (*arr)[left]
	for left < rig {
		for left < rig && (*arr)[rig] >= key {
			rig--
		}
		(*arr)[left] = (*arr)[rig]
		for left < rig && (*arr)[left] <= key {
			left++
		}
		(*arr)[rig] = (*arr)[left]
	}
	(*arr)[left] = key
	return left
}

//topic 215 堆排序

//347 利用sort 和hash,去重
func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int) //存储出现的次数
	sortArr := make([]int, 0) //存储次数用来排序

	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	for _, v := range hash {
		sortArr = append(sortArr, v)
	}
	sort.Ints(sortArr)

	flag := 0
	index := 0
	for index < k {
		if flag == sortArr[len(sortArr)-1-index] {
			index++
			continue
		}
		flag = sortArr[len(sortArr)-1-index]
		for key, v := range hash {
			if v == flag {
				ret = append(ret, key)
			}
		}
		index++
		if len(ret) > k {
			break
		}
	}
	return ret
}
