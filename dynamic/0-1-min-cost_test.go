package dynamic

import (
	"testing"
)

func TestMinMoney(t *testing.T) {
	type args struct {
		arr []int
		aim int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{5, 3, 2},
				aim: 20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMoney(tt.args.arr, tt.args.aim); got != tt.want {
				t.Errorf("MinMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMoneyByRecursive(t *testing.T) {
	type args struct {
		arr []int
		aim int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				arr: []int{5, 3, 2},
				aim: 20,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMoneyByRecursive(tt.args.arr, tt.args.aim); got != tt.want {
				t.Errorf("MinMoneyByRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
