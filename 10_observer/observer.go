package observer

import "fmt"

// Subject defines the interface that subjects implement for attaching, detaching, and notifying observers.
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify() string
}

// SubjectImpl implements the Subject interface and maintains a list of attached observers.
type SubjectImpl struct {
	observers []*Observer
}

// Attach adds an observer to the list of attached observers.
func (s *SubjectImpl) Attach(o Observer) {
	s.observers = append(s.observers, &o)
}

// Detach removes an observer from the list of attached observers.
func (s *SubjectImpl) Detach(o Observer) {
	for i, observer := range s.observers {
		if *observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

// Notify notifies all attached observers and collects their responses.
func (s *SubjectImpl) Notify() string {
	var res string
	for _, observer := range s.observers {
		res += (*observer).Update()
	}
	return res
}

// Observer defines the interface for observers that receive updates from subjects.
type Observer interface {
	Update() string
}

// ObserverImpl implements the Observer interface and provides a basic implementation for Update.
type ObserverImpl struct {
	id int
}

// Update is called by the subject to notify the observer about changes and returns the observer's ID.
func (o *ObserverImpl) Update() string {
	return fmt.Sprintf("%d", o.id)
}
