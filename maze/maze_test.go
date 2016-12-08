package maze

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeMaps(t *testing.T) {
	defer clearTestMaps()

	rooms := createTestRooms()
	computeMaps(rooms)

	assert.Len(t, mazeMap, 4)
	assert.Len(t, objectsMap, 2)
	assert.Len(t, adjacencyMap, 4)

	for _, room := range rooms {
		assert.Equal(t, room, mazeMap[room.Id])
	}

	assert.Equal(t, 3, objectsMap["Knife"])
	assert.Equal(t, 4, objectsMap["Potted Plant"])

	assert.Equal(t, []int{2}, adjacencyMap[1])
	assert.Equal(t, []int{1, 3, 4}, adjacencyMap[2])
	assert.Equal(t, []int{2}, adjacencyMap[3])
	assert.Equal(t, []int{2}, adjacencyMap[4])
}

func TestReadMaze(t *testing.T) {
	defer removeTestMazeFile()

	createTestMazeFile()

	ReadMaze("testMaze.json")
	assert.NotNil(t, mazeMap)
	assert.NotEmpty(t, mazeMap)
	assert.Equal(t, mazeMap[1].Name, "Hallway")
	assert.Equal(t, mazeMap[3].Objects[0].Name, "Knife")
}

func createTestMazeFile() {
	testMaze := `{
	"rooms": 
	[
  		{ "id": 1, "name": "Hallway", "north": 2, "objects": [] },
  		{ "id": 2, "name": "Dining Room", "south": 1, "west": 3, "east": 4, "objects": []
		},
  		{ "id": 3, "name": "Kitchen","east":2, "objects": [ { "name": "Knife" } ] },
  		{ "id": 4, "name": "Sun Room","west":2, "objects": [ { "name": "Potted Plant" } ]
		}
	] 
	}`

	file, err := os.Create("testMaze.json")
	if err != nil {
		log.Fatalln("Error writing testMaze file")
	}
	defer file.Close()

	file.WriteString(testMaze)
}

func removeTestMazeFile() {
	os.Remove("testMaze.json")
}

func TestPrintRoomsPath(t *testing.T) {
	defer clearTestMaps()

	createTestMaze()
	PrintRoomsPath([]int{2, 3, 2, 4}, []string{"Knife", "Potted Plant"})
}

func TestGetObjectsRooms(t *testing.T) {
	defer clearTestMaps()

	createTestObjecsMap()

	nodes := GetObjectsRooms([]string{"Knife", "Potted Plant"})
	assert.Equal(t, []int{3, 4}, nodes)

	nodes = GetObjectsRooms([]string{"Knife", "Potted Plant", "Knife"})
	assert.Equal(t, []int{3, 4}, nodes)

	nodes = GetObjectsRooms([]string{})
	assert.Empty(t, nodes)

	nodes = GetObjectsRooms([]string{"Pippo"})
	assert.Empty(t, nodes)

	nodes = GetObjectsRooms([]string{"Knife", "Pippo"})
	assert.Equal(t, []int{3}, nodes)

	nodes = GetObjectsRooms([]string{"Pippo", "Potted Plant"})
	assert.Equal(t, []int{4}, nodes)
}

func createTestMaze() {
	rooms := createTestRooms()

	for _, room := range rooms {
		mazeMap[room.Id] = room
	}
}

func createTestObjecsMap() {
	objectsMap["Knife"] = 3
	objectsMap["Potted Plant"] = 4
}

func createTestRooms() []Room {
	r1 := Room{
		Id:    1,
		Name:  "Hallway",
		North: 2,
	}

	r2 := Room{
		Id:    2,
		Name:  "Dining Room",
		South: 1,
		West:  3,
		East:  4,
	}

	r3 := Room{
		Id:   3,
		Name: "Kitchen",
		East: 2,
		Objects: []Object{
			Object{
				Name: "Knife",
			},
		},
	}

	r4 := Room{
		Id:   4,
		Name: "Sun Room",
		West: 2,
		Objects: []Object{
			Object{
				Name: "Potted Plant",
			},
		},
	}

	return []Room{r1, r2, r3, r4}
}

func clearTestMaps() {
	mazeMap = make(map[int]Room)
	adjacencyMap = make(map[int][]int)
	objectsMap = make(map[string]int)
}
