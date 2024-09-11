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

	// b := bytes.Buffer{}
	return "/" + strings.Join(stack, "/")
}

// 遇到..如何回退??
// func CleanPathII(path string) string {
// 	result := strings.Builder{}

// 	for i := 0; i < len(path); i++ {
// 		if path[i] == '/' {
// 			if result.Len() == 0 || result.Len() > 0 && result.String()[result.Len()-1] != '/' {
// 				result.WriteByte('/')
// 			}

// 			continue
// 		}

// 		if i+1 < len(path) && path[i:i+2] == ".." {
// 		}

// 	}
// }
