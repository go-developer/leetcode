package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	head = deleteDuplicates(head)
	for {
		if nil != head {
			fmt.Println(head.Val)
			head = head.Next
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
 * 重新调整指针指向即可 1 1 2 3 3
 * @author go_developer@163.com <张德满>
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}

	headPoint := head

	// 1 1 2 3 3
	for {
		if nil == head.Next {
			break
		}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		} else {
			head = head.Next
		}
	}

	return headPoint
}
