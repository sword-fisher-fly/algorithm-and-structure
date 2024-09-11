package array

import "math"

func MinSubArrayLength(arr []int, target int) int {
	sum := 0
	lowIndex := 0

	minLength := math.MaxInt32
	for upperIndex := 0; upperIndex < len(arr); upperIndex++ {
		sum += arr[upperIndex]

		for sum >= target {
			length := upperIndex - lowIndex + 1

			if length < minLength {
				minLength = length
			}

			sum -= arr[lowIndex]
			lowIndex++
		}
	}

	return minLength
}
