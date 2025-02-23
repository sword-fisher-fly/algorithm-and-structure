package special

import "strings"

// https://leetcode.com/problems/swap-for-longest-repeated-character-substring/

// Given a string text, we are allowed to swap two of the characters in the string. Find the length of the longest substring with repeated characters.

// Example 1:

// ```text
// Input: text = "ababa"
// Output: 3
// Explanation: We can swap the first 'b' with the last 'a', or the last 'b' with the first 'a'. Then, the longest repeated character substring is "aaa", which its length is 3.
// ```
func MaxRepOpt1(text string) int {
	rec := [26][][3]int{}
	segments, index := split(text), 0
	for _, s := range segments {
		b, n := int(s[0]-'a'), len(s)
		rec[b] = append(rec[b], [3]int{index, index + n + 1, n})
		index += n
	}

	res := 0
	for _, r := range rec {
		n := len(r)
		if n == 0 {
			continue
		}
		// extension a lonely segment
		ext := 0
		if n > 1 {
			ext = 1
		}
		// connect two neighbor segments
		con := 0
		if n > 2 {
			con = 1
		}
		prev := r[0]
		res = max(res, ext+prev[2])
		for i := 1; i < len(r); i++ {
			cur := r[i]
			if prev[1] == cur[0] {
				res = max(res, con+prev[2]+cur[2])
			} else {
				res = max(res, ext+cur[2])
			}
			prev = cur
		}
	}

	return res
}

func split(s string) []string {
	var sb strings.Builder
	p := s[0]
	sb.WriteByte(p)
	for i := 1; i < len(s); i++ {
		n := s[i]
		if p != n {
			sb.WriteByte('\n')
		}
		sb.WriteByte(n)
		p = n
	}
	return strings.Split(sb.String(), "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
