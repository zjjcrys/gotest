package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(twoSum([]int{-18, 12, 3, 0}, -6))
}

//if use two arrays,it is mistake after translate the value to index
//if use map to restore index, take repeat nums into consideration
//the easiest and fastest is use hashmap to restore index while go through the array, accumulate in time
func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	res := []int{0, 0}
	index[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if _, ok := index[target-nums[i]]; ok {
			res[0] = index[target-nums[i]]
			res[1] = i
			break
		}
		index[nums[i]] = i
	}
	return res
}

func getInt(scanner *bufio.Scanner) int {
	num, error := strconv.Atoi(getStr(scanner))
	if error != nil {
		panic(error)
	}
	return num

}

func getStr(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

/*
func fib(param int) int {
	first := 1
	second := 1
	ret := 0
	if param == first || param == 2 {
		return first
	}
	for i := 3; i <= param; i++ {
		ret = first + second
		first = second
		second = ret
	}
	return ret

}

type byLength [][]int

func (arr byLength) Len() int {
	return len(arr)
}
func (arr byLength) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr byLength) Less(i, j int) bool {
	return arr[i][0] < arr[j][0]
}

//topid 56 自定义排序
func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) < 1 {
		return ret
	}
	sort.Sort(byLength(intervals))
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ret[len(ret)-1][1] {
			if intervals[i][1] > ret[len(ret)-1][1] {
				ret[len(ret)-1][1] = intervals[i][1]
			}
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func maximalSquare(matrix [][]byte) int {
	ret := 0
	if len(matrix) < 1 {
		return ret
	}
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j-1], dp[i-1][j])) + 1
				ret = max(ret, dp[i][j])
			}
		}
	}
	return ret * ret
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}
func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}*/
