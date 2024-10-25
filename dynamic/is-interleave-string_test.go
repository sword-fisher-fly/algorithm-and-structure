package dynamic

import "testing"

func TestIsInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("IsInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
