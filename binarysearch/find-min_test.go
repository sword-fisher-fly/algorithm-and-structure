package binarysearch

import "testing"

func TestFindMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMin(tt.args.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
