package leedcode

import (
	"fmt"
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
