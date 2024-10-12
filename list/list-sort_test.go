package list

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sort random array 1",
			args: args{head: randomList1},
			want: expectedSortedRandomArray1,
		},
		{
			name: "sort random array 2",
			args: args{head: randomList2},
			want: expectedSortedRandomArray2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortList(tt.args.head); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("SortList() = %v, want %v", ToArray(got), tt.want)
			}
		})
	}
}
