package leedcode

//最短路径算法，以集合进行遍历
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

//126 记录所有最短路径
//广度搜索加深度搜索
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
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

	str := make([]string, 0)
	laddersBackTracking(queue, wSet, &ret, str)
	return ret
}

func laddersBackTracking(queue []string, set map[string]int, ret *[][]string, str []string) {
	index := len(queue)
	for i := 0; i < index; i++ {
		head := queue[i]

	}
}
