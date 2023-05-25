package factoryMethod

type NPC struct {
	Character
}

func newNPC() ICharacter {
	return &NPC{
		Character: Character{
			name:  "black widow",
			power: 4,
		},
	}
}
