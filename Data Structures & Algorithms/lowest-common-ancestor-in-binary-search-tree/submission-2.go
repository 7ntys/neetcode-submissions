/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    return dfs(root, p, q)
}

func dfs(r *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if r == nil {return nil}

	if r.Val == p.Val || r.Val == q.Val {
		return r
	}

	left := dfs(r.Left, p, q)
	right := dfs(r.Right, p, q)

	if left != nil && right != nil {return r}

	if left != nil {
		return left
	}
	return right
}