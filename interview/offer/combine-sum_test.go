package offer

import "testing"

func TestCombineSum4(t *testing.T) {
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
			name: "nums=[1,2,3], target=4",
			args: args{nums: []int{1, 2, 3}, target: 4},
			want: 7,
		},
		{
			name: "nums=[9], target=3",
			args: args{nums: []int{9}, target: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("CombineSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
