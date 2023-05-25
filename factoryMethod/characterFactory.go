package factoryMethod

import "fmt"

func getCharacter(characterType string) (ICharacter, error) {
	if characterType == "npc" {
		return newNPC(), nil
	}
	if characterType == "villian" {
		return newVillian(), nil
	}
	return nil, fmt.Errorf("Wrong character type passed")
}
