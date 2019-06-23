package leedcode

//127 广度优先遍历最短路径算法，以集合进行遍历
//两种维护的思考，1.以两个单词为端 2.以worldlist构图
//byte uint8,rune int32,string[i]=byte
func ladderLength(beginWord string, endWord string, wordList []string) int {
	ret := 0
	if beginWord == endWord || len(wordList) < 1 || len(beginWord) != len(endWord) {
		return ret
	}
	wSet := make(map[string]int) //单词集合，用来查找，过滤重复
	queue := make([]string, 0)   //辅助队列
	for i := 0; i < len(wordList); i++ {
		wSet[wordList[i]] = 1
	}
	if wSet[endWord] != 1 {
		return ret
	}

	queue = append(queue, beginWord)
	nextLen := 0
	index := 1 //遍历的位置
	wordLen := len(beginWord)
	find := false

	for index > 0 && !find {
		index--
		head := queue[0]
		for i := 0; i < wordLen && !find; i++ {

			var j byte
			for j = 'a'; j <= 'z'; j++ {
				tmp := []byte(head)
				tmp[i] = j
				str := string(tmp)
				if wSet[str] == 1 {
					queue = append(queue, str)
					nextLen++
					delete(wSet, str)
					if str == endWord {
						find = true
						break
					}
				}
			}
		}
		queue = queue[1:]
		if index == 0 && !find {
			index = nextLen
			nextLen = 0
			ret++
		}
	}
	if find {
		return ret + 2
	}
	return 0
}

func ladder(beginWord string, endWord string, wordList []string) (int, map[string][]string) {
	ret := 0
	if beginWord == endWord || len(wordList) < 1 || len(beginWord) != len(endWord) {
		return 0, nil
	}
	wSet := make(map[string]int) //单词集合，用来查找，过滤重复
	queue := make([]string, 0)   //辅助队列
	resHash := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		wSet[wordList[i]] = 1
	}
	if wSet[endWord] != 1 {
		return 0, nil
	}

	queue = append(queue, beginWord)
	nextLen := 0
	index := 1 //遍历的位置
	wordLen := len(beginWord)
	find := false
	del := make(map[string]bool) //记录删除的string

	for index > 0 {
		index--
		head := queue[0]
		for i := 0; i < wordLen; i++ {

			var j byte
			for j = 'a'; j <= 'z'; j++ {
				tmp := []byte(head)
				tmp[i] = j
				str := string(tmp)
				if wSet[str] == 1 || del[str] {
					if !del[str] {
						queue = append(queue, str)
						nextLen++
					}
					delete(wSet, str)
					del[str] = true

					if head != str {
						resHash[head] = append(resHash[head], str)
					}
					if str == endWord {
						find = true
					}
				}
			}
		}
		queue = queue[1:]
		if index == 0 && find {
			break
		}
		if index == 0 && !find {
			del = make(map[string]bool)
			index = nextLen
			nextLen = 0
			ret++
		}
	}
	if find {
		ret += 2
	} else {
		ret = 0
	}
	return ret, resHash
}

/*126 首先利用127得出层级和图的hash表示类似map[ted:[tad tex] tad:[tax] red:[ted rex]]
然后再用深度遍历
127修改时，各种小问题，不能放入重复的key，获取集合不完整等
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
	count, hash := ladder(beginWord, endWord, wordList)
	if count == 0 {
		return ret
	}
	ladderDFS(&ret, hash, beginWord, []string{}, count, endWord)
	return ret
}
func ladderDFS(res *[][]string, hash map[string][]string, index string, str []string, count int, end string) {
	str = append(str, index)
	if len(str) == count {
		if index == end {
			tmp := make([]string, len(str))
			copy(tmp, str)
			*res = append(*res, tmp)
			//fmt.Println(res)
		}
		return
	}
	for i := 0; i < len(hash[index]); i++ {
		ladderDFS(res, hash, hash[index][i], str, count, end)
	}
}

//topic 17 回溯算法
func letterCombinations(digits string) []string {
	letter := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	enum(&res, letter, []byte{}, digits, 0)
	return res
}

func enum(res *[]string, let map[byte]string, index []byte, digits string, pos int) {
	if len(index) == len(digits) {
		tmp := make([]byte, len(index))
		copy(tmp, index)
		*res = append(*res, string(tmp))
		return
	}
	for i := 0; i < len(let[digits[pos]]); i++ {
		index = append(index, let[digits[pos]][i])
		enum(res, let, index, digits, pos+1)
		index = append(index[:len(index)-1])
	}
}
//pid 22 类似于层次遍历
func generateParenthesis(n int) []string {
	ret:=make([]string,0)
	if n==0 {
		return ret
	}
	ret=append(ret,"()")
	hash:=make(map[string]int)
	hash["()"]=1
	for i:=2;i<=n;i++ {
		nLen:=len(ret)
		for nLen>0 {
			str:=ret[0]
			for j:=0;j<len(str);j++ {
				if str[j]=='('{
					if hash[str[:j+1]+"()"+str[j+1:]]==0 {
						hash[str[:j+1]+"()"+str[j+1:]]=1
						ret=append(ret,str[:j+1]+"()"+str[j+1:])
					}
				}
			}
			ret=ret[1:]
			if hash["()"+str]==0 {
				hash["()"+str]=1
				ret=append(ret,"()"+str)
			}
			nLen--
		}
	}
	return ret
}
