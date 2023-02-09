package main

import (
	"bufio"
	"eli1/game"
	"fmt"
	"os"
	"strings"
)

func main() {

	items := []string{
		"Nothing",
		"Towel",
		"Bed",
		"TV",
		"Baseball Bat",
	}

	monsters := game.GetMonsterList()

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

	player := game.Player{
		Name:    input,
		Life:    100,
		Gold:    0,
		Attack:  5,
		Defense: 1,
	}

	roomName := "bathroom"
	currentRoom := house[roomName]
	exit := false
	for {
		if exit {
			break
		}

		fmt.Printf("\nYou are in the %s\n", currentRoom.Name)

		monster, ok := monsters[currentRoom.Monster]
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
			if command == "go" {
				param := inputArray[1]
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
				fmt.Printf("You look around and see a %s\n", items[currentRoom.Item])
			} else if command == "get" {
				item := currentRoom.Item
				if item == 0 {
					fmt.Printf("There is nothing to get\n")
				} else {
					fmt.Printf("You picked up the %s\n", items[currentRoom.Item])
					currentRoom.Item = 0
					house[roomName] = currentRoom
					player.Inventory = append(player.Inventory, item)
				}
			} else if command == "check" {
				param := inputArray[1]
				if param == "inventory" {
					if len(player.Inventory) == 0 {
						fmt.Printf("You aren't carrying anything.\n")
					} else {
						fmt.Printf("You are carrying:\n")
						for index, itemNumber := range player.Inventory {
							fmt.Printf("\t%d: %s\n", index+1, items[itemNumber])
						}
					}
				} else if param == "life" {
					fmt.Printf("You have %d HP\n", player.Life)
				} else if param == "gold" {
					fmt.Printf("You have %d gold\n", player.Gold)
				} else {
					fmt.Printf("You can't check that!\n")
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
