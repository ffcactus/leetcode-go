package unique_binary_search_trees_ii

import "fmt"

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode is a definition of a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (n TreeNode) String() string {
	return fmt.Sprintf("value = %d", n.Val)
}

func generateTrees(n int) []*TreeNode {
	return dp(1, n)
}

func dp(from, to int) []*TreeNode {
	if to < from {
		return nil
	}
	if to - from == 0 {
		ret := []*TreeNode{&TreeNode{
			Val:   to,
			Left:  nil,
			Right: nil,
		}}
		return ret
	}
	ret := make([]*TreeNode, 0)
	for i:=from; i <= to; i++ {
		left := dp(from, i - 1)
		right := dp(i + 1, to)
		if left != nil && right != nil {
			for _, eachLeft := range left {
				for _, eachRight := range right {
					n := TreeNode{
						Val:   i,
						Left:  eachLeft,
						Right: eachRight,
					}
					ret = append(ret, &n)
				}
			}
		}
		if left != nil && right == nil {
			for _, eachLeft := range left {
				n := TreeNode{
					Val:   i,
					Left:  eachLeft,
					Right: nil,
				}
				ret = append(ret, &n)
			}
		}
		if left == nil && right != nil {
			for _, eachRight := range right {
				n := TreeNode{
					Val:   i,
					Left:  nil,
					Right: eachRight,
				}
				ret = append(ret, &n)
			}
		}
	}
	return ret
}