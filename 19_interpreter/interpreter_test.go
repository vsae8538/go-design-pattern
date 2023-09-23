package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	p := &Parser{}
	p.Parse([]string{"1", "+", "2", "-", "3"})
	res := p.prev.Interpret()
	if res != 0 {
		t.Fatal("Interpreter test fail")
	}
}
