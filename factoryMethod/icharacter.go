package factoryMethod

type ICharacter interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
