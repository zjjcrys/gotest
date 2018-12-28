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

//
func maxPathSum(root *TreeNode) int {

}
