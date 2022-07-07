package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) (head *ListNode) {
	if len(vals) == 0 {
		return
	}

	head = &ListNode{
		Val: vals[0],
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
