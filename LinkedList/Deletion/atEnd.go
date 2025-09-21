package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func deleteAtEnd(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	return head
}
func traverseList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print("->")
		}
		current = current.Next
	}
	fmt.Println()
}
func main() {
	node := &Node{Val: 8, Next: nil}
	node.Next = &Node{Val: 18, Next: nil}
	node.Next.Next = &Node{Val: 20, Next: nil}
	node.Next.Next.Next = &Node{Val: 100, Next: nil}
	traverseList(node)
	node = deleteAtEnd(node)
	traverseList(node)
}
