package offer

import "testing"

func TestLongestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single number",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "10 items with 5 consecutive nums",
			args: args{nums: []int{5, 3, 2, 4, 1, 7, 10, 11, 15}},
			want: 5,
		},
		{
			name: "10 items with repeated nums",
			args: args{nums: []int{5, 2, 2, 4, 1, 7, 10, 11, 15, 15}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
