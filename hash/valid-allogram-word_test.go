package hash

import "testing"

func TestIsValidAllogramWord(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid allogram word",
			args: args{
				s: "ate",
				t: "eat",
			},
			want: true,
		},
		{
			name: "invalid diffent length",
			args: args{
				s: "ate",
				t: "eats",
			},
			want: false,
		},
		{
			name: "invalid different characters",
			args: args{
				s: "ate",
				t: "eats",
			},
			want: false,
		},
		{
			name: "invalid allogram word with different characters",
			args: args{
				s: "word",
				t: "ordv",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidAllogramWord(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsValidAllogramWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
