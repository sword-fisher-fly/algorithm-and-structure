package special

import (
	"strings"
)

func IsValidIP(IP string) string {
	isDigit := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	isHex := func(c byte) bool {
		return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
	}

	ipv4 := strings.Split(IP, ".")
	if len(ipv4) == 4 {
		for i := 0; i < len(ipv4); i++ {
			if len(ipv4[i]) == 0 {
				return "Neither"
			}

			if ipv4[i][0] == '0' && len(ipv4[i]) > 1 {
				return "Neither"
			}

			num := 0
			for j := 0; j < len(ipv4[i]); j++ {
				if !isDigit(ipv4[i][j]) {
					return "Neither"
				}

				num = num*10 + int(ipv4[i][j]-'0')
			}

			if num < 0 || num > 255 {
				return "Neither"
			}
		}

		return "IPv4"
	}

	ipv6 := strings.Split(IP, ":")
	if len(ipv6) == 8 {
		for i := 0; i < len(ipv6); i++ {
			if len(ipv6[i]) == 0 {
				return "Neither"
			}

			for j := 0; j < len(ipv6[i]); j++ {
				if len(ipv6[i]) == 1 && ipv6[i][0] == '0' {
					break
				}

				if len(ipv6[i]) != 4 {
					return "Neither"
				}

				if !isHex(ipv6[i][j]) {
					return "Neither"
				}
			}
		}

		return "IPv6"
	}

	return "Neither"
}
