package str

import "testing"

func TestFindFirstCharAppearOnlyOnce(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "case 1",
			args: args{s: "abcabced"},
			want: 'e',
		},
		{
			name: "case 2",
			args: args{s: "abc"},
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstCharAppearOnlyOnce(tt.args.s); got != tt.want {
				t.Errorf("FindFirstCharAppearOnlyOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
