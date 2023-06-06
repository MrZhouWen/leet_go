package 二叉树

/**
经典问题：二叉树的后序遍历
T= O(n) S= O(n)

 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func postorderTraversal(root *TreeNode) (res []int) {
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}
	postOrder(root)
	return
}
