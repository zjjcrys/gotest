package leedcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

//99 二叉搜索树
//中序遍历找到错误的值，然后交换，是与原值进行比较
//问题：存储的变量值不是全局的，递归中会丢弃数据,初始赋最小值
func recoverTree(root *TreeNode) {
	notes := make([]*TreeNode, 0)
	left := new(TreeNode)
	left.Val = INT_MIN
	traversal(root, &notes, &left)
	if len(notes) == 2 {
		tmp := notes[0].Val
		notes[0].Val = notes[1].Val
		notes[1].Val = tmp
	}
	if len(notes) == 4 {
		tmp := notes[0].Val
		notes[0].Val = notes[3].Val
		notes[3].Val = tmp
	}
}

//中序遍历，左中右,
func traversal(root *TreeNode, arr *[]*TreeNode, left **TreeNode) {
	if root.Left != nil {
		traversal(root.Left, arr, left)
	}
	if (*left).Val > root.Val {
		*arr = append(*arr, *left)
		*arr = append(*arr, root)
	}
	*left = root
	if root.Right != nil {
		traversal(root.Right, arr, left)
	}
}

//124 最大路径和,竟然这么简单，一个递归就行了
func maxPathSum(root *TreeNode) int {
	ret := INT_MIN
	maxPathTrace(root, &ret)
	return ret
}
func maxPathTrace(root *TreeNode, ret *int) int {
	left := 0
	right := 0
	bigger := 0
	if root.Left != nil {
		left = maxPathTrace(root.Left, ret)
	}
	if root.Right != nil {
		right = maxPathTrace(root.Right, ret)
	}

	if left+right+root.Val > *ret {
		*ret = left + right + root.Val
	}
	if root.Val > *ret {
		*ret = root.Val
	}
	if left+root.Val > *ret {
		*ret = left + root.Val
	}
	if right+root.Val > *ret {
		*ret = right + root.Val
	}
	if left > right {
		bigger = left
	} else {
		bigger = right
	}
	if bigger+root.Val > root.Val {
		return bigger + root.Val
	} else {
		return root.Val
	}

}

//199 广度优先遍历,真的就是这么简单
func rightSideView(root *TreeNode) []int {
	record := make([]*TreeNode, 0)
	current := 1
	next := 0
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	record = append(record, root)
	for len(record) > 0 {
		tmp := record[0]
		if tmp.Left != nil {
			record = append(record, tmp.Left)
			next++
		}
		if tmp.Right != nil {
			record = append(record, tmp.Right)
			next++
		}

		current--
		if current == 0 {
			ret = append(ret, tmp.Val)
			current = next
			next = 0
		}
		if len(record) > 1 {
			length := len(record)
			record = record[1:length]
		} else {
			break
		}
		//fmt.Println(record[0])
	}
	return ret
}
