package route

type route struct {
	path []int
	cost int
}

const c_MAX_DISTANCE = 100000000

var (
	routeMap     [][]route
	adjacencyMap map[int][]int
)

func SetAdjacencyMap(adj map[int][]int) {
	adjacencyMap = adj
}

// the size is increased by one because the id are not 0-based
func InitializeRouteMap(size int) {
	routeMap = make([][]route, size+1, size+1)
	for i := 0; i < size+1; i++ {
		routeMap[i] = make([]route, size+1, size+1)
	}
}

// Compute the root from all the nodes to all the nodes
func ComputeRoutes(nodes []int) {
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			n1 := nodes[i]
			n2 := nodes[j]
			// check if the values are valid
			// and if that path has already been computed
			if n1 <= len(adjacencyMap) &&
				n2 <= len(adjacencyMap) &&
				routeMap[n1][n2].path == nil {
				p := Bfs(n1, n2)
				c := len(p)
				n1n2 := route{
					path: p,
					cost: c,
				}

				n2n1 := route{
					path: reverse(p),
					cost: c,
				}

				routeMap[n1][n2] = n1n2
				routeMap[n2][n1] = n2n1
			}
		}
	}
}

// make a reversed copy of a list
func reverse(l []int) []int {
	r := make([]int, len(l))
	copy(r, l)
	for i := len(r)/2 - 1; i >= 0; i-- {
		opp := len(r) - 1 - i
		r[i], r[opp] = r[opp], r[i]
	}

	return r
}

// Get the path only if the start value is valid
func GetPath(start int, targets []int) []int {
	if start >= len(routeMap) {
		return []int{}
	}

	nodes := append([]int{start}, targets...)
	//return findPathOrder(nodes)

	return findPathGreedy(nodes, 0, []int{})

}

func findPathOrder(nodes []int) []int {
	path := []int{nodes[0]}
	for i := 0; i < len(nodes)-1; i++ {
		n1 := nodes[i]
		n2 := nodes[i+1]
		p := routeMap[n1][n2].path
		if len(p) < 2 {
			path = append(path, p...)
		} else {
			path = append(path, p[1:]...)
		}

	}

	return path
}

func findPathGreedy(nodes []int, index int, path []int) []int {
	distance := c_MAX_DISTANCE
	next := 0
	indexNext := 0

	current := nodes[index]
	nodes = append(nodes[0:index], nodes[index+1:]...)

	if len(nodes) < 1 {
		path = append(path, current)
		return path
	}

	for i, node := range nodes {
		if node == current {
			return findPathGreedy(nodes, i, path)
		}

		nextDist := routeMap[current][node].cost
		if nextDist < distance {
			distance = nextDist
			next = node
			indexNext = i
		}
	}

	toAdd := routeMap[current][next].path
	if len(toAdd) > 1 {
		path = append(path, toAdd[:len(toAdd)-1]...)
	}

	return findPathGreedy(nodes, indexNext, path)
}
