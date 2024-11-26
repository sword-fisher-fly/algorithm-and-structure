package special

import "testing"

func TestIsValidIP(t *testing.T) {
	type args struct {
		IP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1 ipv4 valid",
			args: args{IP: "172.16.254.1"},
			want: "IPv4",
		},
		{
			name: "case2 ipv6 valid",
			args: args{IP: "2001:0db8:85a3:0:0:8A2E:0370:7334"},
			want: "IPv6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIP(tt.args.IP); got != tt.want {
				t.Errorf("IsValidIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
