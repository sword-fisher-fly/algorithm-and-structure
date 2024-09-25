package graph

// 用二维数组graph表示有向无环图
// graph[i]: 表示有向图i号节点所能到达的下一个节点, 如graph[0]=[1,2]表示0->1,0->2两个有向边
// 求: 0-n-1的路径

// Input: graph = [[1,2],[3],[3],[]]
// Output: [[0,1,3], [0,2,3]]
func FindAllGraphPath(graph [][]int) [][]int {
	res := [][]int{}

	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		// fmt.Printf("i=%d\n", i)
		if i == len(graph)-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for _, j := range graph[i] {
			path = append(path, j)
			dfs(j, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{0})

	return res
}
