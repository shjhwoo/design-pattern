package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	Ice        int
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

/**
 * Concrete Builders are supposed to provide their own methods for
 * retrieving results. That's because various types of builders may create
 * entirely different products that don't follow the same interface.
 * Therefore, such methods cannot be declared in the base Builder interface
 * (at least in a statically typed programming language).
 *
 * Usually, after returning the end result to the client, a builder instance
 * is expected to be ready to start producing another product. That's why
 * it's a usual practice to call the reset method at the end of the
 * `getProduct` method body. However, this behavior is not mandatory, and
 * you can make your builders wait for an explicit reset call from the
 * client code before disposing of the previous result.
 */

//다같은 인터페이스를 따르기 때문에, 다양한 빌더를 만들기 힘들다.
//이 때문에 고 언어에서는 인터페이스부터 다양하게 만들어줘야 함.
