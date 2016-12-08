package maze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

var (
	mazeMap      map[int]Room
	adjacencyMap map[int][]int
	objectsMap   map[string]int
)

func init() {
	mazeMap = make(map[int]Room)
	adjacencyMap = make(map[int][]int)
	objectsMap = make(map[string]int)
}

func ReadMaze(url string) {
	// Read map file
	raw, err := ioutil.ReadFile(url)
	if err != nil {
		log.Fatalln("Error reading map file")
	}

	var maze Maze

	// Unmarshal JSON
	err = json.Unmarshal(raw, &maze)
	if err != nil {
		log.Fatalln(err)
	}

	// Compute maps
	computeMaps(maze.Rooms)
}

func computeMaps(rooms []Room) {
	for _, room := range rooms {
		// Save room in maze map
		mazeMap[room.Id] = room

		// Get the neighbors of the room (if they exist)
		neighbors := []int{}
		if room.North > 0 {
			neighbors = append(neighbors, room.North)
		}
		if room.South > 0 {
			neighbors = append(neighbors, room.South)
		}
		if room.West > 0 {
			neighbors = append(neighbors, room.West)
		}
		if room.East > 0 {
			neighbors = append(neighbors, room.East)
		}

		// Save the adjacent rooms in the adjacency map
		adjacencyMap[room.Id] = neighbors

		// Save where objects are stored
		for _, obj := range room.Objects {
			objectsMap[obj.Name] = room.Id
		}
	}
}

func GetAdjacencyMap() map[int][]int {
	return adjacencyMap
}

// Get the rooms where are stored the objects
func GetObjectsRooms(objects []string) []int {
	nodes := make([]int, 0, len(objectsMap))
	for _, toCollect := range objects {
		if room, ok := objectsMap[toCollect]; ok {
			if !contains(room, nodes) {
				nodes = append(nodes, room)
			}
		}
	}

	return nodes
}

func contains(i int, l []int) bool {
	for _, ln := range l {
		if i == ln {
			return true
		}
	}

	return false
}

// Print a path of rooms with their object to collect
func PrintRoomsPath(path []int, toCollect []string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintf(w, "\nID\tRoom\t\tObject Collected\n")
	fmt.Fprintf(w, "----------------------------\n")
	for _, room := range path {
		fmt.Fprintf(w, "%d\t%s\t", mazeMap[room].Id, mazeMap[room].Name)
		objects := mazeMap[room].Objects
		if len(objects) < 1 {
			fmt.Fprintf(w, "None")
		} else {
			// Print only the ones to collect, not the others
			for _, obj := range objects {
				for _, collected := range toCollect {
					if obj.Name == collected {
						fmt.Fprintf(w, "%s\t", collected)
					}
				}
			}
		}
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "\n")

	w.Flush()
}
