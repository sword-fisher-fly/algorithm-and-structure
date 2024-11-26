package huawei

import "testing"

func TestCountSameChars(t *testing.T) {
	type args struct {
		s string
		t byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{s: "8 8 8  8A i i OOI              IIIaa", t: 'A'},
			want: 3,
		},
		{
			name: "case 2",
			args: args{s: "ABCabc", t: 'A'},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSameChars(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("CountSameChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
