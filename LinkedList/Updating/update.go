package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func toSlice(head *Node, nums []int) []int {
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
func modifySlice(nums []int) []int {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1]-nums[i], nums[i]
	}
	return nums
}
func toLinkedList(nums []int, head *Node) {
	curr := head
	for i := 0; i < len(nums); i++ {
		curr.Val = nums[i]
		curr = curr.Next
	}
}
func print(head *Node) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
}
func main() {
	node := &Node{Val: 10, Next: nil}
	node.Next = &Node{Val: 4, Next: nil}
	node.Next.Next = &Node{Val: 5, Next: nil}
	node.Next.Next.Next = &Node{Val: 3, Next: nil}
	node.Next.Next.Next.Next = &Node{Val: 6, Next: nil}
	var nums []int
	nums = toSlice(node, nums)
	modifySlice(nums)
	toLinkedList(nums, node)
	print(node)
}
