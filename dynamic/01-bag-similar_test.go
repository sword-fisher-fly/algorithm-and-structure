package dynamic

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
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
			name: "standard example",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTargetSumWaysByBacktracing(t *testing.T) {
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
			name: "standard example",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTargetSumWaysByBacktracing(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindTargetSumWaysByBacktracing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChange(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				amount: 10,
				coins:  []int{1, 10},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChangeWithMinCount(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				amount: 10,
				coins:  []int{1, 10},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				amount: 11,
				coins:  []int{1, 2, 5},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				amount: 2,
				coins:  []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChangeWithMinCount(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinChangeWithMinCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinChangeWithMinCountII(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				amount: 5,
				coins:  []int{1, 2, 5},
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				amount: 10,
				coins:  []int{1, 10},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				amount: 11,
				coins:  []int{1, 2, 5},
			},
			want: 3,
		},
		{
			name: "case 4",
			args: args{
				amount: 2,
				coins:  []int{1},
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				amount: 1,
				coins:  []int{2, 5},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChangeWithMinCountII(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinChangeWithMinCountII() = %v, want %v", got, tt.want)
			}
		})
	}
}
