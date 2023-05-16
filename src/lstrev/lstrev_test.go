package lstrev_test

import (
	"reflect"
	"runtime"
	"testing"

	. "github.com/skurhse/gresto/lstrev"
)

type assert struct {
	input []int
	want  []int
}

func TestReverseLists(t *testing.T) {
	var data = []assert{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	funcs := []func(*ListNode) *ListNode{ReverseListRec, ReverseListItr}

	for _, dtm := range data {
		for _, fnc := range funcs {
			input, want := NewList(dtm.input), NewList(dtm.want)

			if got := fnc(input); !reflect.DeepEqual(got, want) {
				fname := runtime.FuncForPC(reflect.ValueOf(fnc).Pointer()).Name()
				t.Errorf("%s(%v) = %v, want %v", fname, input.ToSlice(), got.ToSlice(), want.ToSlice())
			}
		}
	}
}
