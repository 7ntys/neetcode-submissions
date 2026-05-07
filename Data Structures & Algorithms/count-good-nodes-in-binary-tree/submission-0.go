/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    return isGoodNode(root, root.Val)
}

func isGoodNode(node *TreeNode, threshold int) int {
	if node == nil {return 0}

	goods := 0
	if threshold <= node.Val {
		goods++
		threshold = node.Val
	}

	goods += isGoodNode(node.Left, threshold)
	goods += isGoodNode(node.Right, threshold)
	return goods
}