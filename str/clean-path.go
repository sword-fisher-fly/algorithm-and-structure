package str

import (
	"strings"
)

// https://leetcode.cn/problems/simplify-path/solutions/1193258/jian-hua-lu-jing-by-leetcode-solution-aucq/?envType=study-plan-v2&envId=top-interview-150
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

func SimplifyPathII(path string) string {
	res := []string{}

	part := []byte{}
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if len(part) > 0 && string(part) != "." && string(part) != ".." {
				res = append(res, string(part))
			}

			if len(res) > 0 && string(part) == ".." {
				res = res[:len(res)-1]
			}

			part = part[:0]
			continue
		}

		part = append(part, path[i])
	}

	if len(part) > 0 && string(part) != "." && string(part) != ".." {
		res = append(res, string(part))
	}

	if len(res) > 0 && string(part) == ".." {
		res = res[:len(res)-1]
	}

	return "/" + strings.Join(res, "/")
}
