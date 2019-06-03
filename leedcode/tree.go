package leedcode

import "math"

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
	}
	return ret
}
//pid 98 二叉搜索树 可以用中序遍历或者递归
func isValidBST(root *TreeNode) bool {
	return valid(root,INT_MIN,INT_MAX)
}
func valid(root *TreeNode,min int, max int) bool{
	if root==nil {
		return true
	}
	if root.Val<=min || root.Val>=max {
		return false
	}

	if root.Right!=nil&&root.Right.Val<=root.Val {
		return false
	}
	return valid(root.Left,min,root.Val)&&valid(root.Right,root.Val,max)
}
//中序遍历
func inorderTraversal(root *TreeNode) []int {
	res:=make([]int,0)
	inorder(root,&res)
	return res
}

func inorder(root *TreeNode,res *[]int){
	if root==nil {
		return
	}
	if root.Left!=nil {
		inorder(root.Left,res)
	}
	*res=append(*res,root.Val)
	if root.Right!=nil {
		inorder(root.Right,res)
	}
}
//pid 101 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root==nil {
		return true
	}
	return sym(root.Left,root.Right)

}
func sym(left *TreeNode,rig *TreeNode) bool{
	if left==nil&&rig==nil {
		return true
	}
	if left!=nil&&rig==nil || left==nil&&rig!=nil || left.Val!=rig.Val {
		return false
	}
	return sym(left.Left,rig.Right)&&sym(left.Right,rig.Left)
}
//pid 104 最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := maxDepth(root.Left)+1
	rig:=maxDepth(root.Right)+1
	if left>rig {
		return left
	} else {
		return rig
	}
}
//pid 102层次遍历
func levelOrder(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil {
		return res
	}
	level:=make([]*TreeNode,0)
	level=append(level,root)
	levelArr:=make([]int,0)

	length:=len(level)
	index:=0
	for index<length {
		tmp:=level[0]
		levelArr=append(levelArr,tmp.Val)
		index++
		level=level[1:]
		if tmp.Left!=nil {
			level=append(level,tmp.Left)
		}
		if tmp.Right!=nil {
			level=append(level,tmp.Right)
		}
		if index==length{
			item := make([]int, len(levelArr))
			copy(item, levelArr)
			res=append(res,item)
			levelArr=levelArr[0:0]
			length=len(level)
			index=0
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
//pid 层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res:=make([][]int,0)
	if root==nil {
		return res
	}
	level:=make([]*TreeNode,0)
	level=append(level,root)
	levelArr:=make([]int,0)

	length:=len(level)
	index:=0
	for index<length {
		tmp:=level[0]
		levelArr=append(levelArr,tmp.Val)
		index++
		level=level[1:]
		if tmp.Left!=nil {
			level=append(level,tmp.Left)
		}
		if tmp.Right!=nil {
			level=append(level,tmp.Right)
		}
		if index==length{
			item := make([]int, len(levelArr))
			copy(item, levelArr)
			if len(res)%2==1{
				for i, j := 0, len(item)-1; i < j; i, j = i+1, j-1 {
					item[i], item[j] = item[j], item[i]
				}
			}
			res=append(res,item)
			levelArr=levelArr[0:0]
			length=len(level)
			index=0
		}
	}
	return res
}
//pid 82 删除重复元素
func deleteDuplicates2(head *ListNode) *ListNode {
	if head==nil || head.Next==nil {
		return head
	}
	tmp:=new(ListNode)
	tmp.Next = head
	res:=tmp
	fir:=head
	sed:=head.Next
	for sed!=nil {
		if sed.Val!=fir.Val {
			tmp=fir
			fir=fir.Next
			sed=sed.Next
			continue
		}
		for sed!=nil && sed.Val==fir.Val {
			sed=sed.Next
		}
		tmp.Next=sed
		fir=tmp.Next
		if fir!=nil {
			sed=fir.Next
		}
	}
	return res.Next
}
//pid 判断平衡树
func isBalanced(root *TreeNode) bool {
	if checkDepth(root)==-1 {
		return false
	}
	return true
}
func checkDepth(root *TreeNode) int {
	if root==nil {
		return 0
	}
	left:=checkDepth(root.Left)
	if left==-1 {
		return -1
	}
	right:=checkDepth(root.Right)
	if right==-1 {
		return -1
	}
	diff:=math.Abs(float64(left-right))
	if diff>1 {
		return -1
	} else {
		return  1+max(left,right)
	}
}
