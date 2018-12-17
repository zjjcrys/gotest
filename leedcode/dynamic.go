package leedcode

import "fmt"

// 构造数据结构，
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	arr := new(NumArray)
	length := len(nums)
	sums := make([]int, length)
	for i := 0; i < length; i++ {
		if i == 0 {
			sums[0] = nums[0]
			continue
		}
		sums[i] = sums[i-1] + nums[i]
	}
	arr.sums = sums
	return *arr
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > j {
		return 0
	}
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}

//413 等差数列划分，一个等差数组的子数组个数是有公示的
//找出所有等差数组是容易的
//放数组时，特殊情况忘了考虑,个数统计
func NumberOfArithmeticSlices(A []int) int {
	length := len(A)
	if length < 3 {
		return 0
	}
	tmpCount := 2
	tmpRange := A[1] - A[0]
	ret := 0
	for i := 2; i < length; i++ {
		if A[i]-A[i-1] == tmpRange {
			tmpCount++
		} else {
			if tmpCount >= 3 {
				ret += (tmpCount - 1) * (tmpCount - 2) / 2
			}
			tmpCount = 2 //个数统计
			tmpRange = A[i] - A[i-1]
		}
	}
	if tmpCount >= 3 {
		ret += (tmpCount - 1) * (tmpCount - 2) / 2
	}
	return ret
}

//343 动态规划依赖于之前的结果，是穷举的一种
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	max := 0
	for i := 2; i <= n; i++ {
		if i < n {
			max = i
		}
		for j := 1; j < i; j++ {
			if dp[j]*(i-j) > max {
				max = dp[j] * (i - j)
			}
		}
		dp[i] = max
	}
	return dp[n]
}

//357 用dp方法找规律，dp[i]与dp[i-1]
func CountNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 10*(dp[i-1]-dp[i-2]) - (dp[i-1]-dp[i-2])*(i-1)
	}
	return dp[n]
}

//30 用数组当做栈,go的字符类型,判断的时机不一样,数组的截取
//位置记录代表着全程的扩充过程，通过位置可以实际计算扩充过程
func LongestValidParentheses(s string) int {
	length := len(s)
	if length < 1 {
		return 0
	}
	bracket := make([]int, 0) //记录括号的位置
	ret := 0                  //记录整个流程中的最大值
	index := 0                //记录当前括号
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(':
			bracket = append(bracket, i)
			break
		case ')':
			tmp := len(bracket)
			if tmp == 0 {
				index = i + 1
			} else {
				top := bracket[tmp-1]
				tmp--
				if tmp == 0 {
					ret = max(ret, i-index+1)
				} else {
					ret = max(ret, i-top+1)
				}
				bracket = bracket[:tmp]
				fmt.Println(bracket)
			}
			break
		}
	}
	return ret
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
