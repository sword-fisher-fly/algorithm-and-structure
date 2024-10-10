package monotonic

import (
	"testing"
)

func TestTrapByForce(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 items",
			args: args{heights: []int{0, 1}},
			want: 0,
		},
		{
			name: "3 items",
			args: args{heights: []int{1, 0, 2}},
			want: 1,
		},
		{
			name: "0,1,0,2,1,0,1,3,2,1,2,1",
			args: args{heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "4,2,0,3,2,5",
			args: args{heights: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapByForce(tt.args.heights); got != tt.want {
				t.Errorf("TrapByForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrapByTwoPointer(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 items",
			args: args{heights: []int{0, 1}},
			want: 0,
		},
		{
			name: "3 items",
			args: args{heights: []int{1, 0, 2}},
			want: 1,
		},
		{
			name: "0,1,0,2,1,0,1,3,2,1,2,1",
			args: args{heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "4,2,0,3,2,5",
			args: args{heights: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapByTwoPointer(tt.args.heights); got != tt.want {
				t.Errorf("TrapByTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrapII(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 items",
			args: args{heights: []int{0, 1}},
			want: 0,
		},
		{
			name: "3 items",
			args: args{heights: []int{1, 0, 2}},
			want: 1,
		},
		{
			name: "0,1,0,2,1,0,1,3,2,1,2,1",
			args: args{heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "4,2,0,3,2,5",
			args: args{heights: []int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrapII(tt.args.heights); got != tt.want {
				t.Errorf("TrapII() = %v, want %v", got, tt.want)
			}
		})
	}
}
