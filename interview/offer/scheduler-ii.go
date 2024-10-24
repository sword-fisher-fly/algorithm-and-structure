package offer

// TODO:
// https://leetcode.cn/problems/course-schedule-ii/

func FindOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	nextCourseMap := make(map[int][]int)
	// prerequisites[i][0]: 出度  prerequisites[i][1]: 入度
	for i := range prerequisites {
		inDegree[prerequisites[i][0]]++
		nextCourseMap[prerequisites[i][1]] = append(nextCourseMap[prerequisites[i][1]], prerequisites[i][0])
	}

	queue := []int{}
	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	ans := []int{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		ans = append(ans, cur)
		for _, n := range nextCourseMap[cur] {
			inDegree[n]--
			if inDegree[n] == 0 {
				queue = append(queue, n)
			}
		}
	}

	for i := range inDegree {
		if inDegree[i] != 0 {
			return []int{}
		}
	}

	return ans
}
