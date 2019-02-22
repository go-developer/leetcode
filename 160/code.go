package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 核心 ： 单链表环处理
 * @author go_developer@163.com <张德满>
 */
func getIntersectionNode(headA *ListNode, headB *ListNode) *ListNode {
	var (
		tmpHead *ListNode
	)

	tmpHead = headA

	//将 headB 挂到 headA末尾，如果 A B 相交，则新形成的单链表有环
	for {
		if nil == tmpHead.Next {
			tmpHead.Next = headB
			break
		}
		tmpHead = tmpHead.Next
	}

	//判断单链表是否有环 ： 快慢指针法
}
