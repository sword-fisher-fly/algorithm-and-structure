package greedy

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
			name: "no negative items",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "array hse negative items",
			args: args{nums: []int{-2, -3, 4, -1, -2, 1, 5, -3}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
			}

			if got := MaxSubArrayII(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}
