package 二叉树

/**
题目：将一棵二叉树翻转

思路：递归翻转
先分别将 root的左右子树翻转，再交换root的左右子树的指针
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
