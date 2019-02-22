package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:4,
				Next:&ListNode{
					Val:5,
					Next:nil,
				},
			},
		},
	}

	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val: 3,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}

	merge := mergeTwoLists(l1, l2)
	for  {
		if nil != merge {
			fmt.Println(merge.Val)
			merge = merge.Next
		} else {
			break
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 合并两个有序数组，且合并后依然有序
 * @author go_developer@163.com <张德满>
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}

	if nil == l2 {
		return l1
	}

	var (
		result *ListNode
		originalPoint *ListNode
	)

	result = &ListNode{}
	originalPoint = result

	for  {
		if nil == l1 {
			result.Val = l2.Val
			l2 = l2.Next
		} else if nil == l2 {
			result.Val = l1.Val
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				result.Val = l1.Val
				l1 = l1.Next
			} else {
				result.Val = l2.Val
				l2 = l2.Next
			}
		}

		if nil == l1 && nil == l2 {
			//合并已完成
			result.Next = nil
			break
		}
		result.Next = &ListNode{}
		result = result.Next
	}
	return originalPoint
}
