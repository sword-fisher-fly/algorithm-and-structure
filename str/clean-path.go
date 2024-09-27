package str

import (
	"strings"
)

// illegal path:
// 1) /a/b/.../c
// 2) a/b/./.././c
func CleanPath(path string) string {
	stack := []string{}

	part := ""
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if part != "" {
				stack = append(stack, part)
			}

			part = ""
			continue
		}

		// ..
		if i+1 < len(path) && path[i] == '.' && path[i+1] == '.' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

			i++
			continue
		}

		// .
		if path[i] == '.' {
			continue
		}

		part += string(path[i])
	}

	return "/" + strings.Join(stack, "/")
}
