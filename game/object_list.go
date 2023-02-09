package game

import "fmt"

type Object struct {
	Name        string
	Consumables []string
	Equipables  []string
}

func (r *Object) PrintConsumables(db *GameDatabase) {
	for _, v := range r.Consumables {
		c, ok := db.Consumables[v]
		if !ok {
			panic("invalid consumable " + v)
		}
		fmt.Printf("\t%s\n", c.Name)
	}
}

func (r *Object) RemoveConsumable(k string, db *GameDatabase) bool {
	for ix, v := range r.Consumables {
		if k == v {
			r.Consumables = append(r.Consumables[:ix], r.Consumables[ix+1:]...)
			return true
		}
	}
	return false
}

func (r *Object) PrintEquipables(db *GameDatabase) {
	for _, v := range r.Equipables {
		e, ok := db.Equipables[v]
		if !ok {
			panic("invalid equipable " + v)
		}
		fmt.Printf("\t%s\n", e.Name)
	}
}

func (r *Object) RemoveEquipable(k string, db *GameDatabase) bool {
	for ix, v := range r.Equipables {
		if k == v {
			r.Equipables = append(r.Equipables[:ix], r.Equipables[ix+1:]...)
			return true
		}
	}
	return false
}

func (r *Object) PrintInventory(db *GameDatabase) {
	r.PrintConsumables(db)
	r.PrintEquipables(db)
}

func GetObjectList() map[string]Object {
	return map[string]Object{
		"dresser": {
			Name: "DRESSER",
			Consumables: []string{
				"apple",
				"poison",
			},
			Equipables: []string{
				"hat",
			},
		},
	}
}
