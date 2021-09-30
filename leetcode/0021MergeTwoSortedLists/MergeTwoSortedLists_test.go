package mergetwosortedlist

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
	l2 *ListNode
}

type output struct {
	*ListNode
}

func Test_mergeTwoLists(t *testing.T) {
	qs := []question{
		{
			input{l1: nil, l2: nil},
			output{nil},
		},
		{
			input{
				l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
				},
				l2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				}},
			output{&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},	
					},
				},
			}},
		},
	}

	for _, v := range qs {
		assert.Equal(t, mergeTwoLists(v.l1, v.l2), v.output.ListNode)
	}
}
