package leetcode

import (
	"awesomeProject5/base"
	"fmt"
)

// 翻转链表
func ReveseExample() {
	list := base.NewLinkList()
	list.Add(0, "hello")
	list.Add(1, "hello1")
	list.Add(2, "hello2")
	list.Add(3, "hello3")
	list.Add(4, "hello4")
	PrintNode(list.GetRoot())
	println()
	PrintNode(RevList(list.GetRoot()))

}
func PrintNode(node *base.Node) {
	for node != nil {
		fmt.Println(node.Element)
		node = node.Next
	}
}

// 翻转链表1->2->3-4>-5>null  5 4 3 2-1>null
func RevList(node *base.Node) (head *base.Node) {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}
	head = RevList(node.Next)
	node.Next.Next = node
	node.Next = nil
	return head
}

// 检查链表是否有环
func checkCircle(list base.DoubleLinkList) {
	root := list.Front()
	fast := root.Next()
	slow := root
	for !(fast.Next() == nil && fast.Next().Next() == nil) {
		if fast == slow {
			fmt.Println("true ")
			break
		}
		fast = fast.Next().Next()
		slow = slow.Next()
	}
}
