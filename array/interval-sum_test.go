package array

import "testing"

func TestIntervalSum(t *testing.T) {
	type args struct {
		arr []int
		i   int
		j   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0 item",
			args: args{
				arr: []int{},
				i:   0,
				j:   0,
			},
			want: -1,
		},
		{
			name: "1 item",
			args: args{
				arr: []int{1},
				i:   0,
				j:   1,
			},
			want: 1,
		},
		{
			name: "10 items",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				i:   3,
				j:   10,
			},
			want: 49,
		},
		{
			name: "i is equal to j",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				i:   3,
				j:   3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntervalSum(tt.args.arr, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("IntervalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
