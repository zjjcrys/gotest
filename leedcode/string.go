package leedcode

import "sort"

//68
//首先确认字符串的行数，然后确认字符串的位置,梳理好逻辑直接做
func fullJustify(words []string, maxWidth int) []string {
	rowsRecord := make(map[int]int)  //记录每一行放几个数
	wordsRecord := make(map[int]int) //记录每一行放的总单词数，用来排序
	index := 0                       //每一行的字数和
	rows := 0                        //行数
	for i := 0; i < len(words); i++ {
		index += len(words[i])
		if index > maxWidth {
			rows++
			index = len(words[i])
		}
		rowsRecord[rows] += 1
		wordsRecord[rows] += len(words[i])
		index++
	}
	rows = 0
	index = 0
	ret := make([]string, len(rowsRecord))
	for i := 0; i < len(words); i++ {
		if rowsRecord[rows] == 1 {
			ret[rows] = words[i]
			for j := len(words[i]); j < maxWidth; j++ {
				ret[rows] += " "
			}
			rows++
			index = 0
			continue
		}
		if rows == len(rowsRecord)-1 {
			ret[rows] += words[i]
			if i < len(words)-1 {
				ret[rows] += " "
			} else {
				for j := len(ret[rows]); j < maxWidth; j++ {
					ret[rows] += " "
				}
			}
		}
		if index < rowsRecord[rows] && rows < len(rowsRecord)-1 {
			index++
			space := (maxWidth - wordsRecord[rows]) / (rowsRecord[rows] - 1)
			remind := (maxWidth - wordsRecord[rows]) % (rowsRecord[rows] - 1)
			ret[rows] += words[i]

			if index == rowsRecord[rows] {
				rows++
				index = 0
				continue
			}
			for j := 0; j < space; j++ {
				ret[rows] += " "
			}
			if remind > 0 && index <= remind {
				ret[rows] += " "
			}
		}
	}
	return ret
}

//76 我的问题是在一个循环中即处理left，又处理right，情况太多，只处理right，
//两个指针，滑动窗口，
//时间复杂度几乎是排在了最末，如何优化
func minWindow(s string, t string) string {
	ret := ""
	left := 0
	right := 0
	str := make(map[byte]int) // 快速取出，用hash存储t字符
	tmp := make(map[byte]int) //滑动过程中记录
	desLen := len(t)          //t的长度，判断过程中用到
	if desLen < 1 {
		return ""
	}
	for i := 0; i < desLen; i++ {
		str[t[i]]++
	}
	for ; right < len(s); right++ {
		if str[s[right]] == 0 {
			continue
		}

		tmp[s[right]]++
		if equalMap(str, tmp) {
			if ret == "" || right-left+1 < len(ret) {
				ret = s[left : right+1]
			}
			if tmp[s[left]] > 1 {
				tmp[s[left]]--
			} else {
				delete(tmp, s[left])
			}
			left++
			//更新left指针
			for ; left <= right; left++ {
				if str[s[left]] == 0 {
					continue
				}
				if !equalMap(str, tmp) {
					break
				}
				if equalMap(str, tmp) {
					if right-left+1 < len(ret) {
						ret = s[left : right+1]
					}
					if tmp[s[left]] > 1 {
						tmp[s[left]] -= 1
					} else {
						delete(tmp, s[left])
					}
				}

			}

		}

	}

	return ret
}
func equalMap(map1 map[byte]int, map2 map[byte]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] < v {
			return false
		}
	}
	return true
}

//87 递归求解，任意分解，两个子节点交换位置，
//空时返回true，可以交换多次，感觉别人写的代码比我的要顺畅
type byteSlice []byte

func (s byteSlice) Len() int           { return len(s) }
func (s byteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byteSlice) Less(i, j int) bool { return s[i] < s[j] }

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	//解决排序时间超时
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Sort(byteSlice(b1))
	sort.Sort(byteSlice(b2))
	if string(b1) != string(b2) {
		return false
	}
	length := len(s1)
	for index := 1; index < length; index++ {
		s11 := s1[0:index]
		s12 := s1[index:]

		s21 := s2[0:index]
		s22 := s2[index:]

		if isScramble(s11, s21) && isScramble(s12, s22) {
			return true
		}
		s21 = s2[:length-index]
		s22 = s2[length-index:]
		if isScramble(s11, s22) && isScramble(s12, s21) {
			return true
		}
	}
	return false
}
