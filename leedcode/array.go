package leedcode

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

//37 数独，只能回溯算法了
func solveSudoku(board [][]byte) {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	for i := 1; i <= 9; i++ {

	}
}

//判断有效数字
func checkValid() {

}
