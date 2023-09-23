package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	// Can't use c := &ConcreteComponent{}
	// Because ConcreteComponent can't be transformed into MulDecorator, AddDecorator.c to be declared as Component.
	var c Component = &ConcreteComponent{}
	c = WrapMulDecorator(c, 3)
	c = WrapAddDecorator(c, 6)
	res := c.Calc()
	if res != 6 {
		t.Fatal("Decorator test fail")
	}
}
