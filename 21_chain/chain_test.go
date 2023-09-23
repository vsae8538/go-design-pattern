package chain

import "testing"

func TestChain(t *testing.T) {
	c := NewProjectManagerChain()
	c.SetSuccessor(NewDepManagerChain())
	res := c.HandleFeeRequest("bob", 400)
	if !res {
		t.Fatal("Chain test fail")
	}
	res = c.HandleFeeRequest("tom", 400)
	if res {
		t.Fatal("Chain test fail")
	}
}
