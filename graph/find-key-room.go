package graph

import "fmt"

func CanVisitAllRoomsByBFS(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	queue := []int{0}

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		for _, key := range rooms[room] {
			if visited[key] {
				continue
			}

			visited[key] = true
			queue = append(queue, key)
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

func CanVisitAllRoomsByDFS(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true

	var dfs func(room int)

	dfs = func(room int) {
		for _, key := range rooms[room] {
			if visited[key] {
				continue
			}

			visited[key] = true
			dfs(key)
		}
	}

	dfs(0)
	fmt.Printf("visited: %v\n", visited)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

func CanVisitAllRoomsByDFSII(rooms [][]int) bool {
	visited := make([]bool, len(rooms))

	var dfs func(room int)
	dfs = func(room int) {
		if visited[room] {
			return
		}

		visited[room] = true
		for _, key := range rooms[room] {
			dfs(key)
		}
	}

	dfs(0)

	fmt.Printf("visited: %v\n", visited)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
