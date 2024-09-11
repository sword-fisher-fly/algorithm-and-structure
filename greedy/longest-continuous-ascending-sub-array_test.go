package greedy

import "testing"

func TestLongestContinuousAscendingSubSequence(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10 items non-continuous ascending",
			args: args{
				arr: []int{1, 3, 5, 4, 7, 8, 6, 10, 9, 15},
			},
			want: 3,
		},
		{
			name: "10 items continuous ascending",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestContinuousAscendingSubSequence(tt.args.arr); got != tt.want {
				t.Errorf("LongestContinuousAscendingSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
