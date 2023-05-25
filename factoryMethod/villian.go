package factoryMethod

type villian struct {
	Character
}

func newVillian() ICharacter {
	return &villian{
		Character: Character{
			name:  "purple mamba",
			power: 1,
		},
	}
}
