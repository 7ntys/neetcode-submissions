/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) == 0 {return nil}

	curr := preorder[0]

	node := &TreeNode{Val: curr}
	i := 0
	for i < len(inorder) {
		if inorder[i] == curr {break}
		i++
	}

	// Here we have : 
	// preorder = prorder[1:]
	// left side : inorder[:i]
	// right side : inorder[i+1:]

	node.Left = buildTree(preorder[1:], inorder[:i])

	if i + 1 < len(inorder) {
		node.Right = buildTree(preorder[1 + i:], inorder[i+1:])
	}
	return node
}
