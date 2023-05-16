package lstrev_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"

	. "github.com/skurhse/gresto/lstrev"
)

type assert struct {
	input []int
	want  []int
}

var funcs []func(*ListNode) *ListNode = []func(*ListNode) *ListNode{ReverseListRec, ReverseListItr}

func TestReverseLists(t *testing.T) {
	var data = []assert{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, dtm := range data {
		for _, fnc := range funcs {
			input, want := NewList(dtm.input), NewList(dtm.want)

			fname := runtime.FuncForPC(reflect.ValueOf(fnc).Pointer()).Name()
			testname := fmt.Sprintf("%s(%v)", fname, input.ToSlice())

			t.Run(testname, func(t *testing.T) {
				if got := fnc(input); !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got.ToSlice(), want.ToSlice())
				}
			})
		}
	}
}

func BenchmarkReverseLists(b *testing.B) {
	rand.Seed(time.Now().Unix())
	list := NewList(rand.Perm(1000000))

	for _, fnc := range funcs {
		fname := runtime.FuncForPC(reflect.ValueOf(fnc).Pointer()).Name()

		b.Run(fname, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fnc(list)
			}
		})
	}
}
