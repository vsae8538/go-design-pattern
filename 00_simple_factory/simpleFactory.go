package simpleFactory

import "fmt"

// API is interface of all objects
type API interface {
	DoSomething() string
}

// APIFactory is factory of API
type APIFactory struct {
}

// NewAPI return API instance by type
func (apiFactory *APIFactory) NewAPI(ftype int) API {
	switch ftype {
	case 1:
		return &API1{}
	case 2:
		return &API2{}
	}
	return nil
}

// The following code is the implementation of the API interface.
type API1 struct {
}

// The following code is the implementation of the API interface.
type API2 struct {
}

// DoSomething is API1 implement
func (*API1) DoSomething() string {
	return fmt.Sprintln("API1 do something")
}

// DoSomething is API2 implement
func (*API2) DoSomething() string {
	return fmt.Sprintln("API2 do something")
}
