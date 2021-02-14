// Package _00...
//
// Description : _00...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-13 11:20 下午
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree 深度优先对比
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:23 下午 2021/2/13
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 全部是nil，相等
	if nil == p && nil == q {
		return true
	}
	// 只有一个为nil,不相等
	if nil == p || nil == q {
		return false
	}

	// 值不相等
	if p.Val != q.Val {
		return false
	}

	// 递归对比左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// isSameTreeTwo 广度优先
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:28 下午 2021/2/13
func isSameTreeTwo(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}

	if nil == p || nil == q {
		return false
	}

	treeQueue1, treeQueue2 := []*TreeNode{p}, []*TreeNode{q}

	// 两个队列均长度均大于0
	for len(treeQueue1) > 0 && len(treeQueue2) > 0 {
		if len(treeQueue1) != len(treeQueue2) {
			// 队列长度不一样，说明树一定不相等
			return false
		}
		node1, node2 := treeQueue1[0], treeQueue2[0]
		if node1.Val != node2.Val {
			// 节点值不相等
			return false
		}
		treeQueue1, treeQueue2 = treeQueue1[1:], treeQueue2[1:]
		leftNode1, rightNode1 := node1.Left, node1.Right
		leftNode2, rightNode2 := node2.Left, node2.Right
		if (leftNode1 == nil && leftNode2 != nil) || (leftNode1 != nil && leftNode2 == nil) {
			return false
		}
		if (rightNode1 == nil && rightNode2 != nil) || (rightNode1 != nil && rightNode2 == nil) {
			return false
		}

		if nil != leftNode1 {
			treeQueue1 = append(treeQueue1, leftNode1)
		}

		if nil != rightNode1 {
			treeQueue1 = append(treeQueue1, rightNode1)
		}

		if nil != leftNode2 {
			treeQueue2 = append(treeQueue2, leftNode2)
		}

		if nil != rightNode2 {
			treeQueue2 = append(treeQueue2, rightNode2)
		}
	}

	return len(treeQueue1) == 0 && len(treeQueue2) == 0
}
