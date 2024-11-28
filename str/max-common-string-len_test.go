package str

import "testing"

func TestFindMaxCommonSubStringLength(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s1: "asdfas",
				s2: "werasdfaswer",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxCommonSubStringLength(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("FindMaxCommonSubStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
