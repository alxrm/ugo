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

// Each is a chaining wrapper for #Each
func (wrapper *ChainWrapper) Each(cb Action) *ChainWrapper {
	Each(wrapper.Mid, cb)
	return wrapper
}

// ForEach is a chaining wrapper for #ForEach
func (wrapper *ChainWrapper) ForEach(cb Action) *ChainWrapper {
	return wrapper.Each(cb)
}

// Map is a chaining wrapper for #Map
func (wrapper *ChainWrapper) Map(cb Callback) *ChainWrapper {
	wrapper.Mid = Map(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Collect is a chaining wrapper for #Collect
func (wrapper *ChainWrapper) Collect(cb Callback) *ChainWrapper {
	return wrapper.Map(cb)
}

// Filter is a chaining wrapper for #Filter
func (wrapper *ChainWrapper) Filter(cb Predicate) *ChainWrapper {
	wrapper.Mid = Filter(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Select is a chaining wrapper for #Select
func (wrapper *ChainWrapper) Select(cb Predicate) *ChainWrapper {
	return wrapper.Filter(cb)
}

// Reject is a chaining wrapper for #Reject
func (wrapper *ChainWrapper) Reject(cb Predicate) *ChainWrapper {
	wrapper.Mid = Reject(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Reduce is a chaining wrapper for #Reduce
func (wrapper *ChainWrapper) Reduce(cb Collector, initial Object) *ChainWrapper {
	wrapper.Res = Reduce(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return wrapper
}

// Inject is a chaining wrapper for #Inject
func (wrapper *ChainWrapper) Inject(cb Collector, initial Object) *ChainWrapper {
	return wrapper.Reduce(cb, initial)
}

// FoldL is a chaining wrapper for #FoldL
func (wrapper *ChainWrapper) FoldL(cb Collector, initial Object) *ChainWrapper {
	return wrapper.Reduce(cb, initial)
}

// ReduceRight is a chaining wrapper for #ReduceRight
func (wrapper *ChainWrapper) ReduceRight(cb Collector, initial Object) *ChainWrapper {
	wrapper.Res = ReduceRight(wrapper.Mid, cb, initial)
	wrapper.Mid = nil
	return wrapper
}

// FoldR is a chaining wrapper for #FoldR
func (wrapper *ChainWrapper) FoldR(cb Collector, initial Object) *ChainWrapper {
	return wrapper.ReduceRight(cb, initial)
}

// Min is a chaining wrapper for #Min
func (wrapper *ChainWrapper) Min(cb Comparator) *ChainWrapper {
	wrapper.Res = Min(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Max is a chaining wrapper for #Max
func (wrapper *ChainWrapper) Max(cb Comparator) *ChainWrapper {
	wrapper.Res = Max(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Find is a chaining wrapper for #Find
func (wrapper *ChainWrapper) Find(cb Predicate) *ChainWrapper {
	wrapper.Res = Find(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Detect is a chaining wrapper for #Detect
func (wrapper *ChainWrapper) Detect(cb Predicate) *ChainWrapper {
	return wrapper.Find(cb)
}

// FindLast is a chaining wrapper for #FindLast
func (wrapper *ChainWrapper) FindLast(cb Predicate) *ChainWrapper {
	wrapper.Res = FindLast(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// FindIndex is a chaining wrapper for #FindIndex
func (wrapper *ChainWrapper) FindIndex(cb Predicate) *ChainWrapper {
	wrapper.Res = FindIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// FindLastIndex is a chaining wrapper for #FindLastIndex
func (wrapper *ChainWrapper) FindLastIndex(cb Predicate) *ChainWrapper {
	wrapper.Res = FindLastIndex(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Some is a chaining wrapper for #Some
func (wrapper *ChainWrapper) Some(cb Predicate) *ChainWrapper {
	wrapper.Res = Some(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Any is a chaining wrapper for #Any
func (wrapper *ChainWrapper) Any(cb Predicate) *ChainWrapper {
	return wrapper.Some(cb)
}

// IndexOf is a chaining wrapper for #IndexOf
func (wrapper *ChainWrapper) IndexOf(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	wrapper.Res = IndexOf(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return wrapper
}

// LastIndexOf is a chaining wrapper for #LastIndexOf
func (wrapper *ChainWrapper) LastIndexOf(target Object, cb Comparator) *ChainWrapper {
	wrapper.Res = LastIndexOf(wrapper.Mid, target, cb)
	wrapper.Mid = nil
	return wrapper
}

// Contains is a chaining wrapper for #Contains
func (wrapper *ChainWrapper) Contains(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	wrapper.Res = Contains(wrapper.Mid, target, isSorted, cb)
	wrapper.Mid = nil
	return wrapper
}

// Includes is a chaining wrapper for #Includes
func (wrapper *ChainWrapper) Includes(target Object, isSorted bool, cb Comparator) *ChainWrapper {
	return wrapper.Contains(target, isSorted, cb)
}

// Every is a chaining wrapper for #Every
func (wrapper *ChainWrapper) Every(cb Predicate) *ChainWrapper {
	wrapper.Res = Every(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// All is a chaining wrapper for #All
func (wrapper *ChainWrapper) All(cb Predicate) *ChainWrapper {
	return wrapper.Every(cb)
}

// Uniq is a chaining wrapper for #Uniq
func (wrapper *ChainWrapper) Uniq(cb Comparator) *ChainWrapper {
	wrapper.Mid = Uniq(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Unique is a chaining wrapper for #Unique
func (wrapper *ChainWrapper) Unique(cb Comparator) *ChainWrapper {
	return wrapper.Uniq(cb)
}

// Differenc is a chaining wrapper for #Differenc
func (wrapper *ChainWrapper) Difference(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Difference(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Without is a chaining wrapper for #Without
func (wrapper *ChainWrapper) Without(nonGrata Object, cb Comparator) *ChainWrapper {
	wrapper.Mid = Without(wrapper.Mid, nonGrata, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Intersection is a chaining wrapper for #Intersection
func (wrapper *ChainWrapper) Intersection(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Intersection(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Union is a chaining wrapper for #Union
func (wrapper *ChainWrapper) Union(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Mid = Union(wrapper.Mid, other, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// SortBy is a chaining wrapper for #SortBy
func (wrapper *ChainWrapper) SortBy(cb Comparator) *ChainWrapper {
	wrapper.Mid = SortBy(wrapper.Mid, cb)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// CountBy is a chaining wrapper for #CountBy
func (wrapper *ChainWrapper) CountBy(cb Callback) *ChainWrapper {
	wrapper.Res = CountBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// GroupBy is a chaining wrapper for #GroupBy
func (wrapper *ChainWrapper) GroupBy(cb Callback) *ChainWrapper {
	wrapper.Res = GroupBy(wrapper.Mid, cb)
	wrapper.Mid = nil
	return wrapper
}

// Remove is a chaining wrapper for #Remove
func (wrapper *ChainWrapper) Remove(pos int) *ChainWrapper {
	wrapper.Mid = Remove(wrapper.Mid, pos)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Insert is a chaining wrapper for #Insert
func (wrapper *ChainWrapper) Insert(tg Object, pos int) *ChainWrapper {
	wrapper.Mid = Insert(wrapper.Mid, tg, pos)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Concat is a chaining wrapper for #Concat
func (wrapper *ChainWrapper) Concat(next Seq) *ChainWrapper {
	wrapper.Mid = Concat(wrapper.Mid, next)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Shuffle is a chaining wrapper for #Shuffle
func (wrapper *ChainWrapper) Shuffle() *ChainWrapper {
	wrapper.Mid = ShuffledCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// Reverse is a chaining wrapper for #Reverse
func (wrapper *ChainWrapper) Reverse() *ChainWrapper {
	wrapper.Mid = ReversedCopy(wrapper.Mid)
	wrapper.Res = wrapper.Mid
	return wrapper
}

// EqualsStrict is a chaining wrapper for #EqualsStrict
func (wrapper *ChainWrapper) EqualsStrict(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Res = EqualsStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return wrapper
}

// EqualsNotStrict is a chaining wrapper for #EqualsNotStrict
func (wrapper *ChainWrapper) EqualsNotStrict(other Seq, cb Comparator) *ChainWrapper {
	wrapper.Res = EqualsNotStrict(wrapper.Mid, other, cb)
	wrapper.Mid = nil
	return wrapper
}

// Value returns result of calculations, you've done through chaining calls
func (wrapper *ChainWrapper) Value() Object {
	return wrapper.Res
}
