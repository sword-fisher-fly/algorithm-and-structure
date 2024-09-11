package hash

// 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0
/*
A = [1,2]
B = [-2,-1]
C = [-1,2]
D = [0,2]
*/
func FourSumCount(A []int64, B []int64, C []int64, D []int64) int {
	sumAB := make(map[int64]int)

	for _, a := range A {
		for _, b := range B {
			sum := a + b
			sumAB[sum]++
		}
	}

	var ret int

	for _, c := range C {
		for _, d := range D {
			sum := c + d
			if _, ok := sumAB[0-sum]; ok {
				sumAB[-sum]--
				if sumAB[-sum] == 0 {
					delete(sumAB, -sum)
				}
				ret++
			}
		}
	}

	return ret
}

func FourSumCountII(A []int64, B []int64, C []int64, D []int64) int {
	sumAB := make(map[int64]int)

	for _, a := range A {
		for _, b := range B {
			sum := a + b
			sumAB[sum]++
		}
	}

	var ret int

	for _, c := range C {
		for _, d := range D {
			sum := c + d
			if _, ok := sumAB[0-sum]; ok {
				ret += sumAB[-sum]
			}
		}
	}

	return ret
}
