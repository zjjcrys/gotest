package leedcode

import (
	"math"
	"sort"
	"strconv"
)

//No.31 find regular
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	smaller := len(nums) - 2
	for ; smaller >= 0; smaller-- {
		if nums[smaller] < nums[smaller+1] {
			break
		}
	}
	if smaller < 0 {
		for left, rig := 0, len(nums)-1; left < rig; left, rig = left+1, rig-1 {
			tmp := nums[left]
			nums[left] = nums[rig]
			nums[rig] = tmp
		}
		return
	}
	left := smaller + 1
	rig := len(nums) - 1

	larger := len(nums) - 1
	for nums[larger] <= nums[smaller] {
		larger--
	}
	tmp := nums[larger]
	nums[larger] = nums[smaller]
	nums[smaller] = tmp
	for left < rig {
		tmp = nums[left]
		nums[left] = nums[rig]
		nums[rig] = tmp
		left++
		rig--
	}
}

//No.32 use stack notice: where to put judge
func longestValidParentheses(s string) int {
	stack := make([]int, len(s)) //specify n optimize memory
	sLen := 0
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			if sLen < 1 {
				if res < i {
					res = i
				}
			} else {
				if res < i-stack[sLen-1]-1 {
					res = i - stack[sLen-1] - 1
				}
			}
			stack[sLen] = i
			sLen++
		case ')':
			if sLen < 1 {
				if res < i {
					res = i
				}
				stack[sLen] = i
				sLen++
			} else {
				if s[stack[sLen-1]] == '(' {
					sLen--
				} else {
					if res < i-stack[sLen-1]-1 {
						res = i - stack[sLen-1] - 1
					}
					stack[sLen] = i
					sLen++
				}

			}

		}
	}
	if sLen < 1 {
		res = len(s)
	} else {
		if res < len(s)-stack[sLen-1]-1 {
			res = len(s) - stack[sLen-1] - 1
		}
	}

	return res
}

//No.33 binary search:distinct values,notice =
func search(nums []int, target int) int {
	left, rig := 0, len(nums)-1
	var mid int
	res := -1
	for left <= rig {
		mid = (left + rig) / 2
		if nums[mid] == target {
			res = mid
			break
		} else if target < nums[mid] {
			if nums[left] <= nums[mid] && target < nums[left] {
				left = mid + 1
			} else {
				rig = mid - 1
			}
		} else {
			if nums[rig] >= nums[mid] && nums[rig] < target {
				rig = mid - 1
			} else {
				left = mid + 1
			}

		}
	}
	return res
}

//No.34 binary search
func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	if nums[0] == nums[len(nums)-1] && nums[0] == target {
		return []int{0, len(nums) - 1}
	}
	index := -1
	left, rig := 0, len(nums)-1
	for left <= rig {
		mid := (left + rig) / 2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			rig = mid - 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}

	left = index - 1
	rig = index + 1
	for left >= 0 && nums[left] == target {
		left--
	}
	for rig < len(nums) && nums[rig] == target {
		rig++
	}
	return []int{left + 1, rig - 1}
}

//No.35 binary search
func searchInsert(nums []int, target int) int {
	left, rig := 0, len(nums)-1
	res := -1
	for left <= rig {
		mid := (left + rig) / 2
		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			rig = mid - 1
		}
	}
	if res != -1 {
		return res
	}
	return left
}

//NO.36 represent sodoku
func isValidSudoku(board [][]byte) bool {
	hash := make(map[byte]int)
	res := true
	//check row
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			tmp := board[i][j] - '0'
			if tmp < 0 || tmp > 9 {
				res = false
				break
			}
			if value, ok := hash[tmp]; ok && value == i {
				res = false
				break
			}
			hash[tmp] = i
		}
		if !res {
			break
		}
	}
	if !res {
		return res
	}
	//check column
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			tmp := board[j][i] - '0'
			if value, ok := hash[tmp]; ok && value == 9+i {
				res = false
				break
			}
			hash[tmp] = 9 + i
		}
		if !res {
			break
		}
	}
	if !res {
		return res
	}
	//check sub-box
	for k := 0; k < 9; k++ {
		for i := (k / 3) * 3; i < k/3*3+3; i++ {
			for j := k % 3 * 3; j < k%3*3+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				tmp := board[i][j] - '0'
				if value, ok := hash[tmp]; ok && value == 18+k {
					res = false
					break
				}
				hash[tmp] = 18 + k
			}
			if !res {
				break
			}
		}
		if !res {
			break
		}
	}
	return res
}

/*
No.37 use backTrace as a whole,the key is check sudoku quickly.
version1:check every num,time limit
version2:only check num from .,check range from 1~row to 1~9
version3:use hash,but in backTrace no reset when column=0
version4:use array as hash, to minimum space,use byte.first I don't realize can use so much space
*/
func solveSudoku(board [][]byte) {
	var hashRow, hashColumn [9][9]bool
	var hashBlock [3][3][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				digit := board[i][j] - '0' - 1
				hashRow[i][digit] = true
				hashColumn[j][digit] = true
				hashBlock[i/3][j/3][digit] = true
			}
		}
	}
	backTrace37(&board, 0, 0, hashRow, hashColumn, hashBlock)
}
func backTrace37(board *[][]byte, row int, column int, hashRow, hashColumn [9][9]bool, hashBlock [3][3][9]bool) bool {
	if column == 9 {
		column = 0
		row++
	}

	if row == 9 {
		return true
	}

	if (*board)[row][column] == '.' {
		for index := byte('1'); index <= '9'; index++ {
			digit := index - '1'
			if !hashRow[row][digit] && !hashColumn[column][digit] && !hashBlock[row/3][column/3][digit] {
				(*board)[row][column] = index
				hashRow[row][digit] = true
				hashColumn[column][digit] = true
				hashBlock[row/3][column/3][digit] = true
				if backTrace37(board, row, column+1, hashRow, hashColumn, hashBlock) {
					return true
				}
				hashRow[row][digit] = false
				hashColumn[column][digit] = false
				hashBlock[row/3][column/3][digit] = false
				(*board)[row][column] = '.'
			}
		}
	} else {
		if backTrace37(board, row, column+1, hashRow, hashColumn, hashBlock) {
			return true
		}
	}
	return false
}

//No.38
func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return res
	}
	for i := 2; i <= n; i++ {
		left, rig := 0, 1
		tmp := ""
		for rig < len(res) {
			if res[left] == res[rig] {
				rig++
			} else {
				tmp += strconv.Itoa(rig-left) + string(res[left])
				left = rig
				rig++
			}
		}
		if left < len(res) {
			tmp += strconv.Itoa(rig-left) + string(res[left])
		}
		res = tmp

	}
	return res
}

//No.39 backtrace similar to 40,46,47,51,52
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backTrace39(&res, candidates, target, 0, []int{}, 0)
	return res
}
func backTrace39(res *[][]int, nums []int, target int, sum int, tmp []int, index int) {
	if sum == target {
		tmp1 := make([]int, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
		return
	}
	//nums is distinct
	for i := index; i < len(nums); i++ {
		if sum+nums[i] > target {
			break
		}
		j := (target - sum) / nums[i]

		for k := 0; k < j; k++ {
			tmp = append(tmp, nums[i])
		}
		for ; j > 0; j-- {
			backTrace39(res, nums, target, sum+nums[i]*j, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

//No.40 brute force similar to 46,47,51,52
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	trace(&res, candidates, target, 0, []int{}, 0)
	return res
}
func trace(res *[][]int, nums []int, target int, sum int, tmp []int, index int) {
	if sum == target {
		tmp1 := make([]int, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
		return
	}
	for i := index; i < len(nums); i++ {
		if sum+nums[i] > target {
			break
		}
		tmp = append(tmp, nums[i])
		trace(res, nums, target, sum+nums[i], tmp, i+1)
		tmp = tmp[:len(tmp)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

	}
}

//No.41 for 1<=nums[i]<=lem(nums) put it to correct places
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 || nums[i] == 0 {
			continue
		}
		index := i
		for nums[index] != index+1 {
			tmp := nums[index]
			if tmp > len(nums) || tmp < 1 {
				nums[index] = 0
				break
			} else {
				if nums[tmp-1] == tmp {
					break
				} else {
					nums[index] = nums[tmp-1]
					nums[tmp-1] = tmp
				}

			}
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = i + 1
			break
		}
	}
	if res == 0 {
		res = len(nums) + 1
	}
	return res
}

//No.42 monotonic stack,two points
func trap(height []int) int {
	stack := make([]int, len(height))
	res, sLen := 0, 0

	if len(height) < 2 {
		return res
	}

	stack[sLen] = 0
	sLen++
	for i := 1; i < len(height); i++ {
		if height[i] <= height[stack[sLen-1]] {
			stack[sLen] = i
			sLen++
			continue
		}

		for height[i] > height[stack[sLen-1]] {
			bottom := height[stack[sLen-1]]
			sLen--
			if sLen < 1 {
				break
			}
			min := height[i]
			if min > height[stack[sLen-1]] {
				min = height[stack[sLen-1]]
			}
			res += (i - stack[sLen-1] - 1) * (min - bottom)
		}
		stack[sLen] = i
		sLen++

	}
	return res
}

//No.43 math big data
func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	for i := len(num2) - 1; i >= 0; i-- {
		tmp := make([]byte, len(num1)+len(num2))
		tLen := 0
		carry := uint8(0)
		for k := 1; k <= len(num2)-1-i; k++ {
			tmp[tLen] = 0
			tLen++
		}
		for j := len(num1) - 1; j >= 0; j-- {
			produce := (num2[i]-'0')*(num1[j]-'0') + carry
			tmp[tLen] = produce % 10
			tLen++
			carry = produce / 10
		}
		tmp[tLen] = carry
		tLen++

		carry = 0
		for j := 0; j < tLen; j++ {
			sum := res[j] + tmp[j] + carry
			res[j] = sum % 10
			carry = sum / 10
		}
		for carry > 0 {
			sum := res[tLen] + carry
			res[tLen] = sum % 10
			tLen++
			carry = sum / 10
		}
	}
	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}
	rig := len(res) - 1
	for ; rig >= 0; rig-- {
		if res[rig] != '0' {
			break
		}
	}
	res = res[:rig+1]
	for left := 0; left < rig; {
		tmp := res[left]
		res[left] = res[rig]
		res[rig] = tmp
		left++
		rig--
	}
	return string(res)
}

//No.44 dynamic similar to 10
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ { //var dp [m + 1][n + 1]bool=error writing for go
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
				continue
			}
			if p[j] == '*' {
				if dp[i][j+1] {
					dp[i+1][j+1] = dp[i][j+1]
				} else if dp[i+1][j] {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}

		}
	}
	return dp[m][n]
}

//NO.45 dynamic
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			tmp := j + i
			if tmp < len(nums) {
				if dp[tmp] == 0 || dp[tmp] > dp[i]+1 {
					dp[tmp] = dp[i] + 1
				}
			} else {
				break
			}
		}

	}
	return dp[len(nums)-1]
}

/*
No.46 permutation notice 1:slice,map copy pointer, modify will impact,append not
n2:res append tmp1, if not use new memory, may lead insanity,tmps point to the same memory
*/
func permute(nums []int) [][]int {
	isVisited := make(map[int]bool)
	res := make([][]int, 0)
	permuteTraverse(&res, isVisited, nums, 0, []int{})
	return res
}
func permuteTraverse(res *[][]int, isVisited map[int]bool, nums []int, n int, tmp []int) {
	if n == len(nums) {
		tmp1 := make([]int, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(isVisited)[nums[i]] {
			tmp = append(tmp, nums[i])
			isVisited[nums[i]] = true
			permuteTraverse(res, isVisited, nums, n+1, tmp)
			(isVisited)[nums[i]] = false
			tmp = (tmp)[:len(tmp)-1]
		}
	}
}

//No.47 permutation isvisited store place:i,not nums[i],jump repeated nums
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	isVisited := make(map[int]bool)
	res := make([][]int, 0)
	permuteTraverse2(&res, isVisited, nums, 0, []int{})
	return res
}
func permuteTraverse2(res *[][]int, isVisited map[int]bool, nums []int, n int, tmp []int) {
	if n == len(nums) {
		tmp1 := make([]int, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(isVisited)[i] {
			tmp = append(tmp, nums[i])
			isVisited[i] = true
			permuteTraverse(res, isVisited, nums, n+1, tmp)
			isVisited[i] = false
			tmp = (tmp)[:len(tmp)-1]
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}
}

//No.48 matrix similar to 54,but exchange one by one
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}

}

//No.49 sort string hash
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, len(strs))
	length := 0
	hash := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		if value, ok := hash[string(b)]; ok {
			res[value] = append(res[value], strs[i])
		} else {
			res[length] = []string{strs[i]}
			hash[string(b)] = length
			length++
		}
	}
	res = res[:length]
	return res
}

//No.50 math
func myPow(x float64, n int) float64 {
	var res float64
	res = 1
	flag := true
	symbal := true
	if n < 0 {
		flag = false
		n = -n
	}
	if x < 0 && n%2 == 1 {
		symbal = false
	}

	count := math.Abs(x)
	power := 1
	for n > 0 {
		if power <= n {
			res = res * count
			n -= power
			power += power
			count *= count
		} else {
			count = math.Sqrt(count)
			power /= 2

		}

	}
	if !flag {
		res = 1 / res
	}
	if !symbal {
		res = -res
	}
	return res
}

//NO.51 not in same row,column,diagonal
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	tmp := make([]string, 0)
	hash := make(map[int]bool, n)
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		str[i] = '.'
	}
	backTrace(&res, hash, tmp, 0, n, str)

	return res
}
func backTrace(res *[][]string, hash map[int]bool, tmp []string, layer int, n int, str []byte) {
	if layer == n {
		tmp1 := make([]string, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
	}
	for j := 0; j < n; j++ {
		if !hash[j] {
			//check if attack in diagonal line
			flag := false
			for l, r := len(tmp)-1, j+1; l >= 0 && r < n; l, r = l-1, r+1 {
				if tmp[l][r] == 'Q' {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			for l, r := len(tmp)-1, j-1; l >= 0 && r >= 0; l, r = l-1, r-1 {
				if tmp[l][r] == 'Q' {
					flag = true
				}
			}
			if flag {
				continue
			}

			str[j] = 'Q'
			tmp = append(tmp, string(str))
			hash[j] = true
			str[j] = '.'
			backTrace(res, hash, tmp, layer+1, n, str)
			tmp = tmp[:len(tmp)-1]
			hash[j] = false
		}
	}
}

//NO.52
func totalNQueens(n int) int {
	res := 0
	tmp := make([]string, 0)
	hash := make(map[int]bool, n)
	str := make([]byte, n)
	for i := 0; i < n; i++ {
		str[i] = '.'
	}
	backTrace2(&res, hash, tmp, 0, n, str)

	return res
}
func backTrace2(res *int, hash map[int]bool, tmp []string, layer int, n int, str []byte) {
	if layer == n {
		*res = *res + 1
		return
	}
	for j := 0; j < n; j++ {
		if !hash[j] {
			//check if attack in diagonal line
			flag := false
			for l, r := len(tmp)-1, j+1; l >= 0 && r < n; l, r = l-1, r+1 {
				if tmp[l][r] == 'Q' {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			for l, r := len(tmp)-1, j-1; l >= 0 && r >= 0; l, r = l-1, r-1 {
				if tmp[l][r] == 'Q' {
					flag = true
				}
			}
			if flag {
				continue
			}

			str[j] = 'Q'
			tmp = append(tmp, string(str))
			hash[j] = true
			str[j] = '.'
			backTrace2(res, hash, tmp, layer+1, n, str)
			tmp = tmp[:len(tmp)-1]
			hash[j] = false
		}
	}
}

//No.53 s[j]represent max end with a[j]
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}

//No.54 matrix bound condition
func spiralOrder(matrix [][]int) []int {
	row, column := 0, 0
	res := make([]int, len(matrix)*len(matrix[0]))
	rLen := 0
	end := 101
	for true {
		num4 := true
		for ; column < len(matrix[0]) && matrix[row][column] != end; column++ {
			res[rLen] = matrix[row][column]
			matrix[row][column] = end
			rLen++
		}
		column--
		row++
		for ; row < len(matrix) && matrix[row][column] != end; row++ {
			res[rLen] = matrix[row][column]
			matrix[row][column] = end
			rLen++
		}
		row--
		column--

		for ; column >= 0 && matrix[row][column] != end; column-- {
			res[rLen] = matrix[row][column]
			matrix[row][column] = end
			rLen++
		}
		column++
		row--

		for ; row >= 0 && matrix[row][column] != end; row-- {
			num4 = false
			res[rLen] = matrix[row][column]
			matrix[row][column] = end
			rLen++
		}
		row++

		if num4 {
			break
		}

		column++
	}
	return res
}

//No.55 greedy
func canJump(nums []int) bool {
	maxPo := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > maxPo {
			maxPo = i + nums[i]
		}
		if maxPo >= len(nums)-1 {
			return true
		}
		if i == end {
			end = maxPo
		}
		if i == maxPo {
			return false
		}

	}
	return false
}

//No.56 first sort,and then merge
type interNums [][]int

func (arr interNums) Len() int {
	return len(arr)
}
func (arr interNums) Swap(i, j int) {
	arr[i][0], arr[j][0] = arr[j][0], arr[i][0]
	arr[i][1], arr[j][1] = arr[j][1], arr[i][1]
}
func (arr interNums) Less(i, j int) bool {
	return arr[i][0] < arr[j][0]
}
func merge(intervals [][]int) [][]int {
	sort.Sort(interNums(intervals))
	res := make([][]int, 0)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] {
			if tmp[1] < intervals[i][1] {
				tmp[1] = intervals[i][1]
			}
		} else {
			res = append(res, tmp)
			tmp = intervals[i] //tmp point another arr,not change res value
		}
	}
	res = append(res, []int(tmp))
	return res
}

//No.57
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0)
	index := -1
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] >= newInterval[0] {
			index = i
			break
		} else {
			res = append(res, intervals[i])
		}
	}
	if index == -1 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	tmp := newInterval
	for i := index; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] {
			if tmp[1] < intervals[i][1] {
				tmp[1] = intervals[i][1]
			}
			if tmp[0] > intervals[i][0] {
				tmp[0] = intervals[i][0]
			}
		} else {
			res = append(res, tmp)
			tmp = intervals[i]
		}
	}

	res = append(res, []int(tmp))
	return res
}

//No.58
func lengthOfLastWord(s string) int {
	var left, rig int
	for rig = len(s) - 1; rig >= 0 && s[rig] == ' '; rig-- {
	}
	for left = rig - 1; left >= 0 && s[left] != ' '; left-- {
	}
	return rig - left
}

//N0.59
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	index := 1
	row, column := 0, 0
	for index <= n*n {
		for ; column < n && matrix[row][column] == 0; column++ {
			matrix[row][column] = index
			index++
		}
		column--
		row++
		for ; row < n && matrix[row][column] == 0; row++ {
			matrix[row][column] = index
			index++
		}
		row--
		column--
		for ; column >= 0 && matrix[row][column] == 0; column-- {
			matrix[row][column] = index
			index++
		}
		column++
		row--
		for ; row >= 0 && matrix[row][column] == 0; row-- {
			matrix[row][column] = index
			index++
		}
		row++

		column++

	}
	return matrix
}

//No.60  index=(k-1)/(n-1)! k=k%(n-1)!
func getPermutation(n int, k int) string {
	mark := make([]bool, n+1)
	count := make([]int, n)
	count[0] = 1
	for i := 1; i < n; i++ {
		count[i] = i * count[i-1]
	}
	res := ""
	for i := 1; i < n; i++ {
		if k == 0 {
			break
		}
		index := (k-1)/count[n-i] + 1
		k = k % count[n-i]
		posi := 1
		for j := 1; j <= index; {
			if mark[posi] == true {
				posi++
				continue
			}
			if j == index {
				mark[posi] = true
				res += strconv.Itoa(posi)
				break
			}
			j++
			posi++
		}
	}
	for i := n; i >= 1; i-- {
		if mark[i] == false {
			res += strconv.Itoa(i)
		}
	}

	return res
}
