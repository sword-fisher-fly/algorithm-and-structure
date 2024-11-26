package array

import (
	"testing"
)

func TestFindRepeatedElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one repeated element in a array",
			args: args{nums: []int{2, 1, 3, 2, 5, 4, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatedElement(tt.args.nums); got != tt.want {
				t.Errorf("FindRepeatedElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindRepeatedElementII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one repeated element in a array",
			args: args{nums: []int{2, 1, 3, 2, 5, 4, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindRepeatedElementII(tt.args.nums); got != tt.want {
				t.Errorf("FindRepeatedElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}
