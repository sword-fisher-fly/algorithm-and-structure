package dynamic

import "testing"

func TestMaxSubArray(t *testing.T) {
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
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "case 2",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
