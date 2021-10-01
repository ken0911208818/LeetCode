package removeDuplicatesFromSortedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	input
	output
}

type input struct {
	l1 *ListNode
}

type output struct {
	*ListNode
}

func Test_mergeTwoLists(t *testing.T) {
	qs := []question{
		question{
			input: input{
				l1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{},
						},
					},
				},
			},
			output: output{
				ListNode: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{},
					},
				},
			},
		},
	}

	for _, v := range qs {
		assert.Equal(t, deleteDuplicates(v.l1), v.output.ListNode)
	}
}

var testQS = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{},
		},
	},
}

func BenchmarkSelf(t *testing.B) {
	for i := 0; i < t.N; i++ {
		deleteDuplicates(testQS)
	}
}

func Benchmark01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		deleteDuplicates01(testQS)
	}
}
