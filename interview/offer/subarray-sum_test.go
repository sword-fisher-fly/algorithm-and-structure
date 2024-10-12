package offer

import (
	"testing"
)

func TestSubArraySum(t *testing.T) {
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
			name: "nums=[1,1,2], target=2",
			args: args{
				nums:   []int{1, 1, 2},
				target: 2,
			},
			want: 2,
		},
		{
			name: "nums=[1,2,3], target=9",
			args: args{
				nums:   []int{1, 2, 3},
				target: 9,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArraySum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SubArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSunArraySumWithHashPrefixSum(t *testing.T) {
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
			name: "nums=[1,1,2], target=2",
			args: args{
				nums:   []int{1, 1, 2},
				target: 2,
			},
			want: 2,
		},
		{
			name: "nums=[1,2,3], target=9",
			args: args{
				nums:   []int{1, 2, 3},
				target: 9,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SunArraySumWithHashPrefixSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SunArraySumWithHashPrefixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
