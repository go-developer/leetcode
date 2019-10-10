package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 删除链表倒数第n个节点
 * @author go_developer@163.com
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	inputHead := head
	dataList := make([]*ListNode, 0)
	for nil != head.Next {
		dataList = append(dataList, head)
		head = head.Next
	}
	dataList = append(dataList, head)
	if n > len(dataList) {
		return nil
	}
	if n == 0 {
		return inputHead
	}

	if n == len(dataList) {
		if n == 1 {
			return nil
		}
		return dataList[1]
	}

	if n == 1 {
		dataList[len(dataList) - 2].Next = nil
		return inputHead
	}
	dataList[len(dataList) - n - 1].Next = dataList[len(dataList) - n + 1]
	return inputHead
}
