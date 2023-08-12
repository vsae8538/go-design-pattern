package adapter

import "fmt"

// Target is the interface that Client uses
type Target interface {
	Request() string
}

// Adaptee is the interface that needs to be adapted
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// NewAdaptee return new instance of Adaptee
type adapteeImpl struct {
}

// NewAdaptee return new instance of Adaptee
func (a *adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewTarget return new instance of Target
type adapter struct {
	Adaptee
}

// NewAdapter return new instance of Target
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// NewAdapter return new instance of Target
func (a *adapter) Request() string {
	fmt.Println("adapter do something before")
	return a.SpecificRequest()
}
