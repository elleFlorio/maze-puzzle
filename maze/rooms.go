package maze

type Maze struct {
	Rooms []Room `json:"rooms"`
}

type Room struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	North   int      `json:"north"`
	South   int      `json:"south"`
	West    int      `json:"west"`
	East    int      `json:"east"`
	Objects []Object `json:"objects"`
}

type Object struct {
	Name string `json:"name"`
}
