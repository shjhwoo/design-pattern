package factoryMethod

import "fmt"

func main() {
	ak47, _ := getCharacter("ak47")
	musket, _ := getCharacter("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g ICharacter) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
