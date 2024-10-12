package list

import (
	"reflect"
	"testing"
)

func TestMergeSortedListByIteration(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				l1: sortedList1,
				l2: sortedList2,
			},
			want: expectedMergedSortedListArray,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedListByIteration(tt.args.l1, tt.args.l2); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("MergeSortedListByIteration() = %v, want %v", ToArray(got), tt.want)
			}
		})
	}
}

func TestMergeSortedListByRecursive(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				l1: NewListFromArray(sortedArray1),
				l2: NewListFromArray(sortedArray2),
			},
			want: expectedMergedSortedListArray,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSortedListByRecursive(tt.args.l1, tt.args.l2); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("MergeSortedListByRecursive() = %v, want %v", ToArray(got), tt.want)
			}
		})
	}
}
