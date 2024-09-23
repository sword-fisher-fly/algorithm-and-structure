package dynamic

import "testing"

func TestCanPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "can partition subset",
			args: args{nums: []int{1, 5, 11, 5}},
			want: true,
		},
		{
			name: "can not partition subset",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPartition(tt.args.nums); got != tt.want {
				t.Errorf("CanPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
