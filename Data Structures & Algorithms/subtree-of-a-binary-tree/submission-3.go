/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {return true}
	if root == nil {return false}

	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(ref *TreeNode, sub *TreeNode) bool {
	if ref == nil && sub == nil {return true}
	if ref == nil || sub == nil {return false}

	if ref.Val == sub.Val {
		return isSame(ref.Left, sub.Left) && isSame(ref.Right, sub.Right)
	}
	return false
}
