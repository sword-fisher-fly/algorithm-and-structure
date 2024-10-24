package offer

import "testing"

func TestLenLongestFibSubseq(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,1,2,3,5,8,13,21",
			args: args{[]int{1, 1, 2, 3, 5, 8, 13, 21}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LenLongestFibSubseq(tt.args.arr); got != tt.want {
				t.Errorf("LenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
