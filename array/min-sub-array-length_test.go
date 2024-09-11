package array

import "testing"

func TestMinSubArrayLength(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr:    []int{2, 3, 1, 2, 4, 3},
				target: 7,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				arr:    []int{2, 3, 7, 2, 4, 3},
				target: 7,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubArrayLength(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("MinSubArrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
