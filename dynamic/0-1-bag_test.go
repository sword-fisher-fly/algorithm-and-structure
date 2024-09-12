package dynamic

import "testing"

func TestMaxValueForBag(t *testing.T) {
	type args struct {
		weights   []int
		values    []int
		bagWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard example",
			args: args{
				weights:   []int{1, 3, 4},
				values:    []int{15, 20, 30},
				bagWeight: 4,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxValueForBag(tt.args.weights, tt.args.values, tt.args.bagWeight); got != tt.want {
				t.Errorf("MaxValueForBag() = %v, want %v", got, tt.want)
			}

			if got := MaxValueForBagII(tt.args.weights, tt.args.values, tt.args.bagWeight); got != tt.want {
				t.Errorf("MaxValueForBagII() = %v, want %v", got, tt.want)
			}

			if got := MaxValueForBagIII(tt.args.weights, tt.args.values, tt.args.bagWeight); got != tt.want {
				t.Errorf("MaxValueForBagIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
