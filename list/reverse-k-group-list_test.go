package list

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "10 items in list, k=3",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				k:    3,
			},
			want: []int{3, 2, 1, 6, 5, 4, 9, 8, 7, 10},
		},
		{
			name: "12 items in list, k=4",
			args: args{
				head: NewListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
				k:    4,
			},
			want: []int{4, 3, 2, 1, 8, 7, 6, 5, 12, 11, 10, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("source array: %v", ToArray(tt.args.head))

			if got := ReverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("ReverseKGroup() = %v, want %v", ToArray(got), tt.want)
			}
		})
	}
}
