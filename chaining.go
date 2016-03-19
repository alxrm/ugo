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

// ChainWrapper is the special struct,
// containing resulting and middleware data
type ChainWrapper struct {
	Mid Seq    // Mid is for middleware calculations
	Res Object // Res if for resulting data
}

func (wrapper *ChainWrapper) Each(cb Action) *ChainWrapper {
	Each(wrapper.Mid, cb)
	return wrapper
}

func (wrapper *ChainWrapper) ForEach(cb Action) *ChainWrapper {
	return wrapper.Each(cb)
}

func (wrapper *ChainWrapper) Map(cb Callback) *ChainWrapper {
	wrapper.Mid = Map(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Collect(cb Callback) *ChainWrapper {
	return wrapper.Map(cb)
}

func (wrapper *ChainWrapper) Filter(cb Predicate) *ChainWrapper {
	wrapper.Mid = Filter(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Select(cb Predicate) *ChainWrapper {
	return wrapper.Filter(cb)
}

func (wrapper *ChainWrapper) Reject(cb Predicate) *ChainWrapper {
	wrapper.Mid = Reject(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Reduce(cb Collector, initial Object) *ChainWrapper {
	wrapper.Res = Reduce(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Inject(cb Collector, initial Object) *ChainWrapper {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *ChainWrapper) FoldL(cb Collector, initial Object) *ChainWrapper {
	return wrapper.Reduce(cb, initial)
}

func (wrapper *ChainWrapper) ReduceRight(cb Collector, initial Object) *ChainWrapper {
	wrapper.Res = ReduceRight(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) FoldR(cb Collector, initial Object) *ChainWrapper {
	return wrapper.ReduceRight(cb, initial)
}

func (wrapper *ChainWrapper) Min(cb Comparator) *ChainWrapper {
	wrapper.Res = Min(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Max(cb Comparator) *ChainWrapper {
	wrapper.Res = Max(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Find(cb Predicate) *ChainWrapper {
	wrapper.Res = Find(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Detect(cb Predicate) *ChainWrapper {
	return wrapper.Find(cb)
}

func (wrapper *ChainWrapper) FindLast(cb Predicate) *ChainWrapper {
	wrapper.Res = FindLast(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) FindIndex(cb Predicate) *ChainWrapper {
	wrapper.Res = FindIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) FindLastIndex(cb Predicate) *ChainWrapper {
	wrapper.Res = FindLastIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Some(cb Predicate) *ChainWrapper {
	wrapper.Res = Some(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Any(cb Predicate) *ChainWrapper {
	return wrapper.Some(cb)
}

func (wrapper *ChainWrapper) IndexOf(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	wrapper.Res = IndexOf(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) LastIndexOf(target Object, cb Comparator) *ChainWrapper {
	wrapper.Res = LastIndexOf(wrapper.Mid, target, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Contains(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	wrapper.Res = Contains(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Includes(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	return wrapper.Contains(target, isSorted, cb)
}

func (wrapper *ChainWrapper) Every(cb Predicate) *ChainWrapper {
	wrapper.Res = Every(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) All(cb Predicate) *ChainWrapper {
	return wrapper.Every(cb)
}

func (wrapper *ChainWrapper) Uniq(cb Comparator) *ChainWrapper {
	wrapper.Mid = Uniq(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Unique(cb Comparator) *ChainWrapper {
	return wrapper.Uniq(cb)
}

func (wrapper *ChainWrapper) Difference(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Difference(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Without(nonGrata Object, cb Comparator) *ChainWrapper {
	wrapper.Mid = Without(wrapper.Mid, nonGrata, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Intersection(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Intersection(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Union(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Union(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) SortBy(cb Comparator) *ChainWrapper {
	wrapper.Mid = SortBy(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) CountBy(cb Callback) *ChainWrapper {
	wrapper.Res = CountBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) GroupBy(cb Callback) *ChainWrapper {
	wrapper.Res = GroupBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Remove(pos int) *ChainWrapper {
	wrapper.Mid = Remove(wrapper.Mid, pos)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Insert(tg Object, pos int) *ChainWrapper {
	wrapper.Mid = Insert(wrapper.Mid, tg, pos)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Concat(next Seq) *ChainWrapper {
	wrapper.Mid = Concat(wrapper.Mid, next)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Shuffle() *ChainWrapper {
	wrapper.Mid = ShuffledCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) Reverse() *ChainWrapper {
	wrapper.Mid = ReversedCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return wrapper
}

func (wrapper *ChainWrapper) EqualsStrict(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Res = EqualsStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) EqualsNotStrict(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Res = EqualsNotStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return wrapper
}

func (wrapper *ChainWrapper) Value() Object {
	return wrapper.Res
}
