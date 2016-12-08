package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitializeRouteMaps(t *testing.T) {
	InitializeRouteMap(4)
	assert.Len(t, routeMap, 5)

	r := route{
		path: nil,
		cost: 0,
	}
	assert.Equal(t, r, routeMap[0][0])
}

func TestReverse(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}
	r := reverse(l)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, r)

	l = []int{}
	r = reverse(l)
	assert.Empty(t, r)

	l = []int{1}
	r = reverse(l)
	assert.Equal(t, []int{1}, r)
}

func TestComputeRoutes(t *testing.T) {
	defer clearAdjMap()
	defer clearRoutesMap()

	createAdjMap()
	initTestRoutesMap()

	nodes := []int{3, 4}
	ComputeRoutes(nodes)

	r34 := route{
		path: []int{3, 2, 4},
		cost: 3,
	}

	r43 := route{
		path: []int{4, 2, 3},
		cost: 3,
	}

	assert.Equal(t, r34, routeMap[3][4])
	assert.Equal(t, r43, routeMap[4][3])

	nodes = []int{}
	ComputeRoutes(nodes)

	nodes = []int{3}
	ComputeRoutes(nodes)

	nodes = []int{10000, 3}
	ComputeRoutes(nodes)
}

func TestGetPath(t *testing.T) {
	defer clearRoutesMap()
	defer clearAdjMap()

	createTestRoutesMap()
	createAdjMap()

	path := GetPath(2, []int{3, 4})
	assert.Equal(t, []int{2, 3, 2, 4}, path)

	path = GetPath(2, []int{})
	assert.Equal(t, []int{2}, path)

	path = GetPath(10000, []int{3, 4})
	assert.Empty(t, path)

	path = GetPath(3, []int{3, 4})
	assert.Equal(t, []int{3, 2, 4}, path)
}

func createAdjMap() {
	adjacencyMap = map[int][]int{
		1: []int{2},
		2: []int{1, 3, 4},
		3: []int{2},
		4: []int{2},
	}
}

func clearAdjMap() {
	adjacencyMap = nil
}

func initTestRoutesMap() {
	routeMap = make([][]route, 5, 5)
	for i := 0; i < 5; i++ {
		routeMap[i] = make([]route, 5, 5)
	}
}

func clearRoutesMap() {
	routeMap = nil
}

func createTestRoutesMap() {
	r23 := route{
		path: []int{2, 3},
		cost: 2,
	}

	r32 := route{
		path: []int{3, 2},
		cost: 2,
	}

	r24 := route{
		path: []int{2, 4},
		cost: 2,
	}

	r42 := route{
		path: []int{4, 2},
		cost: 2,
	}

	r34 := route{
		path: []int{3, 2, 4},
		cost: 3,
	}

	r43 := route{
		path: []int{4, 2, 3},
		cost: 3,
	}

	initTestRoutesMap()
	routeMap[2][3] = r23
	routeMap[3][2] = r32
	routeMap[2][4] = r24
	routeMap[4][2] = r42
	routeMap[3][4] = r34
	routeMap[4][3] = r43
}
