// Package _4...
//
// Description : _4...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-12 11:56 下午
package main

import "fmt"

func main() {
	data := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(data))
	fmt.Println(inorderTraversalWithStack(data))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// inorderTraversal 二叉树中序遍历递归实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:57 下午 2021/2/12
func inorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if nil == root {
		return r
	}

	mid(&r, root)
	return r
}

func mid(res *[]int, root *TreeNode) {
	//先遍历左子树
	if root.Left != nil {
		mid(res, root.Left)
	}
	//再遍历自己
	*res = append(*res, root.Val)
	//最后遍历右子树
	if root.Right != nil {
		mid(res, root.Right)
	}
}

// inorderTraversalWithStack 中序遍历迭代实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 1:13 上午 2021/2/13
func inorderTraversalWithStack(root *TreeNode) []int {
	// 中序遍历的顺序是左中右
	// 本算法采用迭代的方法完成，采用的核心工具是栈
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	flag := false
	for len(stack) != 0 {
		if flag {
			// 当前路线左子树已经遍历完
			rightChild := stack[len(stack)-1].Right
			res = append(res, stack[len(stack)-1].Val)
			stack = stack[:len(stack)-1]
			// 说明有右子树，所以将右子树的根结点入栈
			if rightChild != nil {
				stack = append(stack, rightChild)
				// 重新初始化标志，表示该根节点的左子树还没有遍历完
				flag = false
			}
			continue
		}
		leftChild := stack[len(stack)-1].Left
		if leftChild != nil {
			// 左子树入栈
			stack = append(stack, leftChild)
			continue
		}
		// 左子树为空，说明左子树已经遍历完毕
		flag = true
	}
	return res
}
