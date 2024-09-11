package list

import (
	"reflect"
	"testing"
)

func TestRemoveNthItemFromTheBottom(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		{
			name: "1th from bottom in 10 items array",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n:   1,
			},
			want:    10,
			wantArr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "3th from bottom in 10 items array",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n:   3,
			},
			want:    8,
			wantArr: []int{1, 2, 3, 4, 5, 6, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromArray(tt.args.arr)

			t.Logf("list to array: %v", ToArray(head))
			if got := RemoveNthItemFromTheBottomII(head, tt.args.n); !reflect.DeepEqual(ToArray(got), tt.wantArr) {
				t.Errorf("RemoveNthItemFromTheBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
