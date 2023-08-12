package proxy

// Subject defines the interface for the subject and the proxy.
type Subject interface {
	Do() string
}

// RealSubject is the real object that implements the Subject interface.
type RealSubject struct{}

// Do is the implementation of the Do method for the RealSubject.
func (RealSubject) Do() string {
	return "real"
}

// Proxy is a proxy object that wraps around the RealSubject and provides additional functionality.
type Proxy struct {
	real RealSubject
}

// Do is the implementation of the Do method for the Proxy.
func (p Proxy) Do() string {
	var res string
	res += "pre:"      // Pre-processing before invoking the real subject's method.
	res += p.real.Do() // Invoking the real subject's method.
	res += ":after"    // Post-processing after invoking the real subject's method.
	return res
}
