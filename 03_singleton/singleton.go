package singleton

import (
	"sync"
)

// Singleton is interface
type Singleton interface {
	foo()
}

// singletonImpl is a struct to implement Singleton
type singletonImpl struct {
}

// foo is a method of singletonImpl
func (s *singletonImpl) foo() {
}

// singleton is a variable to store singletonImpl instance
var (
	singleton *singletonImpl
	once      sync.Once
)

// GetSingletonInstance return singleton instance
func GetSingletonInstance() Singleton {
	once.Do(func() {
		singleton = &singletonImpl{}
	})
	return singleton
}
