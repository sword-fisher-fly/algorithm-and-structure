package special

import (
	"reflect"
	"testing"
)

func int2List(nums []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return dummyHead.Next
}

func list2Int(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	res := []int{}
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

func TestRemoveZeroSumSublist(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{head: []int{0, 1, 1, -3, 6, -3, 0, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "case 2",
			args: args{head: []int{1, 2, -2, 5, -3, 1}},
			want: []int{1, 5, -3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveZeroSumSublist(int2List(tt.args.head)); !reflect.DeepEqual(list2Int(got), tt.want) {
				t.Errorf("RemoveZeroSumSublist() = %v, want %v", list2Int(got), tt.want)
			}
		})
	}
}
