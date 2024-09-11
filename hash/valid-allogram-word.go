package hash

// 1) s = "ate", t = "tea", return true
// 2) s = "word", t = "ordv", return false
// 3) s = "word", t = "world", return false
func IsValidAllogramWord(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ret := [26]int{}
	for i := 0; i < len(s); i++ {
		ret[s[i]-'a']++
		ret[t[i]-'a']--
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != 0 {
			return false
		}
	}

	return true
}
