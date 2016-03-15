// The MIT License (MIT)
//
// Copyright (c) 2016 Alexey Derbyshev
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ugo

// Wire is an interface with all of the ugo functions, used for chaining
type Wire interface {
	Map(cb Callback) Wire
	Filter(cb Predicate) Wire
	Select(cb Predicate) Wire
	Each(cb Action) Wire
	ForEach(cb Action) Wire
	Reject(cb Predicate) Wire
	Reduce(cb Collector, initial Object) Wire
	Inject(cb Collector, initial Object) Wire
	FoldL(cb Collector, initial Object) Wire
	ReduceRight(cb Collector, initial Object) Wire
	FoldR(cb Collector, initial Object) Wire
	Min(cb Comparator) Wire
	Max(cb Comparator) Wire
	Find(cb Predicate) Wire
	Detect(cb Predicate) Wire
	FindLast(cb Predicate) Wire
	FindIndex(cb Predicate) Wire
	FindLastIndex(cb Predicate) Wire
	Some(cb Predicate) Wire
	Any(cb Predicate) Wire
	IndexOf(target Object, isSorted bool, cb Comparator) Wire
	LastIndexOf(target Object, cb Comparator) Wire
	Contains(target Object, isSorted bool, cb Comparator) Wire
	Includes(target Object, isSorted bool, cb Comparator) Wire
	Every(cb Predicate) Wire
	All(cb Predicate) Wire
	Uniq(cb Comparator) Wire
	Unique(cb Comparator) Wire
	Difference(other Seq, cb Comparator) Wire
	Without(nonGrata Object, cb Comparator) Wire
	Intersection(other Seq, cb Comparator) Wire
	Union(other Seq, cb Comparator) Wire
	SortBy(cb Comparator) Wire
	CountBy(cb Callback) Wire
	GroupBy(cb Callback) Wire
	Remove(pos int) Wire
	Insert(tg Object, pos int) Wire
	Concat(next Seq) Wire
	Shuffle() Wire
	Reverse() Wire
	EqualsStrict(other Seq, cb Comparator) Wire
	EqualsNotStrict(other Seq, cb Comparator) Wire
	Value() Object
}

type chainWrapper struct {
	mid Seq
	res Object
}

func (wrapper *chainWrapper) Each(cb Action) Wire {
	Each(wrapper.mid, cb)
	return Wire(wrapper)
}

func (wrapper *chainWrapper) ForEach(cb Action) Wire {
	return wrapper.Each(cb)
}

func (wrapper *chainWrapper) Map(cb Callback) Wire {
	wrapper.mid = Map(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Collect(cb Callback) Wire {
	return wrapper.Map(cb)
}

func (wrapper *chainWrapper) Filter(cb Predicate) Wire {
	wrapper.mid = Filter(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Select(cb Predicate) Wire {
	return wrapper.Filter(cb)
}

func (wrapper *chainWrapper) Reject(cb Predicate) Wire {
	wrapper.mid = Reject(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Reduce(cb Collector, initial Object) Wire {
	wrapper.res = Reduce(wrapper.mid, cb, initial)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Inject(cb Collector, initial Object) Wire {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *chainWrapper) FoldL(cb Collector, initial Object) Wire {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *chainWrapper) ReduceRight(cb Collector, initial Object) Wire {
	wrapper.res = ReduceRight(wrapper.mid, cb, initial)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) FoldR(cb Collector, initial Object) Wire {
	return wrapper.ReduceRight(cb, initial)
}

func (wrapper *chainWrapper) Min(cb Comparator) Wire {
	wrapper.res = Min(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Max(cb Comparator) Wire {
	wrapper.res = Max(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Find(cb Predicate) Wire {
	wrapper.res = Find(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Detect(cb Predicate) Wire {
	return wrapper.Find(cb)
}

func (wrapper *chainWrapper) FindLast(cb Predicate) Wire {
	wrapper.res = FindLast(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) FindIndex(cb Predicate) Wire {
	wrapper.res = FindIndex(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) FindLastIndex(cb Predicate) Wire {
	wrapper.res = FindLastIndex(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Some(cb Predicate) Wire {
	wrapper.res = Some(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Any(cb Predicate) Wire {
	return wrapper.Some(cb)
}

func (wrapper *chainWrapper) IndexOf(target Object, isSorted bool, cb Comparator) Wire {
	wrapper.res = IndexOf(wrapper.mid, target, isSorted, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) LastIndexOf(target Object, cb Comparator) Wire {
	wrapper.res = LastIndexOf(wrapper.mid, target, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Contains(target Object, isSorted bool, cb Comparator) Wire {
	wrapper.res = Contains(wrapper.mid, target, isSorted, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Includes(target Object, isSorted bool, cb Comparator) Wire {
	return wrapper.Contains(target, isSorted, cb)
}

func (wrapper *chainWrapper) Every(cb Predicate) Wire {
	wrapper.res = Every(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) All(cb Predicate) Wire {
	return wrapper.Every(cb)
}

func (wrapper *chainWrapper) Uniq(cb Comparator) Wire {
	wrapper.mid = Uniq(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Unique(cb Comparator) Wire {
	return wrapper.Uniq(cb)
}

func (wrapper *chainWrapper) Difference(other Seq, cb Comparator) Wire {
	wrapper.mid = Difference(wrapper.mid, other, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Without(nonGrata Object, cb Comparator) Wire {
	wrapper.mid = Without(wrapper.mid, nonGrata, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Intersection(other Seq, cb Comparator) Wire {
	wrapper.mid = Intersection(wrapper.mid, other, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Union(other Seq, cb Comparator) Wire {
	wrapper.mid = Union(wrapper.mid, other, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) SortBy(cb Comparator) Wire {
	wrapper.mid = SortBy(wrapper.mid, cb)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) CountBy(cb Callback) Wire {
	wrapper.res = CountBy(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) GroupBy(cb Callback) Wire {
	wrapper.res = GroupBy(wrapper.mid, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Remove(pos int) Wire {
	wrapper.mid = Remove(wrapper.mid, pos)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Insert(tg Object, pos int) Wire {
	wrapper.mid = Insert(wrapper.mid, tg, pos)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Concat(next Seq) Wire {
	wrapper.mid = Concat(wrapper.mid, next)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Shuffle() Wire {
	wrapper.mid = ShuffledCopy(wrapper.mid)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Reverse() Wire {
	wrapper.mid = ReversedCopy(wrapper.mid)
	wrapper.res = wrapper.mid
	return Wire(wrapper)
}

func (wrapper *chainWrapper) EqualsStrict(other Seq, cb Comparator) Wire {
	wrapper.res = EqualsStrict(wrapper.mid, other, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) EqualsNotStrict(other Seq, cb Comparator) Wire {
	wrapper.res = EqualsNotStrict(wrapper.mid, other, cb)
	wrapper.mid = nil
	return Wire(wrapper)
}

func (wrapper *chainWrapper) Value() Object {
	return wrapper.res
}
