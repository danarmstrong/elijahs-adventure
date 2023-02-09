package game

func GetMonsterList() map[string]Player {
	return map[string]Player{
		"slime": {
			Name:    "SLIME",
			Life:    10,
			Gold:    1,
			Attack:  2,
			Defense: 1,
			Inventory: Inventory{
				Consumables: map[string]int{
					"apple": 1,
				},
			},
		},
	}
}
