package str

import "testing"

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a=101,b=1011",
			args: args{a: "101", b: "1011"},
			want: "10000",
		},
		{
			name: "a=1,b=1",
			args: args{a: "1", b: "1"},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
