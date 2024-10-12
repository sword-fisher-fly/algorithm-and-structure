package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numInput := scanner.Text()

	fields := strings.Fields(numInput)

	nums := make([]int, len(fields))

	for i := range fields {
		nums[i], _ = strconv.Atoi(fields[i])
	}

	for i := 0; i < q; i++ {
		startIndex, endIndex := 0, 0
		fmt.Scan(&startIndex, &endIndex)
		tempNum := make([]int, endIndex-startIndex+1)
		copy(tempNum, nums[startIndex-1:endIndex])

		sort.Ints(tempNum)
		fmt.Println(tempNum[len(tempNum)/2])
	}
}
