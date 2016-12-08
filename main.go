package main

import (
	"log"
	"os"
	"strconv"

	"github.com/elleFlorio/maze-puzzle/maze"
	"github.com/elleFlorio/maze-puzzle/route"
)

var (
	file    string
	start   int
	targets []string
)

func main() {
	// Read and prepare the arguments
	readArgs(os.Args[1:])

	// Read the maze and create all the maps
	maze.ReadMaze(file)
	// Get the rooms I want to reach
	// (i.e., the ones containing the objects to collect)
	targetRooms := maze.GetObjectsRooms(targets)
	// Add the starting room to the list of rooms to reach
	rooms := append([]int{start}, targetRooms...)

	adj := maze.GetAdjacencyMap()
	route.SetAdjacencyMap(adj)
	route.InitializeRouteMap(len(adj))
	// Compute the routes connecting all the rooms I have to reach
	route.ComputeRoutes(rooms)
	// Get the path to collect all the objects I need
	path := route.GetPath(start, targetRooms)

	// Print the path
	maze.PrintRoomsPath(path, targets)
}

func readArgs(args []string) {
	if len(args) < 3 {
		log.Fatalln("Not enough input arguments")
	}

	file = args[0]

	var err error
	start, err = strconv.Atoi(args[1])
	if err != nil || start < 1 {
		log.Fatalln("Invalid starting room")
	}

	targets = args[2:]
}
