package leedcode

import "strconv"

//130 先把边界以及和边界相连的剔除，
func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1) && board[i][j] == 'O' {
				solveDFS(&board, i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '-' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(board *[][]byte, row int, col int) {
	(*board)[row][col] = '-'
	if row-1 >= 0 && (*board)[row-1][col] == 'O' {
		solveDFS(board, row-1, col)
	}
	if row+1 < len(*board) && (*board)[row+1][col] == 'O' {
		solveDFS(board, row+1, col)
	}
	if col-1 >= 0 && (*board)[row][col-1] == 'O' {
		solveDFS(board, row, col-1)
	}
	if col+1 < len((*board)[0]) && (*board)[row][col+1] == 'O' {
		solveDFS(board, row, col+1)
	}
}

//200 dfs+广度遍历
func numIslands(grid [][]byte) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				landDFS(&grid, i, j)
				ret++
			}
		}
	}
	return ret
}

func landDFS(board *[][]byte, row int, col int) {
	(*board)[row][col] = '-'
	if row-1 >= 0 && (*board)[row-1][col] == '1' {
		landDFS(board, row-1, col)
	}
	if row+1 < len(*board) && (*board)[row+1][col] == '1' {
		landDFS(board, row+1, col)
	}
	if col-1 >= 0 && (*board)[row][col-1] == '1' {
		landDFS(board, row, col-1)
	}
	if col+1 < len((*board)[0]) && (*board)[row][col+1] == '1' {
		landDFS(board, row, col+1)
	}
}

//264 丑数2，小值推出大值，
func nthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}
	ret := make([]int, n+1)
	ret[1] = 1
	m2 := 1 //*2的第几个丑数
	m3 := 1 //*3的第几个丑数
	m5 := 1 //*5的第几个丑数

	for i := 2; i <= n; i++ {
		ret[i] = min(min(ret[m2]*2, ret[m3]*3), ret[m5]*5)

		if ret[i] == ret[m2]*2 {
			m2++
		}
		if ret[i] == ret[m3]*3 {
			m3++
		}
		if ret[i] == ret[m5]*5 {
			m5++
		}
	}
	return ret[n]
}

//topic 46 回溯也就是深度优先遍历
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 1 {
		return ret
	}
	backtrack(&ret, nums, 0)
	return ret
}

func backtrack(ret *[][]int, tmp []int, start int) {
	if start >= len(tmp) {
		item := make([]int, len(tmp))
		copy(item, tmp)
		*ret = append(*ret, item)
		return
	}
	for i := start; i < len(tmp); i++ {
		swap(&tmp[start], &tmp[i])
		backtrack(ret, tmp, start+1)
		swap(&tmp[i], &tmp[start])
	}
}

func swap(num1 *int, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

// topic 257
func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	if root == nil {
		return ret
	}

	dfsNode(root, &ret, "")
	return ret
}

func dfsNode(root *TreeNode, strArr *[]string, str string) {
	if root != nil {
		pathSB := str
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*strArr = append(*strArr, pathSB)
		} else {
			pathSB += "->"
			dfsNode(root.Left, strArr, pathSB)
			dfsNode(root.Right, strArr, pathSB)
		}
	}
}
