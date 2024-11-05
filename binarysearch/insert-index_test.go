package binarysearch

import "testing"

func TestInsertIndex(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums=[1,3,5,6], target=5",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "nums=[1,3,5,6], target=2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "nums=[1,3,5,6], target=7",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertIndex(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("InsertIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
