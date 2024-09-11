package backtracking

var letterMap = map[byte]string{
	0: "",
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinationsHelper(digits string, index int, path string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, path)
		return
	}
	// "23"
	digit := digits[index] - '0'
	letters := letterMap[digit]

	for i := 0; i < len(letters); i++ {
		path += string(letters[i])
		letterCombinationsHelper(digits, index+1, path, res)
		path = path[:len(path)-1]
	}
}
func LetterCombinations(digits string) []string {
	var path string
	res := []string{}
	letterCombinationsHelper(digits, 0, path, &res)

	return res
}
