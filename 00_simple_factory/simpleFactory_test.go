package simpleFactory

import (
	"testing"
)

func TestAPI1(t *testing.T) {
	apiFactory := &APIFactory{}
	api := apiFactory.NewAPI(1)

	if api.DoSomething() != "API1 do something\n" {
		t.Fatal("API1 test fail")
	}
}

func TestAPI2(t *testing.T) {
	apiFactory := &APIFactory{}
	api := apiFactory.NewAPI(2)

	if api.DoSomething() != "API2 do something\n" {
		t.Fatal("API1 test fail")
	}
}
