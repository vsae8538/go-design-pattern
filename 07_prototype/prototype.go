package prototype

// Cloneable is an interface that defines the behavior of objects that can be cloned.
type Cloneable interface {
	Clone() Cloneable
}

// PrototypeManager manages a collection of prototypes and provides methods for getting and setting them.
type PrototypeManager struct {
	prototypes map[string]Cloneable
}

// NewPrototypeManager creates a new PrototypeManager instance with an empty map of prototypes.
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{prototypes: make(map[string]Cloneable)}
}

// Get retrieves a clone of a prototype based on the specified name.
func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name].Clone()
}

// Set stores a prototype with the given name in the PrototypeManager.
func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
