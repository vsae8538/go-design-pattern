// Package builder provides an implementation of the builder pattern for constructing complex objects.
package builder

// builder is the interface for builders in the builder pattern, defining the construction steps.
type builder interface {
	Part1()
	Part2()
	Part3()
}

// Director is responsible for executing the builder's steps in a specific order to construct the desired object.
type Director struct {
	builder builder
}

// NewDirector creates a new Director instance associated with the specified builder.
func NewDirector(builder builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct executes the construction process based on the builder's step sequence.
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// builder1 is a concrete implementation of builder, used for constructing a specific type of object.
type builder1 struct {
	result string
}

// Part1 implements the Part1 method of the builder interface.
func (b *builder1) Part1() {
	b.result += "1"
}

// Part2 implements the Part2 method of the builder interface.
func (b *builder1) Part2() {
	b.result += "2"
}

// Part3 implements the Part3 method of the builder interface.
func (b *builder1) Part3() {
	b.result += "3"
}

// GetResult returns the constructed result.
func (b *builder1) GetResult() string {
	return b.result
}

// builder2 is another concrete implementation of builder, used for constructing a different type of object.
type builder2 struct {
	result string
}

// Part1 implements the Part1 method of the builder interface.
func (b *builder2) Part1() {
	b.result += "a"
}

// Part2 implements the Part2 method of the builder interface.
func (b *builder2) Part2() {
	b.result += "b"
}

// Part3 implements the Part3 method of the builder interface.
func (b *builder2) Part3() {
	b.result += "c"
}

// GetResult returns the constructed result.
func (b *builder2) GetResult() string {
	return b.result
}
