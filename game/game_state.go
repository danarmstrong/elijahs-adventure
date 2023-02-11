package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameState struct {
	reader   *bufio.Reader
	database *GameDatabase
	gameMap  map[string]Room
	player   Player

	roomName    string
	currentRoom Room
	exit        bool
}

func CreateGameState() *GameState {
	return &GameState{
		reader:   bufio.NewReader(os.Stdin),
		database: NewGameDatabase(),
		gameMap:  GetMap(),
	}
}

func (state *GameState) Run() {
	state.initGame()
	state.printTitleScreen()
	state.createPlayer()

}

func (state *GameState) initGame() {
	state.roomName = "bathroom"
	state.currentRoom = state.gameMap[state.roomName]
	state.exit = false
}

func (state *GameState) printTitleScreen() {
	fmt.Printf("===================================\n")
	fmt.Printf("|          Welcome to the         |\n")
	fmt.Printf("|         DUNGEON ADVENTURE       |\n")
	fmt.Printf("===================================\n\n")
}

func (state *GameState) createPlayer() {
	input := state.prompt("What is your name?")
	state.player = NewPlayer(input, state.database)
}

func (state *GameState) gameLoop() {
	for {
		if state.exit {
			break
		}

		fmt.Printf("\nYou are in the %s\n", state.currentRoom.Name)

		monster, ok := state.database.Monsters[state.currentRoom.Monster]
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
}

// Handles battling a monster, return true if you win
func (state *GameState) runBattle(monster Player) bool {

	fmt.Printf("==========================================\n")
	fmt.Printf("| BATTLE START                           |\n")
	fmt.Printf("==========================================\n\n")

	fmt.Printf("There is a %s in the room!\n\n", monster.Name)

	for {

		// Players turn
		input := state.prompt("What do you do?")
		inputArray := state.parseInput(input)

		command := inputArray[0]
		switch command {
		case "attack":
			damage := state.player.Attack - monster.Defense
			if damage < 0 {
				damage = 0
			}

			monster.Life -= damage
			fmt.Printf("%s dealt %d damage to the %s\n", state.player.Name, damage, monster.Name)

			if monster.Life <= 0 {
				fmt.Printf("You defeated the %s\n", monster.Name)
				fmt.Printf("\tGot %d gold!\n", monster.Gold)
				state.player.Gold += monster.Gold
				return true
			}
		default:
			fmt.Printf("You did nothing!\n")
		}

		// Monsters turn
		damage := monster.Attack - state.player.Defense
		if damage < 0 {
			damage = 0
		}

		state.player.Life -= damage
		fmt.Printf("The %s dealt %d to %s\n", monster.Name, damage, state.player.Name)

		if state.player.Life <= 0 {
			fmt.Printf("You died a miserable death at the hands of the %s\n", monster.Name)
			return false
		}
	}
}

func (state *GameState) prompt(text string) string {
	fmt.Printf("%s ")
	input, err := state.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return input[:len(input)-1]
}

func (state *GameState) parseInput(input string) []string {
	return strings.Split(input, " ")
}
