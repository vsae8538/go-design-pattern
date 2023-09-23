package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func (n *NumbersIterator) First() {
	n.next = n.numbers.start
}

func (n *NumbersIterator) Next() interface{} {
	if !n.IsDone() {
		next := n.next
		n.next++
		return next
	}
	return nil
}

func (n *NumbersIterator) IsDone() bool {
	return n.next > n.numbers.end
}

func (n *NumbersIterator) CurrentItem() interface{} {
	if !n.IsDone() {
		return n.next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	for i.First(); !i.IsDone(); {
		c := i.CurrentItem()
		if c != nil {
			println(c.(int))
		}
		i.Next()
	}
}
