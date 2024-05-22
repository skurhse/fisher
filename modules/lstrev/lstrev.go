package lstrev

type ListNode struct {
	Value int
	Next  *ListNode
}

func ReverseListRec(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(prev *ListNode, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev

	return reverse(curr, next)
}

func ReverseListItr(head *ListNode) *ListNode {
	var prev, next *ListNode
	for curr := head; curr != nil; prev, curr = curr, next {
		next = curr.Next
		curr.Next = prev
	}

	return prev
}

func NewList(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}

	h := &ListNode{s[0], nil}
	n := h
	for _, v := range s[1:] {
		n.Next = &ListNode{v, nil}
		n = n.Next
	}
	return h
}

func (h *ListNode) ToSlice() []int {
	if h == nil {
		return []int{}
	}

	s := make([]int, 0)
	for n := h; n != nil; n = n.Next {
		s = append(s, n.Value)
	}
	return s
}
