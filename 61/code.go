package main

type ListNode struct {
	Val int
	Next *ListNode
}
func main() {
	
}

/**
 * 简单易懂逻辑
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if nil == head || k == 0 {
		return head
	}
	//链表各个索引的指针
	nodeMap := make(map[int]*ListNode)
	tmpIndex := 0

	for {
		nodeMap[tmpIndex] = head
		tmpIndex = tmpIndex + 1
		if head = head.Next; nil == head {
			break
		}
	}
	head = nodeMap[0]
	offset := k % len(nodeMap)
	if offset == 0 {
		return head
	}
	nodeMap[len(nodeMap) - 1].Next = head
	nodeMap[len(nodeMap) - offset - 1].Next = nil
	return nodeMap[len(nodeMap) - offset]
}

/**
 * 先将链表连成环，在根据旋转的数量找到断开点，即可获得新的线性链表
 */
func rotateRight2(head *ListNode, k int) *ListNode {
	if nil == head || k == 0 {
		return head
	}
	listLen := 0
	headPtr := head
	tmpPtr := head
	for {
		listLen = listLen + 1
		if nil == head.Next {
			head.Next = headPtr
			break
		}
		head = head.Next
	}
	offset := k % listLen
	for step := 0; step < listLen - offset - 1; step ++ {
		tmpPtr = tmpPtr.Next
	}

	newHead := tmpPtr.Next
	tmpPtr.Next = nil

	return newHead
}