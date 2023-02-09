package game

import "fmt"

type Room struct {
	Name        string
	Objects     []string
	Consumables []string
	Equipables  []string
	Monster     string
	North       string
	South       string
	East        string
	West        string
}

func (r *Room) PrintObjects(db *GameDatabase) {
	for _, v := range r.Objects {
		o, ok := db.Objects[v]
		if !ok {
			panic("invalid object " + v)
		}
		fmt.Printf("\t%s\n", o.Name)
	}
}

func (r *Room) PrintConsumables(db *GameDatabase) {
	for _, v := range r.Consumables {
		c, ok := db.Consumables[v]
		if !ok {
			panic("invalid consumable " + v)
		}
		fmt.Printf("\t%s\n", c.Name)
	}
}

func (r *Room) PrintEquipables(db *GameDatabase) {
	for _, v := range r.Equipables {
		e, ok := db.Equipables[v]
		if !ok {
			panic("invalid equipable " + v)
		}
		fmt.Printf("\t%s\n", e.Name)
	}
}

func (r *Room) PrintInventory(db *GameDatabase) {
	r.PrintConsumables(db)
	r.PrintEquipables(db)
}

func (r *Room) HasObject(k string, db *GameDatabase) bool {
	for _, v := range r.Objects {
		if k == v {
			return true
		}
	}

	return false
}

func (r *Room) HasConsumable(k string, db *GameDatabase) bool {
	for _, v := range r.Consumables {
		if k == v {
			return true
		}

	}

	for _, ov := range r.Objects {
		o, ok := db.Objects[ov]
		if !ok {
			panic("invalid object " + ov)
		}
		for _, v := range o.Consumables {
			if k == v {
				return true
			}
		}
	}

	return false
}

func (r *Room) RemoveConsumable(k string, db *GameDatabase) bool {
	for ix, v := range r.Consumables {
		if k == v {
			r.Consumables[ix] = r.Consumables[len(r.Consumables)-1]
			r.Consumables[len(r.Consumables)-1] = ""
			r.Consumables = r.Consumables[:len(r.Consumables)-1]
			return true
		}
	}

	for _, ov := range r.Objects {
		o := db.Objects[ov]
		if o.RemoveConsumable(k, db) {
			db.Objects[ov] = o
			return true
		}
	}

	return false
}

func (r *Room) HasEquipable(k string, db *GameDatabase) bool {
	for _, v := range r.Equipables {
		if k == v {
			return true
		}
	}

	for _, ov := range r.Objects {
		o, ok := db.Objects[ov]
		if !ok {
			panic("invalid object " + ov)
		}
		for _, v := range o.Equipables {
			if k == v {
				return true
			}
		}
	}

	return false
}

func (r *Room) RemoveEquipable(k string, db *GameDatabase) bool {
	for ix, v := range r.Equipables {
		if k == v {
			r.Equipables[ix] = r.Equipables[len(r.Equipables)-1]
			r.Equipables[len(r.Equipables)-1] = ""
			r.Equipables = r.Equipables[:len(r.Equipables)-1]
			return true
		}
	}

	for _, ov := range r.Objects {
		o := db.Objects[ov]
		if o.RemoveEquipable(k, db) {
			db.Objects[ov] = o
			return true
		}
	}

	return false
}

func (r *Room) HasItem(k string, db *GameDatabase) bool {

	if r.HasConsumable(k, db) {
		return true
	} else if r.HasEquipable(k, db) {
		return true
	}

	return false
}

func GetMap() map[string]Room {
	return map[string]Room{
		"bathroom": {
			Name: "Bathroom",
			Objects: []string{
				"dresser",
			},
			Equipables: []string{
				"stick",
			},
			North: "wall",
			South: "wall",
			East:  "exit",
			West:  "bedroom-1",
		},
		"bedroom-1": {
			Name: "Bedroom 1",

			North: "gameroom",
			South: "wall",
			East:  "bathroom",
			West:  "wall",
		},
		"gameroom": {
			Name:    "Game Room",
			Monster: "slime",
			North:   "baseball_room",
			South:   "bedroom-1",
			East:    "wall",
			West:    "wall",
		},
		"baseball_room": {
			Name:  "Baseball Room",
			North: "wall",
			South: "gameroom",
			East:  "wall",
			West:  "locked",
		},
	}
}
