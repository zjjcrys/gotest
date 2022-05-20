package leedcode

import (
	"sort"
	"strconv"
)

//No.61 link bound:len(head)=0
func rotateRight(head *ListNode, k int) *ListNode {
	lLen := 0
	var last *ListNode
	present := head
	for present != nil {
		lLen++
		present = present.Next
	}
	if lLen == 0 {
		return last
	}

	k = k % lLen
	if k == 0 {
		return head
	}
	n := lLen - k
	present = head
	index := 1
	for index < n {
		present = present.Next
		index++
	}
	res := present.Next
	present.Next = nil

	last = res
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head

	return res
}

//No.62 dp{i,j}=dp{i-1,j}+dp{i,j-1}
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//No.63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	isObstacle := false
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	for j := 0; j < n; j++ {
		if !isObstacle {
			if obstacleGrid[0][j] == 1 {
				obstacleGrid[0][j] = -1
				isObstacle = true
			} else {
				obstacleGrid[0][j] = 1
			}
		} else {
			obstacleGrid[0][j] = -1
		}
	}

	if obstacleGrid[0][0] != -1 {
		isObstacle = false
	}

	for i := 1; i < m; i++ {
		if !isObstacle {
			if obstacleGrid[i][0] == 1 {
				obstacleGrid[i][0] = -1
				isObstacle = true
			} else {
				obstacleGrid[i][0] = 1
			}
		} else {
			obstacleGrid[i][0] = -1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
				continue
			}
			if obstacleGrid[i-1][j] != -1 && obstacleGrid[i][j-1] != -1 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else if obstacleGrid[i-1][j] == -1 && obstacleGrid[i][j-1] != -1 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
			} else if obstacleGrid[i-1][j] != -1 && obstacleGrid[i][j-1] == -1 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
			} else {
				obstacleGrid[i][j] = -1
			}
		}
	}

	if obstacleGrid[m-1][n-1] == -1 {
		return 0
	}
	return obstacleGrid[m-1][n-1]
}

//No.64 similar to pre two dp[i,j]=min{dp[i-1,j],dp[i,j-1]}+a[i,j]
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for j := 1; j < n; j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := grid[i-1][j]
			if min > grid[i][j-1] {
				min = grid[i][j-1]
			}
			grid[i][j] += min
		}
	}
	return grid[m-1][n-1]
}

//No.65 it is not a good question,special case e9,0e,.e1=false 46.e3=true
func isNumber(s string) bool {
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	hash := make(map[byte]int)
	if s[0] == '.' {
		hash[s[0]] = 0
	}

	res := true
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'E', 'e':
			if i == 0 || i == len(s)-1 {
				res = false
				break
			}
			if _, ok := hash['e']; ok {
				res = false
				break
			} else {
				if value, ok := hash['.']; ok && value == i-1 && i == 1 {
					res = false
					break
				}
				hash['e'] = i
			}
		case '+', '-':
			value, ok := hash['e']
			if ok && value == i-1 {
				if i == len(s)-1 {
					res = false
					break
				}
			} else {
				res = false
				break
			}
		case '.':
			if value, ok := hash[s[i]]; ok {
				if value != i {
					res = false
					break
				}
			} else {
				if _, ok := hash['e']; ok {
					res = false
					break
				}
				hash['.'] = i
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			res = false
		}

		if !res {
			break
		}

	}

	//only .||e false
	if _, ok := hash['.']; ok && len(s) == 1 {
		res = false
	}
	if _, ok := hash['e']; ok && len(s) == 1 {
		res = false
	}

	return res
}

//66 simulation
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 0 {
			break
		}
		tmp := carry + digits[i]
		digits[i] = tmp % 10
		carry = tmp / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

//67 simulation
func addBinary(a string, b string) string {
	res := make([]byte, len(a)+len(b))
	carry := byte(0)

	length := 0
	var i, j int
	for i, j = len(a)-1, len(b)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		tmp := a[i] - '0' + b[j] - '0' + carry
		res[length] = tmp % 2
		length++
		carry = tmp / 2
	}
	for i >= 0 {
		tmp := a[i] - '0' + carry
		res[length] = tmp % 2
		length++
		carry = tmp / 2
		i--
	}
	for j >= 0 {
		tmp := b[j] - '0' + carry
		res[length] = tmp % 2
		length++
		carry = tmp / 2
		j--
	}
	if carry > 0 {
		res[length] = carry
		length++
	}
	res = res[:length]
	for i, j = 0, length-1; i < j; i, j = i+1, j-1 {
		res[i] += '0'
		res[j] += '0'
		res[i], res[j] = res[j], res[i]
	}
	if i == j {
		res[i] += '0'
	}

	return string(res)
}

//68
func fullJustify(words []string, maxWidth int) []string {
	left, rig := 0, 0
	space, extra, count := 0, 0, -1
	res := make([]string, 0)
	for rig < len(words) {
		if count+1+len(words[rig]) <= maxWidth {
			count += 1 + len(words[rig])
			rig++
			continue
		}

		wordsCount := maxWidth - count + (rig - left - 1)
		tmp := ""
		if rig-left-1 <= 0 {
			space = wordsCount

			tmp += words[left]
			for i := 0; i < space; i++ {
				tmp += " "
			}
			res = append(res, tmp)
		} else {
			space = wordsCount / (rig - left - 1)
			extra = wordsCount % (rig - left - 1)

			for left < rig-1 {
				tmp += words[left]
				for i := 0; i < space; i++ {
					tmp += " "
				}
				if extra > 0 {
					tmp += " "
					extra--
				}

				left++
			}
			tmp += words[left]
			res = append(res, tmp)
		}

		left = rig
		count = -1
	}
	//deal with the last line
	tmp := ""
	for left < rig {
		tmp += words[left] + " "
		left++
	}
	if len(tmp) > maxWidth {
		tmp = tmp[:len(tmp)-1]
	} else {
		for i := len(tmp) + 1; i <= maxWidth; i++ {
			tmp += " "
		}
	}
	res = append(res, tmp)

	return res
}

//69 binary simulation
func mySqrt(x int) int {
	left, rig := 0, x/2
	for left <= rig {
		mid := (left + rig) / 2
		if mid*mid == x {
			left = mid
			break
		} else if mid*mid < x {
			left = mid + 1
		} else {
			rig = mid - 1
		}
	}

	if left*left > x {
		left--
	}
	return left
}

//70 dynamic
func climbStairs(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	if n == 1 {
		return arr[n-1]
	}
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}

//71 case by case discussion ..hidden
func simplifyPath(path string) string {
	arr := make([]string, len(path))
	length := 0
	arr[length] = string(path[0])
	length++
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			if arr[length-1] == "/" {
				continue
			} else {
				arr[length] = "/"
				length++
			}
		} else if path[i] == '.' {
			tmp := "."
			for left := i + 1; left < len(path) && path[left] != '/'; left++ {
				tmp += string(path[left])
			}

			if len(tmp) == 2 {
				if tmp == ".." {
					if length > 1 {
						length -= 2
						if length == 0 {
							length = 1
						}
					}
				} else {
					arr[length] = tmp
					length++
				}
			} else if len(tmp) > 2 {
				arr[length] = tmp
				length++
			}
			i += len(tmp) - 1
		} else {
			tmp := string(path[i])
			for left := i + 1; left < len(path) && path[left] != '/'; left++ {
				tmp += string(path[left])
			}
			arr[length] = tmp
			length++

			i += len(tmp) - 1
		}

	}
	//delete the last / if / is the last word
	if length > 1 && arr[length-1] == "/" {
		length--
	}
	res := ""
	for i := 0; i < length; i++ {
		res += arr[i]
	}
	return res
}

//72 hard dynamic:subQuestion dp[i,j]=1+min{do[i-1,j-1,dp[i,j-1],dp[i,j-1]}
//easier than 87 after know the relation
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for j := 0; j < n; j++ {
		dp[0][j+1] = dp[0][j] + 1
	}
	for i := 0; i < m; i++ {
		dp[i+1][0] = dp[i][0] + 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}
			min := dp[i][j]
			if min > dp[i+1][j] {
				min = dp[i+1][j]
			}
			if min > dp[i][j+1] {
				min = dp[i][j+1]
			}
			dp[i+1][j+1] = min + 1
		}
	}
	return dp[m][n]
}

//73 if a[i][j]==0,specify a[i][0]=0and a[0][j]=0 thus use constant space
//notice a[0][0]
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	zeroRow := false
	zeroColumn := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					zeroRow = true
				}
				if j == 0 {
					zeroColumn = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if zeroRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if zeroColumn {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

//74 bound condition:row<0
func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	m, n := len(matrix), len(matrix[0])
	for ; row < m && matrix[row][0] <= target; row++ {
	}
	row--
	if row < 0 {
		return false
	}

	left, rig := 0, n-1
	for left <= rig {
		mid := (left + rig) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			rig = mid - 1
		}
	}
	return false
}

//75 insert sort
func sortColors(nums []int) {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > tmp {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = tmp
	}
}

//No.76 different to 30 sliding window +hash
func minWindow(s string, t string) string {
	hashCount := make(map[byte]int)
	hashProcess := make(map[byte]int)
	count := 0
	res := ""
	left, rig := 0, 0 //the sliding window
	for i := 0; i < len(t); i++ {
		hashCount[t[i]] += 1
	}
	for rig < len(s) {
		if count < len(t) { //rig++
			if hashCount[s[rig]] > 0 {
				hashProcess[s[rig]] += 1
				if hashProcess[s[rig]] <= hashCount[s[rig]] {
					count++
				}
			}
			rig++
		} else { //left++ if count == len(t)
			if len(res) == 0 || len(res) > rig-left {
				res = s[left:rig]
			}
			if hashCount[s[left]] > 0 {
				if hashProcess[s[left]] <= hashCount[s[left]] {
					count--
				}
				hashProcess[s[left]] -= 1
			}
			left++

		}
	}
	for count == len(t) {
		if len(res) == 0 || len(res) > rig-left {
			res = s[left:rig]
		}
		if hashCount[s[left]] > 0 {
			if hashProcess[s[left]] <= hashCount[s[left]] {
				count--
			}
			hashProcess[s[left]] -= 1
		}
		left++
	}

	return res
}

//77 backtrace similar to 39,40,46,47,51,52
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	backTrace77(&res, []int{}, n, k, 1)
	return res
}
func backTrace77(res *[][]int, tmp []int, n int, k int, index int) {
	if len(tmp) == k {
		tmp1 := make([]int, 0)
		tmp1 = append(tmp1, tmp...)
		*res = append(*res, tmp1)
		return
	}
	for i := index; i <= n; i++ {
		tmp = append(tmp, i)
		backTrace77(res, tmp, n, k, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}

//78 backtrace similar to 39,40,46,47,51,52,77
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backTrace78(&res, nums, 0, []int{})
	return res
}
func backTrace78(res *[][]int, nums []int, index int, tmp []int) {
	tmp1 := make([]int, 0)
	tmp1 = append(tmp1, tmp...)
	*res = append(*res, tmp1)
	if len(tmp) == len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backTrace78(res, nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}

//79 backtrace 39,40,46,47,51,52,77,78
func exist(board [][]byte, word string) bool {
	res := false
	m := len(board)
	n := len(board[0])
	hash := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				backTrace79(&res, board, word, i, j, 0, hash)
			}
			if res {
				break
			}
		}

		if res {
			break
		}
	}
	return res
}
func backTrace79(res *bool, board [][]byte, word string, row int, column int, index int, hash map[int]bool) {
	if *res == true {
		return
	}
	if !hash[row*10+column] && board[row][column] == word[index] {
		hash[row*10+column] = true
		if index == len(word)-1 {
			*res = true
			return
		}

		// up and down
		if row-1 >= 0 {
			backTrace79(res, board, word, row-1, column, index+1, hash)
		}
		if row+1 < len(board) {
			backTrace79(res, board, word, row+1, column, index+1, hash)
		}
		//left and rig
		if column-1 >= 0 {
			backTrace79(res, board, word, row, column-1, index+1, hash)
		}
		if column+1 < len(board[0]) {
			backTrace79(res, board, word, row, column+1, index+1, hash)
		}
		hash[row*10+column] = false
	}
}

//80 two pointers
func removeDuplicates(nums []int) int {
	length := 0
	left, rig := 0, 0
	for rig < len(nums) {
		for rig < len(nums) && nums[rig] == nums[left] {
			rig++
		}
		rig--
		nums[length] = nums[left]
		length++
		if rig != left {
			nums[length] = nums[rig]
			length++
		}
		left = rig + 1
		rig = left
	}
	return length
}

//81 similar to 33,for nums[mid]=nums[left]=nums[rig] only move iteratively
func search(nums []int, target int) bool {
	left, rig := 0, len(nums)-1
	for left <= rig {
		mid := (left + rig) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			if nums[mid] > nums[left] {
				left = mid + 1
			} else if nums[rig] > nums[mid] && nums[rig] < target {
				if nums[rig] < target {
					rig = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				if nums[left] == target {
					return true
				} else {
					left++
				}
			}

		} else {
			if nums[mid] > nums[left] {
				if target < nums[left] {
					left = mid + 1
				} else {
					rig = mid - 1
				}
			} else if nums[mid] < nums[rig] {
				rig = mid - 1
			} else {
				if nums[left] == target {
					return true
				} else {
					left++
				}
			}
		}
	}
	return false
}

//82 two pointers in linklist
func deleteDuplicates(head *ListNode) *ListNode {
	res := new(ListNode)
	res.Next = head
	head = res
	left := res.Next
	rig := left
	for rig != nil {
		rig = left
		for rig != nil && rig.Val == left.Val {
			rig = rig.Next
		}
		if left.Next == rig { //distinct value
			head = left
			left = rig
		} else {
			head.Next = rig
			left = rig
		}
	}
	return res.Next
}

//83 two pointers too simple
//84 monotonic stack,confirm width ,can use centry mini the compare time but not time complexity
func largestRectangleArea(heights []int) int {
	res := 0
	stack := make([]int, len(heights))
	length := 0
	for i := 0; i < len(heights); i++ {
		for length > 0 && heights[stack[length-1]] > heights[i] {
			width := i
			if length > 1 {
				width = i - stack[length-2] - 1
			}
			tmp := heights[stack[length-1]] * width
			if tmp > res {
				res = tmp
			}
			length--
		}
		stack[length] = i
		length++
	}
	for length > 1 {
		tmp := heights[stack[length-1]] * (len(heights) - stack[length-2] - 1)
		if tmp > res {
			res = tmp
		}
		length--
	}
	//the last one
	if heights[stack[0]]*len(heights) > res {
		res = heights[stack[0]] * len(heights)
	}

	return res
}

//85 translate to 84,think every layer as height
func maximalRectangle(matrix [][]byte) int {
	res := 0
	for j := 0; j < len(matrix[0]); j++ {
		matrix[0][j] -= '0'
	}
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = matrix[i-1][j] + 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		max := largestLayerArea(matrix[i])
		if max > res {
			res = max
		}
	}
	return res
}
func largestLayerArea(nums []byte) int {
	res := 0
	stack := make([]int, len(nums))
	length := 0
	for i := 0; i < len(nums); i++ {
		for length > 1 && nums[stack[length-1]] > nums[i] {
			tmp := int(nums[stack[length-1]]) * (i - int(stack[length-2]) - 1)
			if tmp > res {
				res = tmp
			}
			length--
		}
		if length == 1 && nums[stack[length-1]] > nums[i] {
			tmp := int(nums[stack[length-1]]) * i
			if tmp > res {
				res = tmp
			}
			length--
		}
		stack[length] = i
		length++
	}
	for length > 1 {
		tmp := int(nums[stack[length-1]]) * (len(nums) - int(stack[length-2]) - 1)
		if tmp > res {
			res = tmp
		}
		length--
	}
	if length == 1 {
		tmp := int(nums[stack[length-1]]) * len(nums)
		if tmp > res {
			res = tmp
		}
		length--
	}
	return res
}

//86 partition in linklist,use two pointers
func partition(head *ListNode, x int) *ListNode {
	res := new(ListNode)
	res.Next = head
	left, rig := res, head
	for rig != nil {
		if rig.Val >= x {
			break
		}
		rig = rig.Next
	}
	if rig == nil {
		return res.Next
	}
	for rig != nil {
		for left.Next != nil && left.Next.Val < x {
			left = left.Next
		}

		for rig.Next != nil && rig.Next.Val >= x {
			rig = rig.Next
		}

		if rig.Next == nil {
			break
		}
		tmp := rig.Next
		rig.Next = tmp.Next
		tmp.Next = left.Next
		left.Next = tmp
	}
	return res.Next
}

//87 dynamic f{i,j}=f{i,k}+f{k,j},the place from s2turn to s1 must be maintain.f{i,j,place}
func isScramble(s1 string, s2 string) bool {
	hashCount := make(map[byte]int) //word count
	n := len(s1)
	for i := 0; i < n; i++ {
		hashCount[s1[i]]++
	}
	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n)
		}
	}
	//initial
	for i := 0; i < n; i++ {
		if hashCount[s2[i]] == 0 {
			return false
		}
		hashCount[s2[i]]--
		for j := 0; j < n; j++ {
			if s2[i] == s1[j] {
				dp[i][i][j] = true //solve s2的字符映射到多个s1上，不用hash
			}
		}
	}

	for l := 1; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l
			length := j + 1 - i
			for p := 0; p < n-length+1; p++ {
				for k := i + 1; k <= j; k++ {
					if dp[i][k-1][p] && dp[k][j][p+k-i] {
						dp[i][j][p] = true
						break
					}
					if dp[i][k-1][p+j-k+1] && dp[k][j][p] {
						dp[i][j][p] = true
						break
					}

				}
			}
		}
	}
	return dp[0][n-1][0]
}

//88merge in-place, from end iteratively
func merge(nums1 []int, m int, nums2 []int, n int) {
	length := len(nums1)
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] < nums2[n] {
			nums1[length-1] = nums2[n]
			n--
		} else {
			nums1[length-1] = nums1[m]
			m--
		}
		length--
	}
	for n >= 0 {
		nums1[length-1] = nums2[n]
		n--
		length--
	}
}

//89 dynamic symmetry
func grayCode(n int) []int {
	res := make([]int, 1<<uint(n))
	res[0] = 0
	res[1] = 1
	length := 2
	for i := 2; i <= n; i++ {
		tmp := 1 << uint(i-1)
		end := 1 << uint(i)
		for j := tmp; j < end; j++ {
			res[length] = res[end-1-j] + tmp
			length++
		}
	}
	return res
}

//90  39,40,46,47,51,52,77,78,79
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	backTrace90(&res, nums, 0, []int{})
	return res
}
func backTrace90(res *[][]int, nums []int, index int, tmp []int) {
	tmp1 := make([]int, 0)
	tmp1 = append(tmp1, tmp...)
	*res = append(*res, tmp1)
	if index >= len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		backTrace90(res, nums, i+1, tmp)
		tmp = tmp[:len(tmp)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
}

//91 dynamic similar to 70
func numDecodings(s string) int {
	arr := make([]int, len(s)+1)
	arr[0] = 1
	if s[0] >= '1' && s[0] <= '9' {
		arr[1] = 1
	}
	if len(s) < 2 {
		return arr[1]
	}
	flag := false
	for i := 1; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			arr[i+1] += arr[i]
		}
		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			arr[i+1] += arr[i-1]
		}
		if arr[i+1] == 0 {
			flag = true
			break
		}
	}
	if flag {
		return 0
	}

	return arr[len(s)]
}

//92 reverse list
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	res := new(ListNode)
	res.Next = head

	var leftNode, node1, node2, before *ListNode
	index := 0
	leftNode = head
	before = res
	for leftNode != nil {
		index++
		if index == left-1 {
			before = leftNode
		}
		if index == left {
			break
		}
		leftNode = leftNode.Next
	}
	node1 = leftNode

	for index <= right {
		tmp := node1.Next
		node1.Next = node2
		node2 = node1
		node1 = tmp
		index++
	}
	before.Next = node2
	leftNode.Next = node1

	return res.Next
}

//93 backtrace 39,40,46,47,51,52,77,78,79,90
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	backTrace93(&res, s, "", 0, 0)
	return res
}
func backTrace93(res *[]string, s string, tmp string, index int, dot int) {
	if dot > 4 {
		return
	}
	if index >= len(s) {
		if dot < 4 {
			return
		}
		tmp1 := tmp[:len(tmp)-1]
		*res = append(*res, tmp1)
		return
	}
	for i := index + 1; i <= index+3 && i <= len(s); i++ {
		if s[index] == '0' && i > index+1 {
			continue
		}
		num, _ := strconv.Atoi(s[index:i])
		count := i - index
		if num > 255 {
			break
		}
		tmp += s[index:i]
		tmp += "."
		dot++
		count++

		backTrace93(res, s, tmp, i, dot)
		tmp = tmp[:len(tmp)-count]
		dot--

	}
}
