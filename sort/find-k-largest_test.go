package sort

import "testing"

func TestFindKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "case 1",
		// 	args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
		// 	want: 5,
		// },
		{
			name: "case 2",
			args: args{nums: []int{3, 2, 1, 5, 6, 4}, k: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
