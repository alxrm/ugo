package ugo

type wire interface {
	Map(cb Callback) wire
	Filter(cb Predicate) wire
	Each(cb Action) wire
	Reject(cb Predicate) wire
	Reduce(cb Collector, initial Object) wire
	ReduceRight(cb Collector, initial Object) wire
	Min(cb Comparator) wire
	Max(cb Comparator) wire
	Find(cb Predicate) wire
	FindLast(cb Predicate) wire
	FindIndex(cb Predicate) wire
	FindLastIndex(cb Predicate) wire
	Value() Object
}

type chainWrapper struct {
	mid Seq
	res Object
}

func (wrapper *chainWrapper) Each(cb Action) wire {
	Each(wrapper.mid, cb)
	return wire(wrapper)
}

func (wrapper *chainWrapper) Map(cb Callback) wire {
	wrapper.mid = Map(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return wire(wrapper)
}

func (wrapper *chainWrapper) Filter(cb Predicate) wire {
	wrapper.mid = Filter(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return wire(wrapper)
}

func (wrapper *chainWrapper) Reject(cb Predicate) wire {
	wrapper.mid = Reject(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return wire(wrapper)
}

func (wrapper *chainWrapper) Reduce(cb Collector, initial Object) wire {
	wrapper.res = Reduce(wrapper.mid, cb, initial)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) ReduceRight(cb Collector, initial Object) wire {
	wrapper.res = ReduceRight(wrapper.mid, cb, initial)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) Min(cb Comparator) wire {
	wrapper.res = Min(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) Max(cb Comparator) wire {
	wrapper.res = Max(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) Find(cb Predicate) wire {
	wrapper.res = Find(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) FindLast(cb Predicate) wire {
	wrapper.res = FindLast(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) FindIndex(cb Predicate) wire {
	wrapper.res = FindIndex(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) FindLastIndex(cb Predicate) wire {
	wrapper.res = FindLastIndex(wrapper.mid, cb)
	wrapper.mid = nil
	return wire(wrapper)
}

func (wrapper *chainWrapper) Value() Object {
	return wrapper.res
}