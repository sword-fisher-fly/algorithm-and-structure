package dynamic

import "testing"

func TestMaxCommonSubArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "same array",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "different array",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{5, 4, 1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxCommonSubArray(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("MaxCommonSubArray() = %v, want %v", got, tt.want)
			}

			if got := MaxCommonSubArrayII(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("MaxCommonSubArrayII() = %v, want %v", got, tt.want)
			}
		})
	}
}
