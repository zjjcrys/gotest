package leedcode

import (
	"strconv"
)

//166 分数转为小数的计算方法，如果是循环小数，怎么样才算是尽头
//所有的分数都可以化成小数，有限小数和无限循环小数
//0和负数，被除数和除数
func FractionToDecimal(numerator int, denominator int) string {
	symbal := false //默认是整数
	if numerator == 0 {
		return "0"
	}
	if numerator/denominator < 0 || denominator/numerator < 0 {
		symbal = true
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	ret := ""
	index := 0
	hash := make(map[int]int) //记录除数和和余数
	remaind := numerator
	first := true
	var end int

	for remaind != 0 {
		quotient := remaind / denominator
		tmp := remaind % denominator
		ret = ret + strconv.Itoa(quotient)
		index = index + len(strconv.Itoa(quotient))
		if first && tmp != 0 {
			ret = ret + "."
			index++
			first = false
		}
		if hash[tmp] > 0 {
			end = hash[tmp]
			break
		}
		hash[tmp] = index
		remaind = tmp * 10
	}
	if remaind > 0 {
		ret = ret[0:end] + "(" + ret[end:]
		ret += ")"
	}
	if symbal {
		ret = "-" + ret
	}
	return ret
}
