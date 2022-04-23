package leedcode

import (
	"fmt"
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

//295 堆的使用小根堆 只用一个堆算法，超出时间限制，利用两个堆求中位数 o(lgn)
//海量数据
type MedianFinder struct {
	arr *[]int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	median := new(MedianFinder)
	median.arr = new([]int)
	return *median
}

func (this *MedianFinder) AddNum(num int) {
	*this.arr = append(*this.arr, num)
	fixUp(this.arr, len(*this.arr)-1)
}

func (this *MedianFinder) FindMedian() float64 {
	var ret float64
	length := len(*this.arr)
	arr := make([]int, length)
	copy(arr, *this.arr)
	arrTmp := heapSort(arr)
	if length%2 == 0 {
		ret = (float64(arrTmp[length/2] + arrTmp[length/2-1])) / 2
	} else {
		ret = float64(arrTmp[length/2])
	}
	return ret
}

//堆的排序跳转
func fixUp(arr *[]int, index int) {
	if index-1 < 0 {
		return
	}
	parent := (index - 1) / 2
	for parent >= 0 {
		if (*arr)[parent] > (*arr)[index] {
			tmp := (*arr)[parent]
			(*arr)[parent] = (*arr)[index]
			(*arr)[index] = tmp
		} else {
			break
		}
		index = parent
		if index-1 < 0 {
			break
		}
		parent = (index - 1) / 2
	}
}

//根据堆生成有序数组 小跟堆
func heapSort(heap []int) []int {
	fmt.Println("before", heap)
	for i := len(heap) - 1; i >= len(heap)/2-1 && i >= 0; i-- {
		tmp := heap[0]
		heap[0] = heap[i]
		heap[i] = tmp
		fmt.Println("exchange", heap)
		//调节 0~i
		for j := 0; j < i; {
			left := false
			rig := false
			if (2*j + 1) >= i { //没有孩子
				break
			}

			if 2*j+2 < i { //有两个孩子
				if heap[2*j+1] < heap[2*j+2] {
					if heap[j] > heap[2*j+1] {
						left = true
					}
				} else {
					if heap[j] > heap[2*j+2] {
						rig = true
					}
				}
			} else if 2*j+1 < i { //只有一个孩子
				if heap[j] > heap[2*j+1] {
					left = true
				}
			}

			if !left && !rig {
				break
			}

			if left {
				tmp := heap[j]
				heap[j] = heap[2*j+1]
				heap[2*j+1] = tmp
				j = 2*j + 1
			} else {
				tmp := heap[j]
				heap[j] = heap[2*j+2]
				heap[2*j+2] = tmp
				j = 2*j + 2
			}
		}
	}
	fmt.Println("after", heap)
	return heap
}

//292
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}

//263
func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num > 1 {
		flag := false
		if num%2 == 0 {
			flag = true
			num = num / 2
		}
		if num%3 == 0 {
			flag = true
			num = num / 3
		}
		if num%5 == 0 {
			flag = true
			num = num / 5
		}
		if !flag {
			return false
		}
	}
	return true
}

//367
func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	left := 2
	right := left * left
	if right == num {
		return true
	}

	for right*right <= num {
		if right*right == num {
			return true
		}

		left = right
		right = left * left
	}

	middle := (left + right) / 2
	for left < middle && middle < right {
		tmp := middle * middle
		if tmp == num {
			return true
		}
		if tmp < num {
			left = middle
			middle = (left + right) / 2
		} else {
			right = middle
			middle = (left + right) / 2
		}
	}
	return false
}

//9 回文数 转化成字符串，然后双指针
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	left := 0
	right := len(str) - 1
	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//6 Z字形变换，直接找规律
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		str[i] = make([]byte, 0)
	}
	sum := 2 * (numRows - 1)
	for i := 0; i < len(s); i++ {
		tmp := (i + 1) % sum
		if tmp == 0 {
			tmp = 2
		} else if tmp > numRows {
			tmp = 2*numRows - tmp
		}
		str[tmp-1] = append(str[tmp-1], s[i])
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret = ret + string(str[i])
	}
	return ret
}

//7整数反转 转化为字符串, append 是在数组结尾新加item
func Reverse(x int) int {
	max := int32(^uint32(0) >> 1)
	min := ^max
	flag := true //表示整数的flag
	if x == 0 {
		return x
	}
	if x < 0 {
		x = -x
		flag = false
	}
	str := strconv.Itoa(x)
	dest := make([]byte, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		dest[len(str)-1-i] = str[i]
	}
	ret, _ := strconv.Atoi(string(dest))
	if flag == false {
		ret = 0 - ret
	}
	fmt.Println(ret, max)
	if ret > int(max) || ret < int(min) {
		return 0
	}
	return ret
}

//topic 306 additive number
func isAdditiveNumber(num string) bool {

}
