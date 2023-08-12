package factorymethod

// Operator is interface
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory is interface for Operator factory
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase is base struct for Operator
type OperatorBase struct {
	a, b int
}

// SetA sets a
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB sets b
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory is PlusOperator factory
type PlusOperatorFactory struct{}

// Create return a new PlusOperator
func (p *PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

// PlusOperator is operator for Plus
type PlusOperator struct {
	*OperatorBase
}

// Result returns a+b
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory is MinusOperator factory
type MinusOperatorFactory struct{}

// Create return a new MinusOperator
func (m *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}

// MinusOperator is operator for Minus
type MinusOperator struct {
	*OperatorBase
}

// Result returns a-b
func (o MinusOperator) Result() int {
	return o.a - o.b
}
