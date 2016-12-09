# A-Maze-ingly Retro Route Puzzle #

This program can output a valid route one could follow to collect all specified items within a map. The map is a json description of set of rooms with allowed path and contained object. The user should provide a valid map file, a valid starting room and the list of objects to collect.

The room description is composed as follows:

```
Room type allowed fields
  id: Integer
  name: String
  north: Integer //referring to a connected room
  south: Integer //referring to a connected room
  west: Integer  //referring to a connected room
  east: Integer  //referring to a connected room
  objects: List  //of Objects
Object type allowed fields
  name: String //Object Name
```

This is an example of a map:
```
{
"rooms":
  [
    { "id": 1, "name": "Hallway", "north": 2, objects: [] },
    { "id": 2, "name": "Dining Room", "south": 1, "west": 3, "east": 4, objects: [] },
    { "id": 3, "name": "Kitchen","east":2, objects: [ { "name": "Knife" } ] },
    { "id": 4, "name": "Sun Room","west":2, objects: [ { "name": "Potted Plant" } ] }
  ]
}
```

### Example ###
assuming we are using the example map, the input
```
maze-puzzle testMap.json 2 "Knife" "Potted Plant"
```
will produce the following output:
```
ID	Room		Object Collected
----------------------------
2	Dining Room	None
4	Sun Room	Potted Plant	
2	Dining Room	None
3	Kitchen		Knife
```

### Assumptions ###
I made the following assumptions in the development of this software:
- the map file is valid and correct;
- the objects are unique (i.e., the same object cannot be in more than one room);
- the id of the rooms starts from 1;

### Run the software ###
This software can be run through a Docker container: [`elleflorio/mazepuzzle`](https://hub.docker.com/r/elleflorio/mazepuzzle/).

You can attach a volume containing the map file to use and then run the software with a specific input, like this:
```
docker run --rm -v /host/path/to/maps:/tmp elleflorio/mazepuzzle /tmp/testMaze.json 2 "Knife" "Potted Plant"
```
this command will create (and remove at the end) a Docker container attaching the host folder containing the maps to the `\tmp` folder inside the container, then it will run the software using the input `/tmp/testMaze.json 2 "Knife" "Potted Plant"`.

If you prefer to run it without a Docker container, simply `go get` this repo:
```
go get github.com/elleFlorio/maze-puzzle
```
the software does not have external dependencies, except one used for testing. If you want to run the tests, `go get` the following package:
```
go get github.com/stretchr/testify/assert
```

### Enjoy! ###
This should be all, enjoy! :-)
