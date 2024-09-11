package str

import "testing"

// "123", "123.456", "-123.456", "abc", "123abc", "123.456.789", "1e10", "-1E-10", "+1.e3", ".1e1", "1e", "e3", "--1.1", "1.1.1", "1e1.1"
// 有问题??
func TestIsValidScientific(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0.0",
			args: args{s: "0.0"},
			want: true,
		},
		{
			name: "0.0E16",
			args: args{s: "0.0"},
			want: true,
		},
		{
			name: "0.0E-16",
			args: args{s: "0.0"},
			want: true,
		},
		{
			name: "-0.0E-16",
			args: args{s: "0.0"},
			want: true,
		},
		{
			name: "123",
			args: args{s: "123"},
			want: true,
		},
		{
			name: "123.+123",
			args: args{s: "123.+123"},
			want: false,
		},
		{
			name: "123.456",
			args: args{s: "123.456"},
			want: true,
		},
		{
			name: "-123.456",
			args: args{s: "-123.456"},
			want: true,
		},
		{
			name: "abc",
			args: args{s: "abc"},
			want: false,
		},
		{
			name: "123abc",
			args: args{s: "123abc"},
			want: false,
		},

		{
			name: "123.456.789",
			args: args{s: "123.456.789"},
			want: false,
		},
		{
			name: "1e10",
			args: args{s: "1e10"},
			want: true,
		},
		{
			name: "-1E-10",
			args: args{s: "-1E-10"},
			want: true,
		},
		{
			name: "+1.e3",
			args: args{s: "+1.e3"},
			want: false,
		},
		//".1e1", "1e", "e3", "--1.1", "1.1.1", "1e1.1"
		{
			name: ".1e1",
			args: args{s: ".1e1"},
			want: false,
		},
		{
			name: "1e",
			args: args{s: "1e"},
			want: false,
		},
		{
			name: "e3",
			args: args{s: "e3"},
			want: false,
		},
		{
			name: "--1.1",
			args: args{s: "--1.1"},
			want: false,
		},
		{
			name: "1.1.1",
			args: args{s: "1.1.1"},
			want: false,
		},
		{
			name: "1e1.1",
			args: args{s: "1e1.1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidScientific(tt.args.s); got != tt.want {
				t.Errorf("IsValidScientific(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
