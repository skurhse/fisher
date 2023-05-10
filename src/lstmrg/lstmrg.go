package lstmrg

import cl "container/list"

func Merge(listOne *cl.List, listTwo *cl.List) *cl.List {

	if listOne.Front() == nil {
		return listTwo
	}
	if listTwo.Front() == nil {
		return listOne
	}

	if listOne.Front().Value.(int) > listTwo.Front().Value.(int) {
		listOne, listTwo = listTwo, listOne
	}

	frontElement := listTwo.Front()
	listOne.PushBack(frontElement.Value)
	listTwo.Remove(frontElement)

	return Merge(listOne, listTwo)
}
