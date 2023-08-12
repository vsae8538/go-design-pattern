package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	subject := &SubjectImpl{}
	observer1 := &ObserverImpl{id: 1}
	observer2 := &ObserverImpl{id: 2}
	observer3 := &ObserverImpl{id: 3}

	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)

	res := subject.Notify()
	if res != "123" {
		t.Fail()
	}

	subject.Detach(observer2)

	res = subject.Notify()

	if res != "13" {
		t.Fail()
	}
}
