/*
The MIT License (MIT)

Copyright (c) 2016 Alexey Derbyshev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


// Package ugo is a toolbox, inspired by underscore.js
// This package provide some of the underscore most used functions
//
// Usage
//
//  package main
//
//  import (
//  	u "github.com/alxrm/ugo"
//  )
//
//  func main() {
//		strArr := u.Seq{ "nineteen", "three", "eleven", "five", "seventy", "six", "seven", "one" }
//
//		lengths := u.Map(strArr, func(cur, _, _ u.Object) u.Object {
//			return len(cur.(string))
//		})
//
//		fmt.Println(lengths) // Output: [8 5 6 4 7 3 5 3]
//  }
package ugo

import (
	sorter "github.com/alxrm/ugo/timsort"
	"math"
	"math/rand"
	"reflect"
	"time"
)

/**
 An alias type for interface{}
*/
type Object interface{}

/**
 An alias type for (generic) interface{} slice
*/
type Seq []interface{}


/**
 An alias type for Callback function, which has following args:

 - Object memo
 - Object current
 - int index
 - Seq list

 - returns Object: memo, modified after some iteration
*/
type Collector func(memo, current, currentKey, src Object) Object

/**
 An alias type for Callback function, which has following args:

 - Object current
 - int index
 - Seq list

 - returns Object: modified Seq element
*/
type Callback func(current, currentKey, src Object) Object

/**
 An alias type for Predicate function, which has following args:

 - Object left
 - Object right

 - returns int: -1 for less, 0 for equals, 1 for larger
*/
type Comparator func(left, right Object) int

/**
 An alias type for Predicate function, which has following args:

 - Object current
 - int index
 - Seq list

 - returns bool: true if check has been passed
*/
type Predicate func(current, currentKey, src Object) bool

/**
 An alias type for Action function, which has following args:

 - Object current
 - int index
 - Seq list
*/
type Action func(current, currentKey, src Object)


const (
	_DIRECTION_TO_MIN int = -1 /** constant value for incrementing */
	_DIRECTION_TO_MAX int = 1 /** constant value for decrementing */
)

const (
	_LESS = -1
	_EQUAL = 0
	_LARGER = 1
)

/**
 creates new Seq aliased slice with given size
*/
func NewSeq(size int) Seq {
	return make(Seq, size)
}

/**
 Calls cb Action on each element
*/
func Each(seq Seq, cb Action) {
	if cb == nil { return }
	for index, val := range seq {
		cb(val, index, seq)
	}
}

/**
 Creates new slice same size, every element is the result of Callback
*/
func Map(seq Seq, cb Callback) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }

	length := len(seq)
	result := NewSeq(length)

	for index, val := range seq {
		result[index] = cb(val, index, seq)
	}

	return result
}

/**
 Creates new slice, contains only elements that passed Predicate check
*/
func Filter(seq Seq, cb Predicate) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }

	result := NewSeq(0)

	for index, val := range seq {
		if cb(val, index, seq) {
			result = append(result, val)
		}
	}

	return result
}

/**
 Creates new slice, contains only elements that h passed Predicate check
*/
func Reject(seq Seq, cb Predicate) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }

	return Filter(seq, negate(cb))
}

/**
 Makes single value from all of the slice elements, iterating from left
*/
func Reduce(seq Seq, cb Collector, initial Object) Object {
	var memo Object = nil

	if IsEmpty(seq) || cb == nil {
		return nil
	}

	length := len(seq) - 1

	if initial == nil {
		memo = seq[0]
		return createReduce(seq, cb, memo, 1, _DIRECTION_TO_MAX, length-1)
	} else {
		memo = initial
		return createReduce(seq, cb, memo, 0, _DIRECTION_TO_MAX, length)
	}
}

/**
 Makes single value from all of the slice elements, iterating from right
*/
func ReduceRight(seq Seq, cb Collector, initial Object) Object {
	var memo Object = nil

	if IsEmpty(seq) || cb == nil {
		return nil
	}

	length := len(seq) - 1

	if initial == nil {
		memo = seq[length]
		return createReduce(seq, cb, memo, length-1, _DIRECTION_TO_MIN, length-1)
	} else {
		memo = initial
		return createReduce(seq, cb, memo, length, _DIRECTION_TO_MIN, length)
	}
}

/**
 returns min value from slice, calculated in comparator
*/
func Min(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, _DIRECTION_TO_MIN, len(seq))
}

/**
 returns max value from slice, calculated in comparator
*/
func Max(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, _DIRECTION_TO_MAX, len(seq))
}

/**
 returns first found value, passed the predicate check
*/
func Find(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, 0, _DIRECTION_TO_MAX, length)
	return res
}

/**
 returns last found value, passed the predicate check
*/
func FindLast(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, length, _DIRECTION_TO_MIN, length)
	return res
}

/**
 returns first found index, which value passed the predicate check
*/
func FindIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, 0, _DIRECTION_TO_MAX, length)
	return index
}

/**
 returns last found index, which value passed the predicate check
*/
func FindLastIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, length, _DIRECTION_TO_MIN, length)
	return index
}

/**
 returns true if at least one element passed the predicate check
*/
func Some(seq Seq, cb Predicate) bool {
	return FindIndex(seq, cb) != -1
}

/**
 founds index of the first element, which equals to passed one(target)
 NOTE: if slice is sorted, this method can use better search algorithm
*/
func IndexOf(seq Seq, target Object, isSorted bool, cb Comparator) int {
	if cb == nil { return -1 }

	if isSorted {
		_, index := createBinarySearch(seq, target, cb, len(seq))
		return index
	} else {
		var equalityPredicate Predicate = func(cur, _, _ Object) bool { return cb(cur, target) == 0 }
		return FindIndex(seq, equalityPredicate)
	}
}

/**
 founds index of the last element, which equals to passed one(target)
*/
func LastIndexOf(seq Seq, target Object, cb Comparator) int {
	if cb == nil { return -1 }
	var equalityPredicate Predicate = func(cur, _, _ Object) bool { return cb(cur, target) == 0 }
	return FindLastIndex(seq, equalityPredicate)
}

/**
 returns true if slice contains element, which equals to passed one(target)
 NOTE: if slice is sorted, this method can use better search algorithm
*/
func Contains(seq Seq, target Object, isSorted bool, cb Comparator) bool {
	if cb == nil { return false }

	return IndexOf(seq, target, isSorted, cb) != -1
}

/**
 returns true if every element in slice have passed the predicate test
*/
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

/**
 returns slice, which contains only unique elements, calculated by Comparator
*/
func Uniq(seq Seq, cb Comparator) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }

	result := NewSeq(0)
	for _, value := range seq {
		if !Contains(result, value, false, cb) {
			result = append(result, value)
		}
	}
	return result
}

/**
 returns the values from slice that are not present in the other slice
*/
func Difference(seq, other Seq, cb Comparator) Seq {
	if seq == nil { return Seq{} }
	if cb == nil || other == nil { return Seq{} }

	result := NewSeq(0)

	for _, value := range seq {
		if !Contains(other, value, false, cb) {
			result = append(result, value)
		}
	}
	return result
}

/**
 returns the Slice without all instances of nonGrata value
*/
func Without(seq Seq, nonGrata Object, cb Comparator) Seq {
	if seq == nil || cb == nil { return Seq{} }
	if nonGrata == nil { return seq }

	result := NewSeq(0)

	for _, value := range seq {
		if cb(value, nonGrata) != 0 {
			result = append(result, value)
		}
	}

	return result
}

/**
 returns the values that are intersection of two slices
 Each value in the result is present in each of the arrays.
*/
func Intersection(seq, other Seq, cb Comparator) Seq {
	if seq == nil { return Seq{} }
	if cb == nil || other == nil { return Seq{} }

	result := NewSeq(0)

	for _, value := range seq {
		if Contains(other, value, false, cb) {
			result = append(result, value)
		}
	}
	return Uniq(result, cb);
}

/**
 returns the unique values that are union of two slices
 each value in the result appears at least once in one of the passed slices
*/
func Union(seq, other Seq, cb Comparator) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return Seq{} }

	result := Concat(seq, other)

	return Uniq(result, cb)
}

/**
 returns sorted slice, uses very powerful timsort* algorithm
 *timsort obtained from: https://github.com/psilva261/timsort
*/
func SortBy(seq Seq, cb Comparator) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }
	sorter.Sort(seq, lessThan(cb))
	return seq
}

/**
 returns map, which values are count of certain kind of values,
 and keys are names of this kinds
*/
func CountBy(seq Seq, cb Callback) (result map[string]int) {
	if seq == nil || cb == nil {
		return nil
	}

	result = make(map[string]int, 0)
	key := ""

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

/**
 removes an element from given position in slice
*/
func Remove(seq Seq, position int) Seq {
	if seq == nil { return Seq{} }
	position = fixPosition(position, len(seq) - 1)

	return append(seq[:position], seq[position + 1:]...)
}

/**
 inserts an element into given position in slice
*/
func Insert(seq Seq, target Object, position int) Seq {
	if seq == nil { return Seq{} }
	position = fixPosition(position, len(seq))

	seq = append(seq, 0)
	copy(seq[position + 1:], seq[position:])
	seq[position] = target

	return seq
}

/**
 adds another slice to the end of given slice
*/
func Concat(seq, next Seq) Seq {
	if seq == nil { return Seq{} }
	if next == nil { return seq }

	return append(seq, next...)
}

/**
 returns shuffled slice
*/
func Shuffle(seq Seq) Seq {
	if seq == nil { return Seq{} }
	return createShuffle(seq)
}

/**
 returns shuffled copy of slice
*/
func ShuffledCopy(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)
	copied := NewSeq(length)
	copy(copied, seq)

	return createShuffle(copied)
}

/**
 returns reversed slice
*/
func Reverse(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)

	return createReverse(seq, length)
}

/**
 returns reversed copy of slice
*/
func ReversedCopy(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)
	copied := NewSeq(length)
	copy(copied, seq)

	return createReverse(copied, length)
}

/**
 checks whether both of the given slices are strictly equal,
 e. g. they got the same values in the same positions
*/
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

/**
 checks whether both of the given slices are equal, but not strictly,
 e. g. they the same values, but positions can be different
*/
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

/**
 return the random number in given range
*/
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

/**
 returns true if given slice has zero length or it's to nil
*/
func IsEmpty(seq Seq) bool {
	return seq == nil || len(seq) == 0
}

/**
 returns true if given object has type of slice
*/
func IsSlice(target Object) bool {
	return reflect.ValueOf(target).Kind() == reflect.Slice
}

/**
 returns Seq from given object, filling the resulting Seq via reflection
 NOTE: You should possibly avoid using this one
*/
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
func createComparingIterator(seq Seq, cb Comparator, dir, length int) Object {
	var lastComputed Object = nil
	var innerComputed Object = nil

	if IsEmpty(seq) || cb == nil { return -1 }
	if length == 1 { return seq[0] }

	lastComputed = seq[0]

	for left, right := 0, length - 1; left < right; left, right = left + 1, right - 1 {
		if sgn(cb(seq[left], seq[right])) == _LESS {
			innerComputed = seq[right]
		} else {
			innerComputed = seq[left]
		}

		if sgn(cb(innerComputed, lastComputed)) == dir {
			lastComputed = innerComputed
		}
	}

	return lastComputed
}

func createReduce(seq Seq, cb Collector, memo Object, startPoint, direction, length int) Object {
	result := memo
	index := startPoint

	for i := 0; i <= length; i++ {
		result = cb(result, seq[index], index, seq)
		index += direction
	}

	return result
}

func createShuffle(seq Seq) Seq {
	for i := range seq {
		j := Random(0, float64(i))
		seq[i], seq[j] = seq[j], seq[i]
	}

	return seq
}

func createReverse(seq Seq, length int) Seq {
	for left, right := 0, length - 1; left < right; left, right = left + 1, right - 1 {
		seq[left], seq[right] = seq[right], seq[left]
	}

	return seq
}

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

/* this assumes that we operating with sorted Sequence */
func createBinarySearch(sortedSeq Seq, target Object, cb Comparator, length int) (res Object, resIndex int) {
	res = nil
	resIndex = -1

	lo := 0
	hi := length

	for lo < hi {
		mid := (lo + hi) >> 1

		if cb(sortedSeq[mid], target) == 0 {
			resIndex = mid
			res = sortedSeq[mid]
			break
		}

		if cb(sortedSeq[mid], target) < 0 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return
}

func lessThan(cb Comparator) func(l, r interface{}) bool {
	return func(l, r interface{}) bool { return cb(l, r) < 0 }
}

func negate(cb Predicate) Predicate {
	return func(cur, index, list Object) bool { return !cb(cur, index, list) }
}

func sgn(num int) int {
	if num < 0 {
		return _LESS
	} else if num == 0 {
		return _EQUAL
	} else {
		return _LARGER
	}
}

func fixPosition(pos, ableMax int) int {
	if pos < 0 {
		return 0
	} else if pos > ableMax {
		return ableMax
	} else {
		return pos
	}
}

func fixNumber(num float64) float64 {
	if math.IsNaN(num) {
		return 0
	} else if math.IsInf(num, -1) {
		return math.MinInt64
	} else if math.IsInf(num, 1) {
		return math.MaxInt64
	} else {
		return num
	}
}
