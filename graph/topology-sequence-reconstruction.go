package graph

// 检查[1,n]是否为唯一的最短超序列
// 最短超序列: 长度最短的序列, 并且所有sequences[i]都是它的子序列
// 对于给定数组sequences,可能存在多个有效的超序列
// nums: [1,2,3]  sequences: [[1,2], [1,3]] Output: false  两种超序列[1,2,3]和[1,3,2]
// nums: [1,2,3]  sequences: [[1,2], [1,3], [2, 3]] Output: true
func SequenceReconstruction(n int, sequences [][]int) bool {
	indegree := make([]int, n+1)
	graph := make([][]int, n+1)
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			graph[seq[i-1]] = append(graph[seq[i-1]], seq[i])
			indegree[seq[i]]++
		}
	}

	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 1 {
		if len(queue) > 1 {
			return false
		}

		cur := queue[0]
		queue = queue[1:]

		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return true
}
