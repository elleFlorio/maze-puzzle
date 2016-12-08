package route

type node struct {
	id     int
	parent *node
}

// Breadth First Search algorithm with backtracking
func Bfs(start int, end int) []int {
	if adjacencyMap == nil {
		return []int{}
	}

	queue := make([]node, 0, len(adjacencyMap))
	visited := make([]node, 0, len(adjacencyMap))

	root := node{
		id:     end,
		parent: nil,
	}

	queue = enqueue(root, queue)
	isQueueEmpty := false
	for !isQueueEmpty {
		var current node
		current, queue = dequeue(queue)
		visited = append(visited, current)
		if current.id == start {
			return createBackPath(current)
		}

		childs := adjacencyMap[current.id]
		for _, child := range childs {
			chNode := node{
				id:     child,
				parent: &current,
			}

			if !contains(chNode, visited) {
				queue = enqueue(chNode, queue)
			}
		}

		if len(queue) < 1 {
			isQueueEmpty = true
		}
	}

	return []int{}
}

func enqueue(n node, q []node) []node {
	return append(q, n)
}

func dequeue(q []node) (node, []node) {
	if len(q) < 1 {
		n := node{
			id:     -1,
			parent: nil,
		}
		return n, q
	}

	n := q[0]
	q = q[1:]

	return n, q
}

func createBackPath(n node) []int {
	path := []int{n.id}
	next := n.parent
	for next != nil {
		path = append(path, next.id)
		next = next.parent
	}

	return path
}

func contains(n node, l []node) bool {
	for _, ln := range l {
		if n.id == ln.id {
			return true
		}
	}

	return false
}
