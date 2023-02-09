package game

import "fmt"

type Status struct {
	Attack  int
	Defense int
}

type Equipped struct {
	Head      string
	Chest     string
	Legs      string
	Feet      string
	LeftHand  string
	RightHand string
}

type Inventory struct {
	Items       map[string]int
	Consumables map[string]int
	Equipables  map[string]int
}

type Player struct {
	Name    string
	Life    int
	Gold    int
	Attack  int
	Defense int
	Inventory
	Equipped
	Status
}

func (r *Player) PrintConsumables(db *GameDatabase) {
	for k, v := range r.Inventory.Consumables {
		c, ok := db.Consumables[k]
		if !ok {
			panic("invalid consumable " + k)
		}
		if v > 0 {
			fmt.Printf("\t%d X %s\n", v, c.Name)
		}
	}
}

func (r *Player) CountConsumables() int {
	c := 0
	for _, v := range r.Consumables {
		c += v
	}

	return c
}

func (r *Player) PrintEquipables(db *GameDatabase) {
	for k, v := range r.Inventory.Equipables {
		e, ok := db.Equipables[k]
		if !ok {
			panic("invalid equipable " + k)
		}
		if v > 0 {
			fmt.Printf("\t%d X %s\n", v, e.Name)
		}
	}
}

func (r *Player) CountEquipables() int {
	c := 0
	for _, v := range r.Equipables {
		c += v
	}

	return c
}

func (r *Player) PrintInventory(db *GameDatabase) {
	if r.CountConsumables() == 0 && r.CountEquipables() == 0 {
		fmt.Printf("You aren't carrying anything\n")
	} else {
		fmt.Printf("You are carrying:\n")
		r.PrintConsumables(db)
		r.PrintEquipables(db)
	}
}

func (r *Player) PrintStatus() {
	fmt.Printf("  Gold    : %d\n", r.Gold)
	fmt.Printf("  Life    : %d\n", r.Life)
	fmt.Printf("  Attack  : %d\n", r.Attack)
	fmt.Printf("  Defense : %d\n", r.Defense)
}

func (r *Player) Use(k string, db *GameDatabase) {
	i, ok := r.Consumables[k]
	if !ok || i == 0 {
		fmt.Printf("You can't use that\n")
		return
	}

	c := db.Consumables[k]
	r.Life += c.Life
	// TODO figure out the status thing...
}

func NewPlayer(name string, db *GameDatabase) Player {

	/*items := make(map[string]int)
	for k := range db.Items {
		items[k] = 0
	}*/

	consumables := make(map[string]int)
	for k := range db.Consumables {
		consumables[k] = 0
	}

	equipables := make(map[string]int)
	for k := range db.Equipables {
		equipables[k] = 0
	}

	return Player{
		Name:    name,
		Life:    100,
		Gold:    0,
		Attack:  5,
		Defense: 1,
		Inventory: Inventory{
			Equipables:  equipables,
			Consumables: consumables,
		},
	}
}
