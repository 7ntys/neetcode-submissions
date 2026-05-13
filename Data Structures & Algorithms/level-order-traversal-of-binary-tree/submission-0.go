/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    ans := [][]int{}
	if root == nil {return ans}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		level := []int{}
		for i:=0;i<size;i++{
			x := q[i]
			level = append(level, x.Val)
			if x.Left != nil {q = append(q, x.Left)}
			if x.Right != nil {q = append(q, x.Right)}
		}
		ans = append(ans, level)
		q = q[size:]
	}
	return ans
}
