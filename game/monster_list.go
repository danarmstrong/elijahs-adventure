package game

func GetMonsterList() map[string]Player {
	return map[string]Player{
		"slime": {
			Name:    "Slime",
			Life:    10,
			Gold:    1,
			Attack:  2,
			Defense: 1,
		},
	}
}
