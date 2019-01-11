package leedcode

import (
	"strings"
)

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
//动态规划应该怎么做？
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
				tmp--
				if tmp == 0 {
					ret = max(ret, i-index+1)
				} else {
					ret = max(ret, i-bracket[tmp-1])
				}
				bracket = bracket[:tmp]
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
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

//188-需要再看一遍
//但这道题还有个坑，就是如果k的值远大于prices的天数，就是用2的方法
//https://www.cnblogs.com/ariel-dreamland/p/9166176.html
func maxProfit(k int, prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	if k >= len(prices) {
		return multiDeal(prices)
	}
	glo := make([]int, k+1)
	local := make([]int, k+1)
	for i := 0; i < len(prices)-1; i++ {
		diff := prices[i+1] - prices[i]
		for j := k; j >= 1; j-- {
			local[j] = max(glo[j-1], local[j]) + diff
			glo[j] = max(local[j], glo[j])
		}
	}
	return glo[k]

}

//121 最多一笔交易 就是依赖最大和最小之间的差值
//初始状态0，状态转移max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
func maxProfit1(prices []int) int {
	ret := 0
	if len(prices) < 2 {
		return ret
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > ret {
			ret = prices[i] - min
		}
	}
	return ret
}

//122 可以多笔交易
//只要出现差集就卖出买入，这个是贪心算法更贴切
func multiDeal(prices []int) int {
	ret := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ret += prices[i] - prices[i-1]
		}
	}
	return ret
}

//123 最多两笔交易 如果是连续增加，就要在最高点卖，选择两次的差集
//把数组分为两个子数组，分别求最大值
//174 dp：两个条件可以确认下一个状态,从上到下，需要考虑两个记录，从下往上考虑一个记录，好像是一样的
func calculateMinimumHP(dungeon [][]int) int {
	row := len(dungeon)
	if row < 1 {
		return 0
	}
	col := len(dungeon[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			ret := 0
			if i+1 < row && j+1 < col {
				ret = min(dp[i+1][j], dp[i][j+1])
			} else if i+1 < row {
				ret = dp[i+1][j]
			} else if j+1 < col {
				ret = dp[i][j+1]
			}
			dp[i][j] = max(0, ret-dungeon[i][j])
		}
	}
	return dp[0][0] + 1
}

//139 f(i)标识 前面1到i个字符能否被拆分
func isBreak(s string, wordDict []string) ([]bool, map[string]bool) {
	dp := make([]bool, len(s)+1)
	set := make(map[string]bool, len(wordDict))
	minLen := len(wordDict[0])
	maxLen := minLen
	for i := 0; i < len(wordDict); i++ {
		set[wordDict[i]] = true
		if len(wordDict[i]) < minLen {
			minLen = len(wordDict[i])
		}
		if len(wordDict[i]) > maxLen {
			maxLen = len(wordDict[i])
		}

	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ { //这里可以被优化，根据字符串的长度
			if i-j > maxLen && i-j < minLen {
				continue
			}
			left := true
			rig := set[s[j:i+1]]
			if j != 0 {
				left = dp[j]
			}
			if left && rig {
				dp[i+1] = true
				break
			}
		}
	}
	return dp, set
}

//140 先算出字符串是否被拆分，再调用dfs，搅和在一起写，内存使用上升，逻辑不清晰
func WordBreak(s string, wordDict []string) []string {
	if len(s) < 1 || len(wordDict) < 1 {
		return []string{}
	}
	dp, set := isBreak(s, wordDict)
	if !dp[len(s)] {
		return []string{}
	}
	ret := make([]string, 0)
	wordDFS(s, dp, &ret, 0, set, "")
	return ret
}

func wordDFS(s string, dp []bool, res *[]string, index int, set map[string]bool, str string) {
	if index >= len(s) {
		tmp := strings.Trim(str, " ")
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(s); i++ {
		if dp[i+1] && set[s[index:i+1]] {
			str += " " + s[index:i+1]
			wordDFS(s, dp, res, i+1, set, str)
			str = str[0 : len(str)-i-2+index]
		}
	}
}
