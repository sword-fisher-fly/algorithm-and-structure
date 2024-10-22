package list

import (
	"reflect"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{head: NewListFromArray([]int{1, 2, 3, 4, 5})},
			want: []int{1, 3, 5, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddEvenList(tt.args.head); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("OddEvenList() = %v, want %v", ToArray(got), tt.want)
			}
		})
	}
}
