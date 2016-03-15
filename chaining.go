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
	Min(cb Comparator) Wire // tested
	Max(cb Comparator) Wire // tested
	Find(cb Predicate) Wire //
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

// ChainWrapper is the special struct,
// containing resulting and middleware data
type ChainWrapper struct {
	Mid Seq // Mid is for middleware calculations
	Res Object // Res if for resulting data
}

func (wrapper *ChainWrapper) Each(cb Action) Wire {
	Each(wrapper.Mid, cb)
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) ForEach(cb Action) Wire {
	return wrapper.Each(cb)
}

func (wrapper *ChainWrapper) Map(cb Callback) Wire {
	wrapper.Mid = Map(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Collect(cb Callback) Wire {
	return wrapper.Map(cb)
}

func (wrapper *ChainWrapper) Filter(cb Predicate) Wire {
	wrapper.Mid = Filter(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Select(cb Predicate) Wire {
	return wrapper.Filter(cb)
}

func (wrapper *ChainWrapper) Reject(cb Predicate) Wire {
	wrapper.Mid = Reject(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Reduce(cb Collector, initial Object) Wire {
	wrapper.Res = Reduce(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Inject(cb Collector, initial Object) Wire {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *ChainWrapper) FoldL(cb Collector, initial Object) Wire {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *ChainWrapper) ReduceRight(cb Collector, initial Object) Wire {
	wrapper.Res = ReduceRight(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) FoldR(cb Collector, initial Object) Wire {
	return wrapper.ReduceRight(cb, initial)
}

func (wrapper *ChainWrapper) Min(cb Comparator) Wire {
	wrapper.Res = Min(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Max(cb Comparator) Wire {
	wrapper.Res = Max(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Find(cb Predicate) Wire {
	wrapper.Res = Find(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Detect(cb Predicate) Wire {
	return wrapper.Find(cb)
}

func (wrapper *ChainWrapper) FindLast(cb Predicate) Wire {
	wrapper.Res = FindLast(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) FindIndex(cb Predicate) Wire {
	wrapper.Res = FindIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) FindLastIndex(cb Predicate) Wire {
	wrapper.Res = FindLastIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Some(cb Predicate) Wire {
	wrapper.Res = Some(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Any(cb Predicate) Wire {
	return wrapper.Some(cb)
}

func (wrapper *ChainWrapper) IndexOf(target Object, isSorted bool, cb Comparator) Wire {
	wrapper.Res = IndexOf(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) LastIndexOf(target Object, cb Comparator) Wire {
	wrapper.Res = LastIndexOf(wrapper.Mid, target, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Contains(target Object, isSorted bool, cb Comparator) Wire {
	wrapper.Res = Contains(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Includes(target Object, isSorted bool, cb Comparator) Wire {
	return wrapper.Contains(target, isSorted, cb)
}

func (wrapper *ChainWrapper) Every(cb Predicate) Wire {
	wrapper.Res = Every(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) All(cb Predicate) Wire {
	return wrapper.Every(cb)
}

func (wrapper *ChainWrapper) Uniq(cb Comparator) Wire {
	wrapper.Mid = Uniq(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Unique(cb Comparator) Wire {
	return wrapper.Uniq(cb)
}

func (wrapper *ChainWrapper) Difference(other Seq, cb Comparator) Wire {
	wrapper.Mid = Difference(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Without(nonGrata Object, cb Comparator) Wire {
	wrapper.Mid = Without(wrapper.Mid, nonGrata, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Intersection(other Seq, cb Comparator) Wire {
	wrapper.Mid = Intersection(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Union(other Seq, cb Comparator) Wire {
	wrapper.Mid = Union(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) SortBy(cb Comparator) Wire {
	wrapper.Mid = SortBy(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) CountBy(cb Callback) Wire {
	wrapper.Res = CountBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) GroupBy(cb Callback) Wire {
	wrapper.Res = GroupBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Remove(pos int) Wire {
	wrapper.Mid = Remove(wrapper.Mid, pos)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Insert(tg Object, pos int) Wire {
	wrapper.Mid = Insert(wrapper.Mid, tg, pos)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Concat(next Seq) Wire {
	wrapper.Mid = Concat(wrapper.Mid, next)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Shuffle() Wire {
	wrapper.Mid = ShuffledCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Reverse() Wire {
	wrapper.Mid = ReversedCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) EqualsStrict(other Seq, cb Comparator) Wire {
	wrapper.Res = EqualsStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) EqualsNotStrict(other Seq, cb Comparator) Wire {
	wrapper.Res = EqualsNotStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return Wire(wrapper)
}

func (wrapper *ChainWrapper) Value() Object {
	return wrapper.Res
}
