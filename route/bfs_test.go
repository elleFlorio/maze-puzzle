package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := []node{}
	n := node{
		id:     1,
		parent: nil,
	}

	q = enqueue(n, q)
	assert.Contains(t, q, n)
}

func TestDeque(t *testing.T) {
	q := []node{}
	n, q := dequeue(q)
	assert.Equal(t, -1, n.id)
	assert.Empty(t, q)

	n1 := node{
		id:     1,
		parent: nil,
	}

	n2 := node{
		id:     2,
		parent: &n1,
	}
	q = []node{n1, n2}

	n, q = dequeue(q)
	assert.Equal(t, n1.id, n.id)
	assert.Len(t, q, 1)

	n, q = dequeue(q)
	assert.Equal(t, n2.id, n.id)
	assert.Len(t, q, 0)
}

func TestCreateBackPath(t *testing.T) {
	n1 := node{
		id:     1,
		parent: nil,
	}

	n2 := node{
		id:     2,
		parent: &n1,
	}

	n3 := node{
		id:     3,
		parent: &n2,
	}

	n4 := node{
		id:     4,
		parent: &n3,
	}

	current := n1
	path := createBackPath(current)
	assert.Equal(t, []int{1}, path)

	current = n4
	path = createBackPath(current)
	assert.Equal(t, []int{4, 3, 2, 1}, path)
}

func TestContains(t *testing.T) {
	n1 := node{
		id:     1,
		parent: nil,
	}

	n2 := node{
		id:     2,
		parent: &n1,
	}

	l := []node{n1}
	assert.True(t, contains(n1, l))
	assert.False(t, contains(n2, l))
}

func TestBfs(t *testing.T) {
	defer clearAdjMap()

	createAdjMap()
	path := Bfs(3, 3)
	assert.NotEmpty(t, path)
	assert.Equal(t, []int{3}, path)

	path = Bfs(3, 4)
	assert.NotEmpty(t, path)
	assert.Equal(t, []int{3, 2, 4}, path)
}
