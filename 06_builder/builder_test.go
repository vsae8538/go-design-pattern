package builder

import "testing"

func TestBuilder(t *testing.T) {
	builder1 := &builder1{}
	builder2 := &builder2{}

	director := NewDirector(builder1)
	director.Construct()
	if builder1.GetResult() != "123" {
		t.Fatal("builder1 fail")
	}

	director = NewDirector(builder2)
	director.Construct()
	if builder2.GetResult() != "abc" {
		t.Fatal("builder2 fail")
	}
}
