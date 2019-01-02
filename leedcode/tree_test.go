package leedcode

import (
	"fmt"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 3
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 2
	fmt.Println(root, root.Left, "before------------------------")
	recoverTree(root)
	fmt.Println("--------------------------------after", root, root.Left, root.Left.Right)
}

func TestRightSideView(t *testing.T) {
	root := new(TreeNode)
	root.Val = 1
	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 3
	fmt.Println(rightSideView(root))
}
