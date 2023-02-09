package main

import (
	"bufio"
	"eli1/game"
	"fmt"
	"os"
	"strings"
)

func main() {

	database := game.NewGameDatabase()
	house := game.GetMap()

	fmt.Printf("===================================\n")
	fmt.Printf("|          Welcome to the         |\n")
	fmt.Printf("|         DUNGEON ADVENTURE       |\n")
	fmt.Printf("===================================\n\n")

	reader := bufio.NewReader(os.Stdin)

	// Set up our player
	fmt.Printf("What is your name? ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	player := game.NewPlayer(input, database)

	roomName := "bathroom"
	currentRoom := house[roomName]
	exit := false
	for {
		if exit {
			break
		}

		fmt.Printf("\nYou are in the %s\n", currentRoom.Name)

		monster, ok := database.Monsters[currentRoom.Monster]
		if ok {
			success := battle(reader, &player, &monster)
			if !success {
				fmt.Printf("Your adventure has come to an end...\n")
				return
			} else {
				currentRoom.Monster = ""
				house[roomName] = currentRoom
				fmt.Printf("\n==========================================\n")
				fmt.Printf("| BATTLE END                             |\n")
				fmt.Printf("==========================================\n\n")
				continue
			}
		}

		// Handle user input
		fmt.Printf("What do you do? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = input[:len(input)-1]
		inputArray := strings.Split(input, " ")
		fmt.Printf("\n")

		if len(inputArray) < 2 {
			fmt.Printf("I can't understand you\n")
		} else {
			command := inputArray[0]
			param := inputArray[1]
			if command == "go" {
				switch param {
				case "north":
					roomName = currentRoom.North
				case "south":
					roomName = currentRoom.South
				case "east":
					roomName = currentRoom.East
				case "west":
					roomName = currentRoom.West
				default:
					fmt.Printf("You can only go north, south, east or west...\n")
					continue
				}

				switch roomName {
				case "exit":
					fmt.Printf("You have left the house. Thanks for playing!\n")
				case "wall":
					fmt.Printf("There is a wall in the way!\n")
				case "locked":
					fmt.Printf("There is a locked door!\n")
				default:
					nextRoom, ok := house[roomName]
					if !ok {
						fmt.Printf("You have gotten lost in the house!\n")
					} else {
						currentRoom = nextRoom
					}
				}

			} else if command == "look" {
				if param == "around" || param == "room" {
					fmt.Printf("You look around and see:\n")
					currentRoom.PrintObjects(database)
					currentRoom.PrintInventory(database)
				} else {
					if currentRoom.HasObject(param, database) {
						object := database.Objects[param]
						fmt.Printf("You check the %s and find:\n", object.Name)
						object.PrintInventory(database)
					} else {
						fmt.Printf("%s doesn't exist\n", param)
					}
				}
			} else if command == "get" {
				if currentRoom.HasConsumable(param, database) {
					currentRoom.RemoveConsumable(param, database)
					house[roomName] = currentRoom
					player.Consumables[param]++
					fmt.Printf("You picked up %s\n", param)
				} else if currentRoom.HasEquipable(param, database) {
					currentRoom.RemoveEquipable(param, database)
					house[roomName] = currentRoom
					player.Equipables[param]++
					fmt.Printf("You picked up %s\n", param)
				} else {
					fmt.Printf("You can't pick that up\n")
				}
			} else if command == "use" {
				fmt.Printf("using %s\n", param)
			} else if command == "equip" {
				fmt.Printf("equipping %s\n", param)
			} else if command == "check" {
				if param == "inventory" {
					player.PrintInventory(database)
				} else if param == "status" {
					fmt.Printf("Status:\n")
					player.PrintStatus()
				} else {
					fmt.Printf("You can't check that!\n")
				}
			} else if command == "fly" {
				roomName = param
				nextRoom, ok := house[roomName]
				if !ok {
					fmt.Printf("You have gotten lost in the house!\n")
				} else {
					currentRoom = nextRoom
				}
			} else {
				fmt.Printf("I don't understand '%s'\n", input)
			}
		}
	}

	fmt.Printf("You have left the house. Thanks for playing!\n")
}

// Handles battling a monster, return true if you win
func battle(reader *bufio.Reader, player *game.Player, monster *game.Player) bool {

	fmt.Printf("==========================================\n")
	fmt.Printf("| BATTLE START                           |\n")
	fmt.Printf("==========================================\n\n")

	fmt.Printf("There is a %s in the room!\n\n", monster.Name)

	for {

		// Players turn
		fmt.Printf("What do you do? ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = input[:len(input)-1]
		inputArray := strings.Split(input, " ")

		command := inputArray[0]
		switch command {
		case "attack":
			damage := player.Attack - monster.Defense
			if damage < 0 {
				damage = 0
			}

			monster.Life -= damage
			fmt.Printf("%s dealt %d damage to the %s\n", player.Name, damage, monster.Name)

			if monster.Life <= 0 {
				fmt.Printf("You defeated the %s\n", monster.Name)
				fmt.Printf("\tGot %d gold!\n", monster.Gold)
				player.Gold += monster.Gold
				return true
			}
		default:
			fmt.Printf("You did nothing!\n")
		}

		// Monsters turn
		damage := monster.Attack - player.Defense
		if damage < 0 {
			damage = 0
		}

		player.Life -= damage
		fmt.Printf("The %s dealt %d to %s\n", monster.Name, damage, player.Name)

		if player.Life <= 0 {
			fmt.Printf("You died a miserable death at the hands of the %s\n", monster.Name)
			return false
		}
	}
}
