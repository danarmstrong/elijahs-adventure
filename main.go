package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Room struct {
	Name  string
	Item  int
	North string
	South string
	East  string
	West  string
}

func main() {

	items := []string{
		"Nothing",
		"Towel",
		"Bed",
		"TV",
		"Baseball Bat",
	}

	house := map[string]Room{
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
			Name:  "Game Room",
			Item:  3,
			North: "baseball_room",
			South: "bedroom-1",
			East:  "wall",
			West:  "wall",
		},
		"baseball_room": {
			Name:  "Baseball Room",
			Item:  4,
			North: "wall",
			South: "gameroom",
			East:  "wall",
			West:  "wall",
		},
	}

	reader := bufio.NewReader(os.Stdin)
	// var currentRoom Room
	roomName := "bathroom"
	currentRoom := house[roomName]
	exit := false
	for {
		if exit {
			break
		}

		fmt.Printf("You are in the %s\n", currentRoom.Name)

		// Handle user input
		fmt.Printf("What do you do? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			continue
		}
		input = input[:len(input)-1]
		inputArray := strings.Split(input, " ")

		if len(inputArray) < 2 {
			fmt.Printf("I can't understand you\n")
		} else {
			command := inputArray[0]
			if command == "go" {
				if inputArray[1] == "north" {
					roomName = currentRoom.North
					if roomName == "exit" {
						fmt.Printf("You have left the house. Thanks for playing!\n")
					} else if roomName == "wall" {
						fmt.Printf("There is a wall in the way!\n")
					} else {
						nextRoom, ok := house[roomName]
						if !ok {
							fmt.Printf("You have gotten lost in the house!\n")
						} else {
							currentRoom = nextRoom
						}
					}
				} else if inputArray[1] == "south" {
					roomName = currentRoom.South
					if roomName == "exit" {
						fmt.Printf("You have left the house. Thanks for playing!\n")
					} else if roomName == "wall" {
						fmt.Printf("There is a wall in the way!\n")
					} else {
						nextRoom, ok := house[roomName]
						if !ok {
							fmt.Printf("You have gotten lost in the house!\n")
						} else {
							currentRoom = nextRoom
						}
					}
				} else if inputArray[1] == "east" {
					roomName = currentRoom.East
					if roomName == "exit" {
						exit = true
					} else if roomName == "wall" {
						fmt.Printf("There is a wall in the way!\n")
					} else {
						nextRoom, ok := house[roomName]
						if !ok {
							fmt.Printf("You have gotten lost in the house!\n")
						} else {
							currentRoom = nextRoom
						}
					}
				} else if inputArray[1] == "west" {
					roomName = currentRoom.West
					if roomName == "exit" {
						fmt.Printf("You have left the house. Thanks for playing!\n")
					} else if roomName == "wall" {
						fmt.Printf("There is a wall in the way!\n")
					} else {
						nextRoom, ok := house[roomName]
						if !ok {
							fmt.Printf("You have gotten lost in the house!\n")
						} else {
							currentRoom = nextRoom
						}
					}
				} else {
					fmt.Printf("You can only go north, south, east or west...\n")
				}
			} else if command == "look" {
				fmt.Printf("You look around and see a %s\n", items[currentRoom.Item])
			} else if command == "get" {
				item := currentRoom.Item
				if item == 0 {
					fmt.Printf("There is nothing to get\n")
				} else {
					fmt.Printf("You picked up the %s\n", items[currentRoom.Item])
					currentRoom.Item = 0
					house[roomName] = currentRoom
				}
			} else {
				fmt.Printf("I don't understand '%s'\n", input)
			}
		}
	}

	fmt.Printf("You have left the house. Thanks for playing!\n")
}
