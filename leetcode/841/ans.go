package main

// 8ms
// 4.2MB
func canVisitAllRoomsDfs(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	count := 0
	var open func(int)

	open = func(i int) {
		count++
		visited[i] = true
		for _, next := range rooms[i] {
			if visited[next] {
				continue
			}

			open(next)
		}
	}

	open(0)

	return count == n
}

// 8ms
// 4.1MB
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)

	queue := make([]int, 0, n)
	queue = append(queue, 0)
	visited := make([]bool, n)
	visited[0] = true
	count := 1

	for len(queue) > 0 && count < n {
		r := queue[0]
		for _, next := range rooms[r] {
			if visited[next] {
				continue
			}
			visited[next] = true
			count++
			queue = append(queue, next)
		}
		queue = queue[1:]
	}

	return count == n
}
