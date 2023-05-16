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
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}

	funcs := []func(*ListNode) *ListNode{ReverseListRec, ReverseListItr}

	for _, datum := range data {
		for _, fnc := range funcs {

			input := NewList(datum.input)
			want := NewList(datum.want)

			if got := fnc(input); !reflect.DeepEqual(got, want) {
				t.Errorf("%s(%v) = %v, want %v", GetFunctionName(fnc), input.ToSlice(), got.ToSlice(), want.ToSlice())
			}
		}
	}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
