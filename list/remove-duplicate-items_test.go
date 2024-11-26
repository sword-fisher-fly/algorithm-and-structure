package list

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicatesRemainOnlyExistOnce(t *testing.T) {
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
			args: args{
				head: []int{1, 1, 1, 2, 3, 3},
			},
			want: []int{2},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 1, 1, 2, 2, 3, 3},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteDuplicatesRemainOnlyExistOnce(NewListFromArray(tt.args.head)); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("DeleteDuplicatesRemainOnlyExistOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicatesRemainOnlyOne(t *testing.T) {
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
			args: args{
				head: []int{1, 1, 1, 2, 3, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "case 2",
			args: args{
				head: []int{1, 1, 1, 2, 2, 3, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesRemainOnlyOne(NewListFromArray(tt.args.head)); !reflect.DeepEqual(ToArray(got), tt.want) {
				t.Errorf("RemoveDuplicatesRemainOnlyOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
