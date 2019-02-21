package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	for nil != result {
		fmt.Println(result.Val)
		result = result.Next
	}


	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: nil,
		},
	}

	l2 = &ListNode{
		Val: 0,
		Next: nil,
	}

	result = addTwoNumbers(l1, l2)
	for nil != result {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 逐位相加
 * @author go_developer@163.com <张德满>
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result        *ListNode
		originalPoint *ListNode
		plusVal       int
	)

	plusVal = 0

	result = &ListNode{}
	originalPoint = result

	for {

		if nil == l1 && nil == l2 {
			//链表均处理完成
			break
		}

		val := l1.Val + l2.Val + plusVal

		if val >= 10 {
			val = val - 10
			plusVal = 1
		} else {
			plusVal = 0
		}

		result.Val = val
		if nil != l1.Next || nil != l2.Next {
			result.Next = &ListNode{}
			result = result.Next
		} else {
			result.Next = nil
		}

		l1 = l1.Next
		l2 = l2.Next
		//链表未处理完成，补齐
		if l1 == nil && nil != l2 {
			l1 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}

		if nil == l2 && nil != l1 {
			l2 = &ListNode{
				Val:  0,
				Next: nil,
			}
		}
	}

	if plusVal > 0 {
		result.Next = &ListNode{
			Val:  plusVal,
			Next: nil,
		}
	}

	return originalPoint
}
