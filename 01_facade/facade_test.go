package facade

import (
	"testing"
)

func TestAPI(t *testing.T) {
	api := NewAPI()
	s := api.Test()
	if s != "A module running\nB module running" {
		t.Fatal("error")
	}
}
