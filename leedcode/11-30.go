package leedcode

import (
	"math"
	"sort"
	"strconv"
)

/*
NO.11 if go through array use min{a[i],a[j]}*{j-1},will time limit
use two points,left=0,rig=n-1,narrow the window gradually,the premise is the smaller one move,
because width is smaller until left=rig
*/
func maxArea(height []int) int {
	res := 0
	for left, rig := 0, len(height)-1; left < rig; {
		tmp := min(height[left], height[rig]) * (rig - left)
		if res < tmp {
			res = tmp
		}
		if height[left] < height[rig] {
			left++
		} else {
			rig--
		}

	}

	return res
}
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
NO.12 go through
*/
func intToRoman(num int) string {
	hash := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	arr := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	res := ""

	for i := len(arr) - 1; i >= 0; i-- {
		if num >= arr[i] {
			tmp := num / arr[i]
			for j := 0; j < tmp; j++ {
				res += hash[arr[i]]
			}
			num = num % arr[i]
		}
	}

	return res
}

//NO.13
func romanToInt(s string) int {
	hash := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	arr := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	res := 0
	left, rig := 0, 1
	for i := len(arr) - 1; i >= 0 && left < len(s); {
		if s[left:rig] == arr[i] {
			res += hash[arr[i]]

			left++
			rig++
		} else if rig < len(s) && s[left:rig+1] == arr[i] {
			res += hash[arr[i]]
			left += 2
			rig += 2
		} else {
			i--
		}

	}
	return res
}

/*
NO.14 prefix is not substring,directly compare from beginning
*/
func longestCommonPrefix(strs []string) string {
	res := make([]byte, 0)
	if len(strs) < 2 {
		return strs[0]
	}
	//find the first two common prefix
	for i, j := 0, 0; i < len(strs[0]) && j < len(strs[1]) && strs[0][i] == strs[1][j]; i, j = i+1, j+1 {
		res = append(res, strs[0][i])
	}
	if len(res) == 0 {
		return ""
	}
	for i := 2; i < len(strs); i++ {
		j := 0
		for ; j < len(res) && j < len(strs[i]) && strs[i][j] == res[j]; j++ {
		}
		if j == 0 {
			return ""
		}
		res = res[:j]
	}
	return string(res)
}

/*
NO.15 two points,O(n*n)
notice1:repeated nums use hash or extract the repeated
notice2:have multi must go through over
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	hash := make(map[string]bool)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, rig := i+1, len(nums)-1
		for left < rig {
			if nums[left]+nums[rig] == 0-nums[i] {
				tmp := strconv.Itoa(nums[i]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[rig])
				if _, ok := hash[tmp]; !ok {
					res = append(res, []int{nums[i], nums[left], nums[rig]})
					hash[tmp] = true
				}
				left++
				rig--
			} else if nums[left]+nums[rig] < 0-nums[i] {
				left++
			} else {
				rig--
			}
		}
	}
	return res

}

/*
NO.16 similar toNO.15, //delete some possibility
*/
func threeSumClosest(nums []int, target int) int {
	res := 0
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	res = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		fix := nums[i]

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		rig := len(nums) - 1
		for left < rig {
			tmp := fix + nums[left] + nums[rig]

			//for left < rig && nums[left] == nums[left+1] {
			//	left++
			//}
			//for left < rig && nums[rig] == nums[rig-1] {
			//	rig--
			//}
			if math.Abs(float64(tmp-target)) < math.Abs(float64(res-target)) {
				res = tmp
			}
			if tmp < target {
				left++
			} else if tmp > target {
				rig--
			} else {
				break
			}
		}

	}
	return res
}

//NO.17 brute force
func letterCombinations(digits string) []string {
	hash := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	res := make([]string, 0)
	if len(digits) < 1 {
		return res
	}
	before := hash[digits[0]]
	for i := 1; i < len(digits); i++ {
		tmp := hash[digits[i]]
		for k := 0; k < len(before); k++ {
			for j := 0; j < len(tmp); j++ {
				res = append(res, before[k]+tmp[j])
			}
		}
		before = res
		res = []string{}
	}

	return before
}

/*
NO.18 1.can't set nums[l]>target,for have negative nums.
2.delete repeated result for used nums
*/
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	res := make([][]int, 0)
	sort.Ints(nums)

	for l := 0; l < len(nums)-3; l++ {
		for l > 0 && l < len(nums) && nums[l] == nums[l-1] {
			l++
		}
		for i := l + 1; i < len(nums)-2; i++ {
			for i > l+1 && i < len(nums) && nums[i] == nums[i-1] {
				i++
			}

			for left, rig := i+1, len(nums)-1; left < rig; {
				tmp := nums[l] + nums[i] + nums[left] + nums[rig]
				if tmp == target {
					res = append(res, []int{nums[l], nums[i], nums[left], nums[rig]})
					left++
					for left < len(nums) && nums[left] == nums[left-1] {
						left++
					}
					rig--
					for rig >= 0 && nums[rig] == nums[rig+1] {
						rig--
					}
				} else if tmp < target {
					left++
				} else {
					rig--
				}

			}

		}
	}
	return res
}

//NO.19directly writing
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	before := head
	tmp := 0
	for before != nil {
		tmp++
		before = before.Next
	}
	index := tmp - n + 1
	if index == 1 {
		return head.Next
	}
	tmp = 1
	before = head
	for before != nil && tmp < index-1 {
		tmp++
		before = before.Next
	}
	before.Next = before.Next.Next
	return head
}

/*
NO.20 case is go out default
notice1:bound condition
*/
func isValid(s string) bool {
	arr := make([]byte, 10000)
	index := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			arr[index] = s[i]
			index++
		case ')':
			if index < 1 || arr[index-1] != '(' {
				return false
			}
			index--
		case ']':
			if index < 1 || arr[index-1] != '[' {
				return false
			}
			index--
		case '}':
			if index < 1 || arr[index-1] != '{' {
				return false
			}
			index--
		}

	}
	if index > 0 {
		return false
	}
	return true
}

//NO.21 merge
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	head := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp := new(ListNode)
			tmp.Val = list1.Val
			list1 = list1.Next
			head.Next = tmp
			head = head.Next
		} else {
			tmp := new(ListNode)
			tmp.Val = list2.Val
			list2 = list2.Next
			head.Next = tmp
			head = head.Next
		}
	}
	if list1 != nil {
		head.Next = list1
	} else if list2 != nil {
		head.Next = list2
	}

	return res.Next
}

//No.22 directly writing
func generateParenthesis(n int) []string {
	hash := make(map[string]bool)
	before := []string{"()"}
	after := make([]string, 0)

	for i := 1; i < n; i++ {
		for j := 0; j < len(before); j++ {
			for k := 0; k < len(before[j]); k++ {
				tmp := before[j][0:k] + "()" + before[j][k:]
				if _, ok := hash[tmp]; !ok {
					after = append(after, tmp)
					hash[tmp] = true
				}

			}
		}
		before = after
		after = make([]string, 0)
	}

	return before
}

/*
NO.23 algorithm 6.5-8 use heap&&merge
Notice1:lists=nil or lists[i]=nil
*/
type heapNode struct {
	Val      int
	Location int
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]heapNode, 0)
	//create min root heap
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap = append(heap, heapNode{lists[i].Val, i})
			lists[i] = lists[i].Next
		}
	}
	for i := (len(heap) - 1) / 2; i >= 0; i-- {
		minHeapIfy(&heap, i)
	}

	res := new(ListNode)
	head := res
	//merge
	for len(heap) > 0 {
		node := new(ListNode)
		node.Val = heap[0].Val
		head.Next = node
		head = head.Next

		if lists[heap[0].Location] != nil {
			heap[0].Val = lists[heap[0].Location].Val
			lists[heap[0].Location] = lists[heap[0].Location].Next
		} else {
			heap[0] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		}
		minHeapIfy(&heap, 0)
	}

	return res.Next
}

func minHeapIfy(arr *[]heapNode, i int) {
	left, rig, min := 2*i, 2*i+1, 0
	if left < len(*arr) && (*arr)[left].Val < (*arr)[i].Val {
		min = left
	} else {
		min = i
	}
	if rig < len(*arr) && (*arr)[rig].Val < (*arr)[min].Val {
		min = rig
	}
	if min != i {
		tmpNode := (*arr)[min]
		(*arr)[min] = (*arr)[i]
		(*arr)[i] = tmpNode
		minHeapIfy(arr, min)
	}
}

//No.24 directly writing
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := new(ListNode)
	res.Next = head

	node1 := res
	node2 := head
	for node2 != nil && node2.Next != nil {
		node1.Next = node2.Next
		node1 = node1.Next
		node2.Next = node1.Next
		node1.Next = node2
		node1 = node2
		node2 = node2.Next
	}

	return res.Next
}

//NO.25 once pass
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	llen := 0
	n1 := head
	for n1 != nil {
		llen++
		n1 = n1.Next
	}
	round := llen / k
	res := new(ListNode)
	res.Next = head
	left := res
	for round > 0 {
		n1 := left.Next
		n2 := n1.Next
		index := 1
		var tmp *ListNode
		for index < k {
			tmp = n2.Next
			n2.Next = n1
			n1 = n2
			n2 = tmp
			index++
		}
		tmp = left.Next
		left.Next.Next = n2
		left.Next = n1
		left = tmp
		round--
	}

	return res.Next
}

//NO.29 use << or >> notice bound
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return dividend
	}
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	res := 0
	count := 1
	symbol := true

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		symbol = false
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	multiDivisor := divisor
	for dividend >= divisor {
		if dividend > multiDivisor {
			dividend -= multiDivisor
			res += count
			multiDivisor += multiDivisor
			count += count
		} else if dividend == multiDivisor {
			dividend -= multiDivisor
			res += count
			break
		} else {
			if multiDivisor > divisor { //use >>
				multiDivisor = multiDivisor >> 1
				count = count >> 1
			}
		}
	}
	if !symbol {
		res = -res
	}
	return res
}

/*
NO.30 使用了很多方法去优化，总是无法兼顾所有的case，原来思路有问题
https://leetcode.com/problems/substring-with-concatenation-of-all-words/discuss/1753357/Clear-solution!-Easy-to-understand-with-diagrams%5C
*/
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	hash := make(map[string][]int, 0)
	for i := 0; i < len(words); i++ {
		hash[words[i]] = append(hash[words[i]], -1)
	}
	k := len(words[0]) //k=length of words[0]
	wlen := len(words)
	for i := 0; i < k; i++ {
		subString(hash, i, k, wlen, &res, s)
		//init hash
		for l := 0; l < len(words); l++ {
			for j := 0; j < len(hash[words[l]]); j++ {
				hash[words[l]][j] = -1
			}

		}
	}

	return res
}
func subString(hash map[string][]int, left int, k int, wlen int, res *[]int, s string) {
	rig := left
	for rig <= len(s)-k {
		value, ok := hash[s[rig:rig+k]]
		if !ok {
			if rig-left >= wlen*k {
				*res = append(*res, left)
			}
			left = rig + k
			rig = left
			continue
		}

		//deal with repeated word
		update := 0
		for i := 1; i < len(value); i++ {
			if value[i] < value[update] {
				update = i
			}
		}

		if value[update] < left {
			hash[s[rig:rig+k]][update] = rig
			if rig-left >= wlen*k {
				*res = append(*res, left)
				left = rig + k
				rig = left
			} else {
				rig += k
			}

		} else {
			tmp := value[update]
			hash[s[rig:rig+k]][update] = rig
			if rig-left >= wlen*k {
				*res = append(*res, left)
			}
			left = tmp + k
			rig += k
		}

	}
	if rig-left >= wlen*k {
		*res = append(*res, left)
	}
}
