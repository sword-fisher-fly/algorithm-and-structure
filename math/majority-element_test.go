package math

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3 items with majority",
			args: args{nums: []int{1, 2, 1}},
			want: 1,
		},
		{
			name: "10 items with majority",
			args: args{nums: []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
