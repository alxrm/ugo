package ugo

import (
	sorter "github.com/alxrm/ugo/timsort"
	"math"
	"math/rand"
	"reflect"
	"time"
)

type Object interface{}
type Seq []interface{}

type Collector func(memo, current, currentKey, src Object) Object
type Callback func(current, currentKey, src Object) Object
type Comparator func(left, right Object) int
type Predicate func(current, currentKey, src Object) bool
type Action func(current, currentKey, src Object)

const _DIRECTION_TO_MAX int = 1
const _DIRECTION_TO_MIN int = -1

/* Slice methods */
func NewSeq(size int) Seq {
	return make(Seq, size)
}

/**
Calls cb Action on each element

@param {Seq} seq
@param {Action} cb

Action(cur, index, list Object)
	@param {Object} current
	@param {int} index
	@param {Seq} list
*/
func Each(seq Seq, cb Action) {
	if cb == nil { return }
	for index, val := range seq {
		cb(val, index, seq)
	}
}

/**
Creates new slice same size, every element is the result of Callback

@param {Seq} seq
@param {Callback} cb
@return {Seq} mutated slice

Callback(cur, index, list Object) Object
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {Object} mutated element
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

@param {Seq} seq
@param {Predicate} cb
@return {Seq} filtered slice

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
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

@param {Seq} seq
@param {Predicate} cb
@return {Seq} reverse filtered slice

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func Reject(seq Seq, cb Predicate) Seq {
	if seq == nil { return Seq{} }
	if cb == nil { return seq }

	return Filter(seq, negate(cb))
}

/**
Makes single value from all of the slice elements, iterating from left

@param {Seq} seq
@param {Collector} cb
@param {Object|nil} initial
@return {Object} memo, collected all the elements in one

Collector(memo, cur, index, list Object) bool
	@param {Object} memo
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {Object} memo, updated with new element
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

@param {Seq} seq
@param {Collector} cb
@param {Object|nil} initial
@return {Object} memo, collected all the elements in one

Collector(memo, cur, index, list Object) bool
	@param {Object} memo
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {Object} memo, updated with new element
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

@param {Seq} seq
@param {Comparator} cb
@return {Object} min element

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
*/
func Min(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, _DIRECTION_TO_MIN, len(seq))
}

/**
returns max value from slice, calculated in comparator

@param {Seq} seq
@param {Comparator} cb
@return {Object} max element

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
*/
func Max(seq Seq, cb Comparator) Object {
	return createComparingIterator(seq, cb, _DIRECTION_TO_MAX, len(seq))
}

/**
returns first found value, passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {Object} found element

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func Find(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, 0, _DIRECTION_TO_MAX, length)
	return res
}

/**
returns last found value, passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {Object} last found element

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func FindLast(seq Seq, cb Predicate) Object {
	length := len(seq) - 1
	res, _ := createPredicateSearch(seq, cb, length, _DIRECTION_TO_MIN, length)
	return res
}

/**
returns first found index, which value passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {int} found index

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func FindIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, 0, _DIRECTION_TO_MAX, length)
	return index
}

/**
returns last found index, which value passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {int} last found index

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func FindLastIndex(seq Seq, cb Predicate) int {
	length := len(seq) - 1
	_, index := createPredicateSearch(seq, cb, length, _DIRECTION_TO_MIN, length)
	return index
}

/**
returns true if at least one element passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {bool} whether some of the elements passed the predicate test

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
*/
func Some(seq Seq, cb Predicate) bool {
	return FindIndex(seq, cb) != -1
}

/**
founds index of the first element, which equals to passed one(target)
NOTE: if slice is sorted, this method can use better search algorithm

@param {Seq} seq
@param {Object} target
@param {bool} isSorted
@param {Comparator} cb
@return {int} index of found element

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Object} target
@param {bool} isSorted
@param {Comparator} cb
@return {int} index of found element

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
*/
func LastIndexOf(seq Seq, target Object, cb Comparator) int {
	if cb == nil { return -1 }
	var equalityPredicate Predicate = func(cur, _, _ Object) bool { return cb(cur, target) == 0 }
	return FindLastIndex(seq, equalityPredicate)
}

/**
returns true if slice contains element, which equals to passed one(target)
NOTE: if slice is sorted, this method can use better search algorithm

@param {Seq} seq
@param {Object} target
@param {bool} isSorted
@param {Comparator} cb
@return {bool} whether the slice contains target element

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
*/
func Contains(seq Seq, target Object, isSorted bool, cb Comparator) bool {
	if cb == nil { return false }

	return IndexOf(seq, target, isSorted, cb) != -1
}

/**
returns true if every element in slice have passed the predicate test

@param {Seq} seq
@param {Predicate} cb
@return {bool} whether all of the elements passed the predicate test

Predicate(cur, index, list Object) bool
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {bool} whether element passed the Predicate check
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

@param {Seq} seq
@param {Comparator} cb
@return {Seq} slice, which contains only unique elements


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Seq} other
@param {Comparator} cb
@return {Seq} resulting slice


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Object} nonGrata
@param {Comparator} cb
@return {Seq} result


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Seq} other
@param {Comparator} cb
@return {Seq} resulting intersection of slices


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Seq} other
@param {Comparator} cb
@return {Seq} resulting intersection of slices


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Comparator} cb
@return {Seq} resulting sorted slice

Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} seq
@param {Callback} cb
@return {map[string]int} resulting map

Callback(cur, index, list Object) Object
	@param {Object} current
	@param {int} index
	@param {Seq} list
	@return {Object(string)} name of countable values
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

@param {Seq} seq
@param {int} position
@return {Seq} slice, without removed element

*/
func Remove(seq Seq, position int) Seq {
	if seq == nil { return Seq{} }
	position = fixPosition(position, len(seq) - 1)

	return append(seq[:position], seq[position + 1:]...)
}

/**
inserts an element into given position in slice

@param {Seq} seq
@param {Object} target
@param {int} position
@return {Seq} slice, with new inserted element (target)

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
concates another slice to the end of given slice

@param {Seq} seq
@param {Seq} next
@return {Seq} concatenated slice

*/
func Concat(seq, next Seq) Seq {
	if seq == nil { return Seq{} }
	if next == nil { return seq }

	return append(seq, next...)
}

/**
returns shuffled slice

@param {Seq} seq
@return {Seq} shuffled

*/
func Shuffle(seq Seq) Seq {
	if seq == nil { return Seq{} }

	for i := range seq {
		j := Random(0, float64(i))
		seq[i], seq[j] = seq[j], seq[i]
	}

	return seq
}

/**
returns shuffled copy of slice

@param {Seq} seq
@return {Seq} shuffled copy

*/
func ShuffledCopy(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)
	shuffled := NewSeq(length)
	copy(shuffled, seq)

	for i := range seq {
		j := Random(0, float64(i))
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}

/**
returns reversed slice

@param {Seq} seq
@return {Seq} reversed

*/
func Reverse(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)

	for left, right := 0, length - 1; left < right; left, right = left + 1, right - 1 {
		seq[left], seq[right] = seq[right], seq[left]
	}

	return seq
}

/**
returns reversed copy of slice

@param {Seq} seq
@return {Seq} reversed copy

*/
func ReversedCopy(seq Seq) Seq {
	if seq == nil { return Seq{} }

	length := len(seq)
	reversed := NewSeq(length)
	copy(reversed, seq)

	for left, right := 0, length - 1; left < right; left, right = left + 1, right - 1 {
		reversed[left], reversed[right] = reversed[right], reversed[left]
	}

	return reversed
}

/**
checks whether both of the given slices are strictly equal,
e. g. they got the same values in the same positions

@param {Seq} leftSeq
@param {Seq} rightSeq
@param {Comparator} cb
@return {boo} whether they are equals or not


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {Seq} leftSeq
@param {Seq} rightSeq
@param {Comparator} cb
@return {bool} whether slices are equal


Comparator(left, right Object) int
	@param {Object} left
	@param {Object} right
	@return {int} -1 for less, 0 for equals, 1 for larger
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

@param {float64} min
@param {float64} max
@return {int} random number

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

@param {Seq} seq
@return {bool} whether this slice is empty or nil

*/
func IsEmpty(seq Seq) bool {
	return seq == nil || len(seq) == 0
}

/**
returns true if given object has type of slice

@param {Object} target
@return {bool} whether given object has type of slice

*/
func IsSlice(target Object) bool {
	return reflect.ValueOf(target).Kind() == reflect.Slice
}

/**
returns Seq from given object, filling the resulting Seq via reflection
NOTE: You should possibly avoid using this one

@param {Object} target
@param {int} size
@return {Seq} resulting filled(or not) Seq

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

	if IsEmpty(seq) || cb == nil {
		return -1
	}
	if length == 1 {
		return seq[0]
	}

	lastComputed = seq[0]

	for _, val := range seq {
		if sgn(cb(val, lastComputed)) == dir {
			lastComputed = val
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
		return -1
	} else if num == 0 {
		return 0
	} else {
		return 1
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
