package facade

import "fmt"

// API is facade interface of facade package
type API interface {
	Test() string
}

// NewAPI return implementation of API
func NewAPI() API {
	aModuleImpl := NewAModuleAPI()
	bModuleImpl := NewBModuleAPI()
	return &apiImpl{
		a: aModuleImpl,
		b: bModuleImpl,
	}
}

// facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

// Test is facade method
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// The following code is the implementation of the A module and the B module.
type AModuleAPI interface {
	TestA() string
}

// NewAModuleAPI return new instance of AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// aModuleImpl is a struct to implement AModuleAPI
type aModuleImpl struct {
}

// TestA is a method of aModuleImpl
func (a *aModuleImpl) TestA() string {
	return "A module running"
}

// The following code is the implementation of the B module.
type BModuleAPI interface {
	TestB() string
}

// NewBModuleAPI return new instance of BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// bModuleImpl is a struct to implement BModuleAPI
type bModuleImpl struct {
}

// TestB is a method of bModuleImpl
func (b *bModuleImpl) TestB() string {
	return "B module running"
}
