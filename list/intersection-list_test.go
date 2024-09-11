package list

import (
	"testing"
)

/*
  1->2->3
            ->4->5->6
     2->3
*/

func buildIntersectionList(prefixNums1, prefixNums2 []int, commonNums []int) (*ListNode, *ListNode) {
	list1 := NewListFromArray(prefixNums1)
	list2 := NewListFromArray(prefixNums2)

	commonList := NewListFromArray(commonNums)

	return ConcatenateList(list1, commonList), ConcatenateList(list2, commonList)
}

func TestIntersectionNode(t *testing.T) {
	type args struct {
		prefixNums1 []int
		prefixNums2 []int
		commonNums  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "intersection node",
			args: args{
				prefixNums1: []int{1, 2, 3},
				prefixNums2: []int{2, 3},
				commonNums:  []int{4, 5, 6},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headA, headB := buildIntersectionList(tt.args.prefixNums1, tt.args.prefixNums2, tt.args.commonNums)

			if got := IntersectionNode(headA, headB); got == nil || got.Val != tt.want {
				t.Errorf("IntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
