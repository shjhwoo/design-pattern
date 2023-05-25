package factoryMethod

type Character struct {
	name  string
	power int
}

func (g *Character) setName(name string) {
	g.name = name
}

func (g *Character) getName() string {
	return g.name
}

func (g *Character) setPower(power int) {
	g.power = power
}

func (g *Character) getPower() int {
	return g.power
}
