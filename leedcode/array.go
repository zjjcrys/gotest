package leedcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// make 和 []int的不同
func SortArrayByParity(A []int) []int {
	length := len(A)
	ret := make([]int, length)
	i := 0
	j := length - 1
	for index := 0; index < length; index++ {
		if A[index]%2 == 0 { //偶数
			ret[i] = A[index]
			i++
		} else { //奇数
			ret[j] = A[index]
			j--
		}
	}
	return ret
}

// 熟悉二维数组的初始化
func Transpose(A [][]int) [][]int {
	row := len(A)
	column := len(A[0])
	ret := make([][]int, column)
	for i := 0; i < column; i++ {
		ret[i] = make([]int, row)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			ret[j][i] = A[i][j]
		}
	}
	return ret
}

//创建一个指向不同内存的新数组，用copy
func CombinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	tmp := []int{}
	backTracking(&ret, k, n, tmp, 0)
	return ret
}
func backTracking(ret *[][]int, k int, n int, tmp []int, sum int) {
	if len(tmp) == k {
		if sum == n {
			item := make([]int, len(tmp))
			copy(item, tmp)
			*ret = append(*ret, item)
		}
		return
	}
	var beg int
	if len(tmp) == 0 {
		beg = 1
		sum = 0
	} else {
		beg = tmp[len(tmp)-1] + 1
	}

	for i := beg; i <= 9; i++ {
		if sum+i > n {
			return
		}
		tmp = append(tmp, i)
		backTracking(ret, k, n, tmp, sum+i)
		tmp = tmp[:len(tmp)-1]
	}
}

//289 生命游戏,本地算法
func GameOfLife(board [][]int) {
	row := len(board)
	column := len(board[0])
	ret := make([][]int, row)

	for i := 0; i < row; i++ {
		ret[i] = make([]int, column)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if liveJudge(board, i, j, row, column) {
				ret[i][j] = 1
			} else {
				ret[i][j] = 0
			}
		}
	}
	copy(board, ret)
}

func liveJudge(board [][]int, i int, j int, row int, column int) bool {
	self := board[i][j]
	var top, bottom, left, right, lt, lb, rt, rb int
	if i-1 < 0 {
		top = 0
	} else {
		top = board[i-1][j]
	}

	if i+1 >= row {
		bottom = 0
	} else {
		bottom = board[i+1][j]
	}

	if j-1 < 0 {
		left = 0
	} else {
		left = board[i][j-1]
	}

	if j+1 >= column {
		right = 0
	} else {
		right = board[i][j+1]
	}

	if i-1 < 0 || j-1 < 0 {
		lt = 0
	} else {
		lt = board[i-1][j-1]
	}

	if i+1 >= row || j-1 < 0 {
		lb = 0
	} else {
		lb = board[i+1][j-1]
	}

	if i-1 < 0 || j+1 >= column {
		rt = 0
	} else {
		rt = board[i-1][j+1]
	}

	if i+1 >= row || j+1 >= column {
		rb = 0
	} else {
		rb = board[i+1][j+1]
	}

	sum := top + bottom + left + right + lt + lb + rt + rb
	if self == 0 {
		if sum == 3 {
			return true
		}
		return false
	}
	if sum < 2 {
		return false
	}
	if sum == 2 || sum == 3 {
		return true
	}
	return false
}

//229 求众数的算法，不利用排序，多数投票算法,go 数组动态增加
func MajorityElement(nums []int) []int {
	length := len(nums)
	ret := make([]int, 0)
	if length < 2 {
		return nums
	}

	var ma1, ma2, cou1, cou2 int
	ma1 = nums[0]
	ma2 = nums[1]
	cou1 = 0
	cou2 = 0
	third := length/3 + 1

	for i := 0; i < length; i++ {
		if nums[i] == ma1 {
			cou1++
			continue
		}
		if nums[i] == ma2 {
			cou2++
			continue
		}
		if cou1 == 0 {
			ma1 = nums[i]
			cou1++
			continue
		}
		if cou2 == 0 {
			ma2 = nums[i]
			cou2++
			continue
		}
		cou1--
		cou2--
	}
	cou1 = 0
	cou2 = 0
	for i := 0; i < length; i++ {
		if nums[i] == ma1 {
			cou1++
		} else if nums[i] == ma2 {
			cou2++
		}
	}
	if cou1 >= third {
		ret = append(ret, ma1)
	}
	if cou2 >= third {
		ret = append(ret, ma2)
	}
	return ret
}

//37 数独，只能回溯算法了，byte 和int 不同，加引号和不加引号不同
//优化方法，优化比较次数，利用空间存储
func solveSudoku(board [][]byte) {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	sudoKu(board, 0, 0)
	fmt.Println(board)
}
func sudoKu(board [][]byte, row int, col int) bool {
	if row == 9 {
		return true
	}

	var i byte
	for i = '1'; i <= '9'; i++ {
		if (board)[row][col] == '.' {
			if !checkValid(board, row, col, i) {
				continue
			}
			(board)[row][col] = i
			newRow := row
			newCol := col + 1
			if newCol == 9 {
				newRow++
				newCol = 0
			}
			if sudoKu(board, newRow, newCol) {
				return true
			}
			(board)[row][col] = '.'
		} else {
			newRow := row
			newCol := col + 1
			if newCol == 9 {
				newRow++
				newCol = 0
			}
			if sudoKu(board, newRow, newCol) {
				return true
			}
			break
		}

	}
	return false
}

//判断当前放的值是不是有效
func checkValid(board [][]byte, row int, col int, val byte) bool {
	//判断列
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if board[i][col] == val {
			return false
		}
	}
	//判断行
	for j := 0; j < 9; j++ {
		if j == col {
			continue
		}
		if board[row][j] == val {
			return false
		}
	}
	//判断3*3宫格
	gridRow := (row / 3) * 3
	gridCol := (col / 3) * 3
	for i := gridRow; i < gridRow+3; i++ {
		for j := gridCol; j < gridCol+3; j++ {
			if i == row && j == col {
				continue
			}
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

//44 通配符匹配又是回溯题,深度优先遍历,先把情况列举完,超时了，动态规划，找关系
//程序中直接调用的正则匹配是怎么实现的，底层算法是什么
//编译原理中涉及到的词法分析树和语法分析数是怎么遍历的
func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	if p == "*" {
		return true
	}
	if s != "" && p != "" {
		return match(s, p, 0, 0)
	}
	return false
}
func match(s string, p string, sIndex int, pIndex int) bool {
	if sIndex >= len(s) {
		if pIndex >= len(p) {
			return true
		}
		for i := pIndex; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	if sIndex < len(s) && pIndex >= len(p) {
		return false
	}

	if s[sIndex] == p[pIndex] || p[pIndex] == '?' {
		return match(s, p, sIndex+1, pIndex+1)
	}
	if p[pIndex] != '*' && s[sIndex] != p[pIndex] {
		return false
	}
	for i := sIndex; i <= len(s); i++ {
		if match(s, p, i, pIndex+1) {
			return true
		}
	}
	return false
}

//动态规划的解法，如果使用穷举超时，尝试dp优化
//自己在观念上无法把回溯推导到dp，走不通
func isMatch2(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	ret := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		ret[i] = make([]bool, pLen+1)
	}
	ret[0][0] = true
	for i := 1; i <= sLen; i++ {
		ret[i][0] = false
	}
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			ret[0][j] = ret[0][j-1]
		} else {
			ret[0][j] = false
		}
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				ret[i][j] = ret[i-1][j-1]
			} else if p[j-1] == '*' {
				ret[i][j] = ret[i-1][j] || ret[i][j-1]
			} else {
				ret[i][j] = false
			}
		}
	}

	return ret[sLen][pLen]
}

//85 找不出对应关系，怎么能从子问题推到到解
//动态规划每一步是最优解吗，二叉树对于为什么能取到最优解
//首先简化问题，怎么简化 是84问题的递增，把二维转化为一维
func maximalRectangle(matrix [][]byte) int {
	ret := 0
	if len(matrix) < 1 {
		return ret
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '0' {
				matrix[i][j] = matrix[i-1][j] + 1
			}
		}
	}

	arr := make([][]int, len(matrix))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			arr[i][j] = int(matrix[i][j]) - '0'

		}
	}
	for i := 0; i < len(arr); i++ {
		tmp := largestRectangleArea(arr[i])
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

//84 都是单调栈的使用
//https://zhuanlan.zhihu.com/p/26465701 单调栈的使用
//左边界问题
func largestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	heights = append(heights, 0)
	stack := make([]int, 0) //用来做递增栈
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			left := 0
			if len(stack) > 1 {
				left = stack[len(stack)-2] + 1
			}
			if (i-left)*height > maxArea {
				maxArea = (i - left) * height
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return maxArea
}

//146 LRU缓存 使用len cap
//如果存在直接覆盖,并且修改顺序,改为双链表
//go的双链表如何实现，在生产环境中多个线程访问，如何保证数据的一致性
type LRUCache struct {
	link []int
	hash map[int]int
	cap  int
}

/*func Constructor(capacity int) LRUCache {
	lru := new(LRUCache)
	if capacity < 1 {
		return *lru
	}
	lru.link = make([]int, 0, capacity)
	lru.hash = make(map[int]int)
	lru.cap = capacity
	return *lru
}*/

func (this *LRUCache) Get(key int) int {
	ret := -1
	if this.hash[key] == 0 {
		return ret
	}
	ret = this.hash[key]
	index := -1
	length := len(this.link)
	for i := 0; i < length; i++ {
		if this.link[i] == key {
			index = i
			break
		}
	}
	if index != -1 && index != length-1 {
		tmp := this.link[index]
		for j := index + 1; j < length; j++ {
			this.link[j-1] = this.link[j]
		}
		this.link[length-1] = tmp
	}
	return ret
}

func (this *LRUCache) Put(key int, value int) {
	var head int
	if len(this.link) > 0 {
		head = this.link[0]
	}
	if this.hash[key] != 0 {
		this.hash[key] = value
		this.Get(key)
		return
	}

	if len(this.link) >= this.cap {
		delete(this.hash, head)
		this.link = this.link[1:]
	}
	this.link = append(this.link, key)
	this.hash[key] = value
}

//299 hash的简单使用 把条件放在for循环；就直接退出l
func getHint(secret string, guess string) string {
	ret := ""
	countA := 0
	countB := 0
	if len(secret) != len(guess) || len(secret) == 0 {
		return ret
	}
	hash := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			countA++
			continue
		}
		hash[secret[i]]++
	}
	for i := 0; i < len(guess); i++ {
		if hash[guess[i]] > 0 && secret[i] != guess[i] {
			countB++
			hash[guess[i]]--
		}
	}
	ret += strconv.Itoa(countA) + "A"
	ret += strconv.Itoa(countB) + "B"
	return ret
}

//topic 1
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

//topic 10 内存怎么优化，双指针，指针迁移的时机根据题目实际来说
func maxArea(height []int) int {
	ret := 0
	if len(height) < 2 {
		return ret
	}
	left := 0
	rig := len(height) - 1

	for left < rig {
		tmp := min(height[left], height[rig]) * (rig - left)
		if ret < tmp {
			ret = tmp
		}
		if left <= rig {
			left++
		} else {
			rig--
		}
	}
	return ret
}

//qid 罗马数字转换
func romanToInt(s string) int {
	mapTable := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	special := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	ret := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && special[s[i:i+2]] > 0 {
			ret += special[s[i:i+2]]
			i++
		} else {
			ret += mapTable[s[i]]
		}

	}
	return ret
}

//topic 19 双指针 特殊case n=节点个数逻辑,n是有效的，否则逻辑需要变化
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	rig := head
	for i := 1; i <= n; i++ {
		if rig.Next != nil {
			rig = rig.Next
		} else {
			return left.Next
		}
	}
	for rig.Next != nil {
		left = left.Next
		rig = rig.Next
	}
	left.Next = left.Next.Next
	return head
}

//15 三数求和 利用双指针，固定其中一个数
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		fix := nums[i]
		if fix > 0 { //删选第一个条件
			break
		}
		if i > 0 && nums[i] == nums[i-1] { //排除重复数
			continue
		}
		left := i + 1
		rig := len(nums) - 1
		for left < rig { //只能小于，因为是三个不同的数
			if nums[left]+nums[rig] == -fix {
				tmp := []int{fix, nums[left], nums[rig]}
				res = append(res, tmp)                         //tmp重新生成一块新区域，避免被修改
				for left < rig && nums[left] == nums[left+1] { //再做一遍重复过滤
					left++
				}
				for left < rig && nums[rig] == nums[rig-1] {
					rig--
				}
				left++ //移动指针
				rig--
			} else if nums[left]+nums[rig] > -fix {
				rig--
			} else {
				left++
			}
		}

	}
	return res
}

//topic 26 双指针删除重复项
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	left := 0
	rig := 1

	for rig < len(nums) {
		if nums[left] == nums[rig] {
			rig++
		} else {
			nums = append(nums[:left+1], nums[rig:]...)
			left++
			rig = left
		}
	}
	nums = append(nums[:left+1])
	return len(nums)
}

//topic 88 双指针 注意两个是不是都跑完了
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
		index--
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		j--
		index--
	}
}

//topic 16 找三个离target最近的三个数，双指针
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for point := 0; point < len(nums)-2; point++ {
		left := point + 1
		rig := len(nums) - 1
		for left < rig {
			sum := nums[point] + nums[left] + nums[rig]
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
			if sum < target {
				left++
			} else {
				rig--
			}
		}
	}
	return res
}

//qid 18 在三层的基础上在套一层,特殊case[0,0,0,0]，三个数的时候和0比较，这个题是和target比较
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			if i > 0 && nums[i] == nums[i-1] { //排除重复数
				continue
			}
			if j > i+1 && nums[j] == nums[j-1] { //排除重复数,应该>i+1
				continue
			}
			fix := nums[i] + nums[j]
			left := j + 1
			rig := len(nums) - 1
			for left < rig { //只能小于，因为是三个不同的数
				if nums[left]+nums[rig]+fix == target {
					tmp := []int{nums[i], nums[j], nums[left], nums[rig]}
					res = append(res, tmp)                         //tmp重新生成一块新区域，避免被修改
					for left < rig && nums[left] == nums[left+1] { //再做一遍重复过滤
						left++
					}
					for left < rig && nums[rig] == nums[rig-1] {
						rig--
					}
					left++ //移动指针
					rig--
				} else if nums[left]+nums[rig]+fix > target {
					rig--
				} else {
					left++
				}
			}
		}
	}
	return res
}

//pid 27 双指针
func removeElement(nums []int, val int) int {
	left := 0
	rig := 0
	length := len(nums)
	if length < 1 {
		return left
	}

	for left < length && rig < length {
		if nums[rig] != val {
			nums[left] = nums[rig]
			left++
			rig++
		} else {
			rig++
		}
	}
	return left
}

//pid 28 特殊情况下应该返回0，注意越界问题
//kmp 原理没用上
func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}

	for left := 0; left < len(haystack); left++ {
		if haystack[left] == needle[0] {
			if left+len(needle) > len(haystack) {
				break
			}
			if needle == haystack[left:left+len(needle)] {
				return left
			}
		}
	}
	return -1
}

//pid 29 二分查找 特殊case,0 越界问题
func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	flag := true //符号位，默认
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		flag = false
	}
	//排除符号干扰
	dividend = abs(dividend)
	divisor = abs(divisor)
	res := dividend - divisor
	if res < 0 {
		return 0
	}
	index := divisor
	count := 1
	consult := 1
	for res >= divisor {
		if res > index+index {
			index += index
			count += count
		} else if index > divisor {
			for res < index {
				index = index >> 1
				count = count >> 1
			}

		}
		consult += count
		res -= index
		//fmt.Println(consult, res, index)
	}
	if !flag {
		return -consult
	}
	return consult

}
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

//先找到目标值，再从两边扩展 pid 34
func searchRange(nums []int, target int) []int {
	if len(nums)<1 {
		return []int{-1,-1}
	}

	if target<nums[0]||target>nums[len(nums)-1] {
		return []int{-1,-1}
	}
	leftRes:=-1
	rigRes:=-1
	leftIndex:=0
	rigIndex:=len(nums)-1
	find:=-1
	for leftIndex<=rigIndex {
		mid:=(leftIndex+rigIndex)/2
		if target==nums[mid] {
			find=mid
			break
		}else if target<nums[mid] {
			rigIndex=mid-1
		} else {
			leftIndex=mid+1
		}
	}
	if find==-1 {
		return []int{leftRes,rigRes}
	}
	for i:=find;i>=0;i-- {
		if nums[i]==target {
			leftRes=i
		}
	}
	for i:=find;i<len(nums);i++ {
		if nums[i]==target{
			rigRes=i
		}
	}
	return []int{leftRes,rigRes}
}
//pid 41 使用了sort，复杂度是nlog(n)
//pid 41 1放到nums[0] 2 nums[1] nums[i]==nums[nums[i]-1] 这种思想比较常见
func firstMissingPositive(nums []int) int {
	res:=1
	sort.Ints(nums)
	for i:=0;i<len(nums); i++{
		if nums[i]<=0 {
			continue
		}
		if res<nums[i] {
			break
		}
		if nums[i]==res {
			res++
		}
	}
	return res
}
//pid 81搜索旋转数组
func search(nums []int, target int) bool {
	if len(nums)<1 {
		return false
	}
	left:=0
	rig:=len(nums)-1
	for left<=rig {
		mid:=(left+rig)/2
		if nums[mid]==target {
			return true
		}
		if nums[left]<=nums[mid] { //
			if target<nums[mid] && target>=nums[left]{
				rig=mid-1
			} else {
				left=mid+1
			}
		} else {
			if target>nums[mid] && target<=nums[rig]{
				left=mid+1
			}else {
				rig=mid-1
			}
		}
	}
	return false
}
//pid 80
func removeDuplicates2(nums []int) int {
	if len(nums)<3 {
		return len(nums)
	}
	new:=1
	for old:=2;old<len(nums);old++ {
		if nums[old]==nums[new] && nums[old]==nums[new-1] {
			continue
		}
		nums[new+1]=nums[old]
		new++
	}
	return new+1
}
//pid 54 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	if len(matrix)<1 {
		return []int{}
	}
	m:=len(matrix)
	n:=len(matrix[0])
	ret:=make([]int,m*n)
	circle:=(min(m,n)+1)/2
	index:=0
	for i:=0;i<circle;i++ {
		for j:=i;j<n-i;j++ {
			ret[index]=matrix[i][j]
			index++
		}
		for j:=i+1;j<m-i&&index<m*n;j++ {
			ret[index]=matrix[j][n-i-1]
			index++
		}
		for j:=n-i-2;j>=i&&index<m*n;j-- {
			ret[index]=matrix[m-i-1][j]
			index++
		}
		for j:=m-i-2;j>i&&index<m*n;j-- {
			ret[index]=matrix[j][i]
			index++
		}
	}
	return ret
}

func myPow(x float64, n int) float64 {
	var res float64
	res=1
	for i:=n;i!=0;i/=2 {
		if i%2!=0 {
			res=res*x
		}
		x=x*x
	}
	if n>0 {
		return res
	} else {
		return 1/res
	}
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix)<1 {
		return []int{}
	}
	m:=len(matrix)
	n:=len(matrix[0])
	ret:=make([]int,m*n)
	circle:=(n+1)/2
	index:=0
	for i:=0;i<circle;i++ {
		for j:=i;j<n-i;j++ {
			ret[index]=matrix[i][j]
			index++
		}
		for j:=i+1;j<m-i&&index<m*n;j++ {
			ret[index]=matrix[j][n-i-1]
			index++
		}
		for j:=n-i-2;j>=i&&index<m*n;j-- {
			ret[index]=matrix[m-i-1][j]
			index++
		}
		for j:=m-i-2;j>i&&index<m*n;j-- {
			ret[index]=matrix[j][i]
			index++
		}
	}
	return ret
}
func generateMatrix(n int) [][]int {
	if n<1 {
		return [][]int{}
	}
	matrix:=make([][]int,n)
	for i:=0;i<n;i++ {
		matrix[i]=make([]int,n)
	}
	circle:=(n+1)/2
	index:=1
	for i:=0;i<circle;i++ {
		for j:=i;j<n-i;j++ {
			matrix[i][j]=index
			index++
		}
		for j:=i+1;j<n-i;j++ {
			matrix[j][n-i-1]=index
			index++
		}
		for j:=n-i-2;j>=i&&index<=n*n;j-- {
			matrix[n-i-1][j]=index
			index++
		}
		for j:=n-i-2;j>i&&index<=n*n;j-- {
			matrix[j][i]=index
			index++
		}
	}
	return matrix
}
//pid 187不能有重复，注意边界 用bit节省空间
func findRepeatedDnaSequences(s string) []string {
	if len(s)<11 {
		return []string{}
	}
	hash:=make(map[string]int)
	ret:=make([]string,0)
	for i:=10;i<len(s)+1;i++ {
		if hash[s[i-10:i]]==1 {
			ret=append(ret,s[i-10:i])
			hash[s[i-10:i]]++
		}else {
			hash[s[i-10:i]]++
		}
	}
	return ret
}

//445 递归解法
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil {
		return l2
	}
	if l2==nil {
		return l1
	}
	count1:=length(l1)
	count2:=length(l2)
	ret:=new(ListNode)
	if count1>=count2 {
		ret.Next=helper(l1,l2,count1-count2)
	}else {
		ret.Next=helper(l2,l1,count2-count1)
	}
	if ret.Next.Val>9 {
		ret.Next.Val%=10
		ret.Val=1
		return ret
	}
	return ret.Next
}

func length(root *ListNode) int {
	ret:=0
	if root==nil {
		return ret
	}

	for (root!=nil) {
		ret++
		root=root.Next
	}
	return ret
}

func helper(long *ListNode,short *ListNode,diff int) *ListNode{
	if long==nil {
		return nil
	}
	res:=new(ListNode)
	var post *ListNode
	if diff>0 {
		res.Val=long.Val
		post=helper(long.Next,short,diff-1)
	} else {
		res.Val=long.Val+short.Val
		post=helper(long.Next,short.Next,diff)
	}
	if post!=nil &&post.Val>9{
		post.Val%=10
		res.Val++
	}
	res.Next=post
	return res
}