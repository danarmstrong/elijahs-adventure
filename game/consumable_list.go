package game

type Consumable struct {
	Name    string
	Life    int
	Attack  int
	Defense int
	Value   int
}

func GetConsumableList() map[string]Consumable {
	return map[string]Consumable{
		"apple": {
			Name:  "APPLE",
			Life:  3,
			Value: 1,
		},
		"potion": {
			Name:  "POTION",
			Life:  20,
			Value: 5,
		},
		"poison": {
			Name:    "POISON",
			Life:    -10,
			Defense: -2,
			Value:   10,
		},
		"cure": {
			Name:    "CURE",
			Life:    100,
			Attack:  0,
			Defense: 0,
			Value:   50,
		},
		"protein": {
			Name:   "PROTEIN",
			Attack: 3,
		},
		"biscuit": {
			Name:    "BISCUIT",
			Defense: 3,
		},
	}
}
