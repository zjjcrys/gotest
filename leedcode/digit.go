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
