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

// Package ugo is a toolbox, inspired by underscore.js
// This package provide some of the underscore most used functions
//
// Usage:
//
//
//  package main
//
//  import (
//    u "github.com/alxrm/ugo"
//  )
//
//  func main() {
//	  strArr := u.Seq{ "nineteen", "three", "eleven", "five", "seventy", "six", "seven", "one" }
//
//	  lengths := u.Map(strArr, func(cur, _, _ u.Object) u.Object {
//	  	return len(cur.(string))
//	  })
//
//	  fmt.Println(lengths) // Output: [8 5 6 4 7 3 5 3]
//  }
package ugo

import (
	sorter "github.com/alxrm/ugo/timsort"
	"math"
	"math/rand"
	"reflect"
	"time"
)

// Object is an alias type for interface{}
type Object interface{}

// Seq is an alias type for (generic) interface{} slice e. g. []interface{}
type Seq []interface{}

// Collector is an alias type for function, used in Reduce based methods, which has following args:
//
// * Object memo
//
// * Object current
//
// * int index
//
// * Seq list
//
// * returns Object: memo, modified after some iteration
type Collector func(memo, current, currentKey, src Object) Object

// Callback is an alias type for function, used to get calculated result, which has following args:
//
// * Object current
//
// * int index
//
// * Seq list
//
// * returns Object: modified Seq element
type Callback func(current, currentKey, src Object) Object

// Comparator is an alias type for function, used to compare one value with another one, which has following args:
//
// * Object left
//
// * Object right
//
// * returns int: -1 for less, 0 for equals, 1 for larger
type Comparator func(left, right Object) int

// Predicate is an alias type for function, used to check value for some condition, which has following args:
//
// * Object current
//
// * int index
//
// * Seq list
//
// * returns bool: true if check has been passed
type Predicate func(current, currentKey, src Object) bool

// Action is an alias type for function, used to do some action based on given values, which has following args:
//
// * Object current
//
// * int index
//
// * Seq list
type Action func(current, currentKey, src Object)

const (
	toMin int = -1 /** constant value for incrementing */
	toMax int = 1  /** constant value for decrementing */
)

const (
	less   = -1
	equal  = 0
	larger = 1
)

// Chain method is a start point for chaining behaviour
// like that: u.Chain(Seq).Map(...).Filter(...).Reduce(...).Value()
func Chain(target Seq) *ChainWrapper {
	if IsEmpty(target) {
		target = Seq{}
	}
	return &ChainWrapper{Mid: target, Res: target}
}

// NewSeq creates a new Seq instance with given size
func NewSeq(size int) Seq {
	return make(Seq, size)
}

// Each Calls cb Action on each element
func Each(seq Seq, cb Action) {
	if cb == nil {
		return
	}
	for index, val := range seq {
		cb(val, index, seq)
	}
}

// ForEach is an alias for Each (see #Each)
func ForEach(seq Seq, cb Action) {
	Each(seq, cb)
}

// Map creates new slice same size, every element is the result of Callback
func Map(seq Seq, cb Callback) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return seq
	}

	length := len(seq)
	result := NewSeq(length)

	for index, val := range seq {
		result[index] = cb(val, index, seq)
	}

	return result
}

// Collect is an alias for Map (see #Map)
func Collect(seq Seq, cb Callback) Seq {
	return Map(seq, cb)
}

// Filter creates new slice, contains only elements that passed Predicate check
func Filter(seq Seq, cb Predicate) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return seq
	}

	result := NewSeq(0)

	for index, val := range seq {
		if cb(val, index, seq) {
			result = append(result, val)
		}
	}

	return result
}

// Select is an alias for Filter (see #Filter)
func Select(seq Seq, cb Predicate) Seq {
	return Filter(seq, cb)
}

// Reject creates new slice, contains only elements that haven't passed Predicate check
func Reject(seq Seq, cb Predicate) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return seq
	}

	return Filter(seq, negate(cb))
}

// Reduce makes single value from all of the slice elements, iterating from left
func Reduce(seq Seq, cb Collector, initial Object) Object {
	var memo Object

	if IsEmpty(seq) || cb == nil {
		return nil
	}

	length := len(seq) - 1

	if initial == nil {
		memo = seq[0]
		return createReduce(seq, cb, memo, 1, toMax, length-1)
	}

	memo = initial
	return createReduce(seq, cb, memo, 0, toMax, length)
}

// Inject is an alias for Reduce (see #Reduce)
func Inject(seq Seq, cb Collector, initial Object) Object {
	return Reduce(seq, cb, initial)
}

// FoldL is an alias for Reduce (see #Reduce)
func FoldL(seq Seq, cb Collector, initial Object) Object {
	return Reduce(seq, cb, initial)
}

// ReduceRight makes single value from all of the slice elements, iterating from right
func ReduceRight(seq Seq, cb Collector, initial Object) Object {
	var memo Object

	if IsEmpty(seq) || cb == nil {
		return nil
	}

	length := len(seq) - 1

	if initial == nil {
		memo = seq[length]
		return createReduce(seq, cb, memo, length-1, toMin, length-1)
	}

	memo = initial
	return createReduce(seq, cb, memo, length, toMin, length)
}

// FoldR is an alias for Reduce (see #ReduceRight)
func FoldR(seq Seq, cb Collector, initial Object) Object {
	return ReduceRight(seq, cb, initial)
}

// Min returns min value from slice, calculated in comparator
func Min(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, toMin, len(seq))
}

// Max returns max value from slice, calculated in comparator
func Max(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, toMax, len(seq))
}

// Find returns first found value, passed the predicate check
func Find(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, 0, toMax, length)
	return res
}

// Detect is an alias for Find (see #Find)
func Detect(seq Seq, cb Predicate) Object {
	return Find(seq, cb)
}

// FindLast returns last found value, passed the predicate check
func FindLast(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, length, toMin, length)
	return res
}

// FindIndex returns first found index, which value passed the predicate check
func FindIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, 0, toMax, length)
	return index
}

// FindLastIndex returns last found index, which value passed the predicate check
func FindLastIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, length, toMin, length)
	return index
}

// Some returns true if at least one element passed the predicate check
func Some(seq Seq, cb Predicate) bool {
	return FindIndex(seq, cb) != -1
}

// Any is an alias for Some (see #Some)
func Any(seq Seq, cb Predicate) bool {
	return Some(seq, cb)
}

// IndexOf founds index of the first element, which equals to passed one(target)
// NOTE: if slice is sorted, this method can use better search algorithm
func IndexOf(seq Seq, target Object, isSorted bool, cb Comparator) int {
	if cb == nil {
		return -1
	}

	if isSorted {
		return createBinarySearch(seq, target, cb, len(seq))
	}

	equalityPredicate := func(cur, _, _ Object) bool { return cb(cur, target) == 0 }
	return FindIndex(seq, equalityPredicate)
}

// LastIndexOf founds index of the last element, which equals to passed one(target)
func LastIndexOf(seq Seq, target Object, cb Comparator) int {
	if cb == nil {
		return -1
	}
	var equalityPredicate Predicate = func(cur, _, _ Object) bool { return cb(cur, target) == 0 }
	return FindLastIndex(seq, equalityPredicate)
}

// Contains returns true if slice contains element, which equals to passed one(target)
// NOTE: if slice is sorted, this method can use better search algorithm
func Contains(seq Seq, target Object, isSorted bool, cb Comparator) bool {
	if cb == nil {
		return false
	}

	return IndexOf(seq, target, isSorted, cb) != -1
}

// Includes is an alias for Contains (see #Contains)
func Includes(seq Seq, target Object, isSorted bool, cb Comparator) bool {
	return Contains(seq, target, isSorted, cb)
}

// Every returns true if every element in slice have passed the predicate test
func Every(seq Seq, cb Predicate) bool {
	if IsEmpty(seq) || cb == nil {
		return false
	}

	for index, val := range seq {
		if !cb(val, index, seq) {
			return false
		}
	}
	return true
}

// All is an alias for Every (see #Every)
func All(seq Seq, cb Predicate) bool {
	return Every(seq, cb)
}

// Uniq returns slice, which contains only unique elements, calculated by Comparator
func Uniq(seq Seq, cb Comparator) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return seq
	}

	result := NewSeq(0)
	for _, value := range seq {
		if !Contains(result, value, false, cb) {
			result = append(result, value)
		}
	}
	return result
}

// Unique is an alias for Uniq (see #Uniq)
func Unique(seq Seq, cb Comparator) Seq {
	return Uniq(seq, cb)
}

// Difference returns the values from slice that are not present in the other slice
func Difference(seq, other Seq, cb Comparator) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil || other == nil {
		return Seq{}
	}

	result := NewSeq(0)

	for _, value := range seq {
		if !Contains(other, value, false, cb) {
			result = append(result, value)
		}
	}
	return result
}

// Without returns the Slice without all instances of nonGrata value
func Without(seq Seq, nonGrata Object, cb Comparator) Seq {
	if seq == nil || cb == nil {
		return Seq{}
	}
	if nonGrata == nil {
		return seq
	}

	result := NewSeq(0)

	for _, value := range seq {
		if cb(value, nonGrata) != 0 {
			result = append(result, value)
		}
	}

	return result
}

// Intersection returns the values that are intersection of two slices
// Each value in the result is present in each of the arrays.
func Intersection(seq, other Seq, cb Comparator) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil || other == nil {
		return Seq{}
	}

	result := NewSeq(0)

	for _, value := range seq {
		if Contains(other, value, false, cb) {
			result = append(result, value)
		}
	}

	return Uniq(result, cb)
}

// Union returns the unique values that are union of two slices
// each value in the result appears at least once in one of the passed slices
func Union(seq, other Seq, cb Comparator) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return Seq{}
	}

	result := Concat(seq, other)

	return Uniq(result, cb)
}

// SortBy returns sorted slice, uses very powerful timsort* algorithm
// *timsort obtained from: https://github.com/psilva261/timsort
func SortBy(seq Seq, cb Comparator) Seq {
	if seq == nil {
		return Seq{}
	}
	if cb == nil {
		return seq
	}
	sorter.Sort(seq, lessThan(cb))
	return seq
}

// CountBy returns map, which values are count of certain kind of values,
// and keys are names of this kinds
func CountBy(seq Seq, cb Callback) (result map[string]int) {
	result = make(map[string]int, 0)
	key := ""

	if seq == nil || cb == nil {
		return result
	}

	for index, val := range seq {
		key = cb(val, index, seq).(string)

		if result[key] == 0 {
			result[key] = 1
		} else {
			result[key] = result[key] + 1
		}
	}

	return result
}

// GroupBy returns map, which keys are results of Callback calculation,
// and the value is the slice of elements, which gave such result
func GroupBy(seq Seq, cb Callback) map[Object]Seq {
	var key Object
	var length int

	result := make(map[Object]Seq, 0)

	if seq == nil || cb == nil {
		return result
	}

	for index, val := range seq {
		key = cb(val, index, seq)
		length = len(result[key])

		if length == 0 {
			result[key] = Seq{val}
		} else {
			result[key] = Insert(result[key], val, length)
		}
	}

	return result
}

// Remove takes an element from given position in slice
func Remove(seq Seq, position int) Seq {
	if IsEmpty(seq) {
		return Seq{}
	}
	position = fixPosition(position, len(seq)-1)

	bef := NewSeq(len(seq[:position]))
	aft := NewSeq(len(seq[position+1:]))

	copy(bef, seq[:position])
	copy(aft, seq[position+1:])

	return Concat(bef, aft)
}

// Insert pushes an element into given position in slice
func Insert(seq Seq, target Object, position int) Seq {
	if seq == nil {
		return Seq{}
	}
	position = fixPosition(position, len(seq))

	seq = append(seq, 0)
	copy(seq[position+1:], seq[position:])
	seq[position] = target

	return seq
}

// Concat adds another slice to the end of given slice
func Concat(seq, next Seq) Seq {
	if seq == nil {
		return Seq{}
	}
	if next == nil {
		return seq
	}

	return append(seq, next...)
}

// Shuffle returns shuffled slice
func Shuffle(seq Seq) Seq {
	if seq == nil {
		return Seq{}
	}
	return createShuffle(seq)
}

// ShuffledCopy returns shuffled copy of slice
func ShuffledCopy(seq Seq) Seq {
	if seq == nil {
		return Seq{}
	}

	length := len(seq)
	copied := NewSeq(length)
	copy(copied, seq)

	return createShuffle(copied)
}

// Reverse returns reversed slice
func Reverse(seq Seq) Seq {
	if seq == nil {
		return Seq{}
	}

	length := len(seq)

	return createReverse(seq, length)
}

// ReversedCopy returns reversed copy of slice
func ReversedCopy(seq Seq) Seq {
	if seq == nil {
		return Seq{}
	}

	length := len(seq)
	copied := NewSeq(length)
	copy(copied, seq)

	return createReverse(copied, length)
}

// EqualsStrict checks whether both of the given slices are strictly equal,
// e. g. they got the same values in the same positions
func EqualsStrict(seqLeft, seqRight Seq, cb Comparator) bool {
	if len(seqLeft) != len(seqRight) || cb == nil {
		return false
	}

	for index, value := range seqLeft {
		if cb(seqRight[index], value) != 0 {
			return false
		}
	}

	return true
}

// EqualsNotStrict checks whether both of the given slices are equal, but not strictly,
// e. g. they the same values, but positions can be different
func EqualsNotStrict(seqLeft, seqRight Seq, cb Comparator) bool {
	lengthLeft := len(seqLeft)
	lengthRight := len(seqRight)
	target := NewSeq(lengthLeft)
	foundIndex := -1

	if lengthLeft != lengthRight || cb == nil {
		return false
	}

	copy(target, seqRight)

	for _, value := range seqLeft {
		if foundIndex = IndexOf(target, value, false, cb); foundIndex != -1 {
			target = Remove(target, foundIndex)
		} else {
			return false
		}
	}

	return true
}

/* Utils */

// Random returns the random number in given range
func Random(min, max float64) int {
	if min == max {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	min = fixNumber(min)
	max = fixNumber(max)

	if min > max {
		min, max = max, min
	}

	res := min + math.Floor(rand.Float64()*(max-min+1))
	if math.IsNaN(res) {
		return 0
	}

	return int(res)
}

// IsEmpty returns true if given slice has zero length or it's to nil
func IsEmpty(seq Seq) bool {
	return seq == nil || len(seq) == 0
}

// IsSlice returns true if given object has type of slice
func IsSlice(target Object) bool {
	return reflect.ValueOf(target).Kind() == reflect.Slice
}

// From returns Seq from given object, filling the resulting Seq via reflection
// NOTE: You should possibly avoid using this one
func From(target Object, size int) Seq {
	if !IsSlice(target) || size < 0 {
		return Seq{}
	}

	s := reflect.ValueOf(target)
	res := make(Seq, size)

	for i := 0; i < size; i++ {
		res[i] = s.Index(i).Interface()
	}

	return res
}

/* private methods */
// createComparingIterator returns single value, which conforms some condition (dir)
// it works for O(n/2) and uses in Min/Max methods
func createComparingIterator(seq Seq, cb Comparator, dir, length int) Object {
	var lastComputed Object
	var innerComputed Object

	if IsEmpty(seq) || cb == nil {
		return -1
	}
	if length == 1 {
		return seq[0]
	}

	lastComputed = seq[0]

	for left, right := 0, length-1; left <= right; left, right = left+1, right-1 {
		if sgn(cb(seq[left], seq[right])) == dir {
			innerComputed = seq[left]
		} else {
			innerComputed = seq[right]
		}

		if sgn(cb(innerComputed, lastComputed)) == dir {
			lastComputed = innerComputed
		}
	}

	return lastComputed
}

// createShuffle returns single value folded to it from given slice
func createReduce(seq Seq, cb Collector, memo Object, startPoint, direction, length int) Object {
	result := memo
	index := startPoint

	for i := 0; i <= length; i++ {
		result = cb(result, seq[index], index, seq)
		index += direction
	}

	return result
}

// createShuffle returns slice shuffled by Fisher-Yates algorithm
func createShuffle(seq Seq) Seq {
	for i := range seq {
		j := Random(0, float64(i))
		seq[i], seq[j] = seq[j], seq[i]
	}

	return seq
}

// createReverse returns the slice, shuffled in O(n/2) operations
func createReverse(seq Seq, length int) Seq {
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		seq[left], seq[right] = seq[right], seq[left]
	}

	return seq
}

// createBinarySearch returns the index of the value we want to find,
// it just goes through given slice and invokes the Predicate check
func createPredicateSearch(seq Seq, cb Predicate, startPoint, direction, length int) (res Object, resIndex int) {
	res = nil
	resIndex = -1

	if seq == nil || cb == nil {
		return
	}

	index := startPoint

	for i := 0; i <= length; i++ {
		if cb(seq[index], index, seq) {
			res = seq[index]
			resIndex = index
			break
		}
		index += direction
	}

	return
}

// createBinarySearch returns the index of the value we want to find,
// it uses the Binary Search algorithm to reduce iteration count
// this assumes that we operating with sorted Sequence
func createBinarySearch(sortedSeq Seq, target Object, cb Comparator, length int) int {
	resIndex := -1

	lo := 0
	hi := length

	for lo < hi {
		mid := (lo + hi) >> 1

		if cb(sortedSeq[mid], target) == 0 {
			resIndex = mid
			break
		}

		if cb(sortedSeq[mid], target) < 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return resIndex
}

// lessThen returns function we can use for sorting or comparing,
// it simply returns true if the left value is less than the right value
func lessThan(cb Comparator) func(l, r interface{}) bool {
	return func(l, r interface{}) bool { return cb(l, r) < 0 }
}

// negate returns the given Predicate but with opposite result
func negate(cb Predicate) Predicate {
	return func(cur, index, list Object) bool { return !cb(cur, index, list) }
}

// sgn returns the sign of the passed number, which can be, as follows,
// -1, 0, 1 (negative, zero, positive)
func sgn(num int) int {
	if num < 0 {
		return less
	} else if num == 0 {
		return equal
	}

	return larger
}

// fixPosition returns robust index, that is inside the slice bounds
func fixPosition(pos, ableMax int) int {
	if pos < 0 {
		return 0
	} else if pos > ableMax {
		return ableMax
	}

	return pos
}

// fixNumber returns fixed float64 number we able to use in calculations
func fixNumber(num float64) float64 {
	if math.IsNaN(num) {
		return 0
	} else if math.IsInf(num, -1) {
		return math.MinInt64
	} else if math.IsInf(num, 1) {
		return math.MaxInt64
	}

	return num
}
