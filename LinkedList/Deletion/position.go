package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func deleteAtPosition(head *Node, pos int) *Node {
	if pos < 1 {
		return head
	}
	if pos == 1 {
		return head.Next
	}
	curr := head
	var ptr *Node
	for i := 1; i < pos && curr != nil; i++ {
		ptr = curr
		curr = curr.Next
	}
	if curr == nil {
		return head
	}
	ptr.Next = curr.Next
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

// 1 2 3 5
func main() {
	node := &Node{Val: 8, Next: nil}
	node.Next = &Node{Val: 18, Next: nil}
	node.Next.Next = &Node{Val: 20, Next: nil}
	node.Next.Next.Next = &Node{Val: 100, Next: nil}
	traverseList(node)
	node = deleteAtPosition(node, 4)
	traverseList(node)
}
