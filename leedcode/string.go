package leedcode

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

//76
//两个指针，滑动窗口
func MinWindow(s string, t string) string {
	ret := s
	left := 0
	right := 0
	str := make(map[byte]int) // 快速取出，用hash存储t字符
	tmp := make(map[byte]int) //滑动过程中记录
	desLen := len(t)          //t的长度，判断过程中用到
	if desLen < 1 {
		return ""
	}
	for i := 0; i < desLen; i++ {
		str[t[i]] = 1
	}
	for left < len(s) && right < len(s) && left <= right {
		if str[s[left]] != 1 {
			left++
			continue
		}
		if str[s[right]] != 1 && len(tmp) != desLen {
			right++
			continue
		}
		if str[s[right]] == 1 {
			if tmp[s[right]] > 0 {
				tmp[s[right]] += 1
			} else {
				tmp[s[right]] = 1
			}
		}
		if len(tmp) == desLen {
			tmpRight := right
			if str[s[right]] != 1 {
				tmpRight = right - 1
			}

			if tmpRight-left+1 < len(ret) {
				ret = s[left : tmpRight+1]
				if tmp[s[left]] > 1 {
					tmp[s[left]] -= 1
				} else {
					delete(tmp, s[left])
				}
			}
			left++
		}
		if str[s[right]] == 1 {
			right++
		}
	}
	return ret
}
