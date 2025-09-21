package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func deleteAtBegin(head *Node) *Node {
	if head == nil {
		return head
	}
	head = head.Next
	return head
}

func main() {
	node := &Node{Val: 8, Next: nil}
	node.Next = &Node{Val: 18, Next: nil}
	node.Next.Next = &Node{Val: 20, Next: nil}
	fmt.Println(node.Val, node.Next.Val, node.Next.Next.Val)
	node = deleteAtBegin(node)
	fmt.Println(node.Val, node.Next.Val)
}
