package str

import "testing"

func Test_isValidInterger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive sign interger",
			args: args{s: "+123"},
			want: true,
		},
		{
			name: "positive sign interger with leading zero",
			args: args{s: "+021"},
			want: false,
		},
		{
			name: "positive sign interger with only leading zero",
			args: args{s: "+0"},
			want: true,
		},
		{
			name: "negative sign interger",
			args: args{s: "-123"},
			want: true,
		},
		{
			name: "negative sign interger with leading zero",
			args: args{s: "-021"},
			want: false,
		},
		{
			name: "negative sign interger with only leading zero",
			args: args{s: "-0"},
			want: true,
		},
		{
			name: "interger without sign",
			args: args{s: "123"},
			want: true,
		},
		{
			name: "interger without sign and with leading zero",
			args: args{s: "021"},
			want: false,
		},
		{
			name: "interger without sign and with only leading zero",
			args: args{s: "0"},
			want: true,
		},
		{
			name: "illegal char with dot",
			args: args{s: "1.1"},
			want: false,
		},
		{
			name: "illegal char with E",
			args: args{s: "1.1E"},
			want: false,
		},
		{
			name: "illegal char with E and dot",
			args: args{s: "1.1E3"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidInterger(tt.args.s); got != tt.want {
				t.Errorf("isNonZeroInterger(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func Test_isNonZeroInterger(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "positive sign interger",
			args: args{s: "+123"},
			want: true,
		},
		{
			name: "positive sign interger with leading zero",
			args: args{s: "+021"},
			want: false,
		},
		{
			name: "positive sign interger with only leading zero",
			args: args{s: "+0"},
			want: false,
		},
		{
			name: "negative sign interger",
			args: args{s: "-123"},
			want: true,
		},
		{
			name: "negative sign interger with leading zero",
			args: args{s: "-021"},
			want: false,
		},
		{
			name: "negative sign interger with only leading zero",
			args: args{s: "-0"},
			want: false,
		},
		{
			name: "interger without sign",
			args: args{s: "123"},
			want: true,
		},
		{
			name: "interger without sign and with leading zero",
			args: args{s: "021"},
			want: false,
		},
		{
			name: "interger without sign and with only leading zero",
			args: args{s: "0"},
			want: false,
		},
		{
			name: "illegal char with dot",
			args: args{s: "1.1"},
			want: false,
		},
		{
			name: "illegal char with E",
			args: args{s: "1.1E"},
			want: false,
		},
		{
			name: "illegal char with E and dot",
			args: args{s: "1.1E3"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNonZeroInterger(tt.args.s); got != tt.want {
				t.Errorf("isNonZeroInterger(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
