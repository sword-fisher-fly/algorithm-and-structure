package twopointer

import "testing"

func TestMinWindows(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s=wordgoodgoodgoodbestword,t=goodword",
			args: args{s: "wordgoodgoodgoodbestword", t: "goodword"},
			want: "wordgood",
		},
		{
			name: "s=wordgoodgoodgoodbestword,t=goodword",
			args: args{s: "wordgoodgoodgoodbestword", t: "goodbestword"},
			want: "dgoodbestwor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindows(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("MinWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}
