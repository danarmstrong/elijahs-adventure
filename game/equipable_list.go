package game

type BodyPart string

const (
	BodyHead      BodyPart = "HEAD"
	BodyChest     BodyPart = "CHEST"
	BodyLegs      BodyPart = "LEGS"
	BodyFeet      BodyPart = "FEET"
	BodyLeftHand  BodyPart = "LEFT_HAND"
	BodyRightHand BodyPart = "RIGHT_HAND"
)

type Equipable struct {
	Name string
	BodyPart
	Life    int
	Attack  int
	Defense int
	Value   int
}

func GetEquipableList() map[string]Equipable {
	return map[string]Equipable{
		"hat": {
			Name:     "HAT",
			BodyPart: BodyHead,
			Defense:  1,
			Value:    0,
		},
		"stick": {
			Name:     "STICK",
			BodyPart: BodyLeftHand,
			Attack:   1,
			Value:    0,
		},
	}
}
