package game

type Room struct {
	Name    string
	Item    int
	Monster string
	North   string
	South   string
	East    string
	West    string
}

func GetMap() map[string]Room {
	return map[string]Room{
		"bathroom": {
			Name:  "Bathroom",
			Item:  1,
			North: "wall",
			South: "wall",
			East:  "exit",
			West:  "bedroom-1",
		},
		"bedroom-1": {
			Name:  "Bedroom 1",
			Item:  2,
			North: "gameroom",
			South: "wall",
			East:  "bathroom",
			West:  "wall",
		},
		"gameroom": {
			Name:    "Game Room",
			Item:    3,
			Monster: "slime",
			North:   "baseball_room",
			South:   "bedroom-1",
			East:    "wall",
			West:    "wall",
		},
		"baseball_room": {
			Name:  "Baseball Room",
			Item:  4,
			North: "wall",
			South: "gameroom",
			East:  "wall",
			West:  "locked",
		},
	}
}
