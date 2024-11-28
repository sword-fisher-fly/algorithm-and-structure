package str

func FindFirstCharAppearOnlyOnce(s string) byte {
	cMap := make(map[byte]bool)
	hasSame := false
	for i := 0; i < len(s); i++ {
		if cMap[s[i]] {
			continue
		}
		j := i + 1
		for j < len(s) {
			if s[i] == s[j] {
				hasSame = true
				break
			}
			j++
		}

		if !hasSame || i == len(s)-1 {
			return s[i]
		}
		hasSame = false
		cMap[s[i]] = true
	}

	return 0
}
