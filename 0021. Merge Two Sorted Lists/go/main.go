package main

import (
	. "linkedlist"
)

func main() {

	lists := [3][2][]int{
		{{1, 2, 3, 4}, {1, 2, 3, 4}},
		{{}, {}},
		{{}, {0}},
	}

	for _, v := range lists {
		mergedList := mergeTwoLists(NewList(v[0]), NewList(v[1]))
		PrintList(mergedList)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}

	list1.Next = mergeTwoLists(list1.Next, list2)
	return list1
}
