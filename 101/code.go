// Package _01...
//
// Description : _01...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2021-02-14 4:25 下午
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric 递归实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:37 下午 2021/2/14
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if nil == root.Left && nil == root.Right {
		return true
	}

	if nil == root.Left || nil == root.Right {
		return false
	}
	return check(root.Left, root.Right)
}

// check 递归检查
//
// Author : go_developer@163.com<张德满>
//
// Date : 7:46 下午 2021/2/14
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// 根节点值相等,左节点值等于右节点值
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// isSymmetricLoop 循环实现
//
// Author : go_developer@163.com<张德满>
//
// Date : 7:57 下午 2021/2/14
func isSymmetricLoop(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if nil == root.Left && nil == root.Right {
		return true
	}

	if nil == root.Left || nil == root.Right {
		return false
	}

	queue1, queue2 := []*TreeNode{root.Left}, []*TreeNode{root.Right}

	for len(queue1) > 0 && len(queue2) > 0 {
		if len(queue1) != len(queue2) {
			// 数量不等,不可能是镜像
			return false
		}
		node1, node2 := queue1[0], queue2[0]
		if node1.Val != node2.Val {
			// 节点的根植不相等
			return false
		}
		if node1.Left != nil && node2.Right == nil || node1.Right != nil && node2.Left == nil {
			return false
		}

		if node1.Left == nil && node2.Right != nil || node1.Right == nil && node2.Left != nil {
			return false
		}
		queue1, queue2 = queue1[1:], queue2[1:]
		if node1.Left != nil {
			queue1 = append(queue1, node1.Left)
		}
		if node1.Right != nil {
			queue1 = append(queue1, node1.Right)
		}
		if node2.Right != nil {
			queue2 = append(queue2, node2.Right)
		}
		if node2.Left != nil {
			queue2 = append(queue2, node2.Left)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
