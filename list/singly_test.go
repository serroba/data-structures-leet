package list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				list1: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 4,
						}},
				}, list2: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{Val: 1, Next: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4, Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			}},
		},
		{
			name: "example 2",
			args: args{
				list1: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 3},
				}, list2: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 4,
					},
				},
			},
			want: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
