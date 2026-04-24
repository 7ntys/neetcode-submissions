/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return getHeight(root) != -1
}

func getHeight(node *TreeNode) int {
    if node == nil {return 0}

    left := getHeight(node.Left)
    right := getHeight(node.Right)

    if left == -1 || right == -1 {return -1}
    if abs(left + 1, right + 1) > 1 {return -1}
    return max(left,right) + 1
}

func abs(x,y int) int {
    if x > y {
        return x-y
    }
    return y-x
}