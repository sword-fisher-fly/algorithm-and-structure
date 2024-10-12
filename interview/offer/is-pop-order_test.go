package offer

import "testing"

func TestIsPopOrder(t *testing.T) {
	type args struct {
		pushV []int
		popV  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is pop order",
			args: args{
				pushV: []int{1, 2, 3, 4, 5},
				popV:  []int{4, 5, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "no pop order",
			args: args{
				pushV: []int{1, 2, 3, 4, 5},
				popV:  []int{4, 3, 5, 1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPopOrder(tt.args.pushV, tt.args.popV); got != tt.want {
				t.Errorf("IsPopOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
