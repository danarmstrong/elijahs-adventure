package game

type GameDatabase struct {
	Consumables map[string]Consumable
	Equipables  map[string]Equipable
	Monsters    map[string]Player
	Objects     map[string]Object
}

func NewGameDatabase() *GameDatabase {
	return &GameDatabase{
		Consumables: GetConsumableList(),
		Equipables:  GetEquipableList(),
		Monsters:    GetMonsterList(),
		Objects:     GetObjectList(),
	}
}
