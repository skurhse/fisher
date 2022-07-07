package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals []int) (head *ListNode) {
	if len(vals) == 0 {
		return
	}

	head = &ListNode{
		Val: vals[0],
	}
	if len(vals) == 1 {
		return
	}

	prev := head

	for _, val := range vals[1 : len(vals)-1] {
		next := &ListNode{
			Val: val,
		}
		prev.Next = next
		prev = next
	}

	next := &ListNode{
		Val: vals[len(vals)-1],
	}
	prev.Next = next

	return
}

func PrintList(list *ListNode) {
	fmt.Printf("%v\n", sliceFromList(list))
}

func sliceFromList(list *ListNode) (slice []int) {
	if list == nil {
		return []int{}
	}

	slice = append(slice, list.Val)
	slice = append(slice, sliceFromList(list.Next)...)

	return
}
