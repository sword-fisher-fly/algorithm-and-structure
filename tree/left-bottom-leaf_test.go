package tree

import (
	"reflect"
	"testing"
)

func TestFindLeftBottomLeafByRecursive(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only root tree node",
			args: args{t: NewTreeFromArray([]any{1})},
			want: 1,
		},
		{
			name: "2-level full tree",
			args: args{t: level2FullTree},
			want: 2,
		},
		{
			name: "4-level full tree",
			args: args{t: level4FullTree},
			want: 8,
		},
		{
			name: "random tree",
			args: args{t: level4RandomTree},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindLeftBottomLeafByRecursive(tt.args.t.root)
			if got == nil || !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("FindLeftBottomLeafByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLeftBottomLeafByIteration(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "2-level full tree",
			args: args{t: level2FullTree},
			want: 2,
		},
		{
			name: "4-level full tree",
			args: args{t: level4FullTree},
			want: 8,
		},
		{
			name: "random tree",
			args: args{t: level4RandomTree},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindLeftBottomLeafByIteration(tt.args.t.root)
			if got == nil || !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("FindLeftBottomLeafByIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
