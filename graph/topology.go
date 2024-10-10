package graph

// N个文件之间拥有M条依赖关系
// 转化为有向图, 把这个有向图转成线性的排序, 称为拓扑排序
// 前提有向无环图
// 解题思路: 迭代查找入度为0的节点, 找到后删除, 继续查找, 直到没有节点可以删除
/*
0 1
0 2
1 3
2 4
*/

func TopologicalSort(n int64, edges [][2]int64) []int64 {
	umap := make(map[int64][]int64)
	inDegree := make(map[int64]int)

	for i := int64(0); i < n; i++ {
		inDegree[i] = 0
	}

	// s -> t 表示节点s指向t
	for _, edge := range edges {
		inDegree[edge[1]]++
		umap[edge[0]] = append(umap[edge[0]], edge[1])
	}

	queue := []int64{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	result := []int64{}
	for len(queue) > 0 {
		curVertex := queue[0]
		queue = queue[1:]
		result = append(result, curVertex)
		curNexts := umap[curVertex]
		for i := 0; i < len(curNexts); i++ {
			inDegree[curNexts[i]]--
			if inDegree[curNexts[i]] == 0 {
				queue = append(queue, curNexts[i])
			}
		}
	}

	return result
}

func TopologicalSortFromGraphList(graph *GraphList) []int64 {
	umap := make(map[int64][]int64)
	inDegree := make(map[int64]int)

	for k := range graph.Vertex {
		inDegree[k] = 0
	}

	for s, t := range graph.Edges {
		umap[s] = append(umap[s], t...)
	}

	for _, vv := range umap {
		for _, v := range vv {
			inDegree[v]++
		}
	}

	queue := []int64{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	result := []int64{}
	for len(queue) > 0 {
		curVertex := queue[0]
		queue = queue[1:]
		result = append(result, curVertex)
		curNexts := umap[curVertex]
		// if len(curNexts) > 0 {
		for i := 0; i < len(curNexts); i++ {
			inDegree[curNexts[i]]--
			if inDegree[curNexts[i]] == 0 {
				queue = append(queue, curNexts[i])
			}
		}
		// }
	}

	return result
}
