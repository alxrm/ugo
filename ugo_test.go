package ugo_test

import (
	. "github.com/alxrm/ugo"
	. "github.com/franela/goblin"
	"math"
	"testing"
)

func TestSingleValues(t *testing.T) {
	g := Goblin(t)

	inSeq := Seq{2, 2, 4, 6, 7, 8, 10, 10, 17, 120}
	inSeqDifOrder := Seq{17, 2, 8, 6, 2, 4, 10, 7, 10, 120}
	intComparator := func(l, r Object) int { return l.(int) - r.(int) }
	reduceCollector := func(memo, cur, _, _ Object) Object { return memo.(int) + cur.(int) }
	searchPredicate := func(cur, _, _ Object) bool { return cur.(int) > 7 }

	g.Describe("#Min()", func() {
		g.It("Should return min value ", func() {
			g.Assert(Min(inSeq, intComparator)).Equal(2)
			g.Assert(Min(nil, intComparator)).Equal(-1)
			g.Assert(Min(inSeq, nil)).Equal(-1)
			g.Assert(Min(nil, nil)).Equal(-1)
		})
	})

	g.Describe("#Max()", func() {
		g.It("Should return max value ", func() {
			g.Assert(Max(inSeq, intComparator)).Equal(120)
			g.Assert(Max(nil, intComparator)).Equal(-1)
			g.Assert(Max(inSeq, nil)).Equal(-1)
			g.Assert(Max(nil, nil)).Equal(-1)
		})
	})

	g.Describe("#Reduce()", func() {
		g.It("Should return sum of slice elements", func() {
			g.Assert(Reduce(inSeq, reduceCollector, 0)).Equal(186)
			g.Assert(Reduce(inSeq, reduceCollector, nil)).Equal(186)
			g.Assert(Reduce(inSeq, nil, nil)).Equal(nil)
			g.Assert(Reduce(nil, nil, nil)).Equal(nil)
		})
	})

	g.Describe("#ReduceRigth()", func() {
		g.It("Should return sum of slice elements, starts from right", func() {
			g.Assert(ReduceRight(inSeq, reduceCollector, 0)).Equal(186)
			g.Assert(ReduceRight(inSeq, reduceCollector, nil)).Equal(186)
			g.Assert(ReduceRight(inSeq, nil, nil)).Equal(nil)
			g.Assert(ReduceRight(nil, nil, nil)).Equal(nil)
		})
	})

	g.Describe("#Find()", func() {
		g.It("Should return first value, which passes predicate test", func() {
			g.Assert(Find(inSeq, searchPredicate)).Equal(8)
			g.Assert(Find(nil, searchPredicate)).Equal(nil)
			g.Assert(Find(inSeq, nil)).Equal(nil)
			g.Assert(Find(nil, nil)).Equal(nil)
		})
	})

	g.Describe("#FindLast()", func() {
		g.It("Should return last value, which passes predicate test", func() {
			g.Assert(FindLast(inSeq, searchPredicate)).Equal(120)
			g.Assert(FindLast(nil, searchPredicate)).Equal(nil)
			g.Assert(FindLast(inSeq, nil)).Equal(nil)
			g.Assert(FindLast(nil, nil)).Equal(nil)
		})
	})

	g.Describe("#FindIndex()", func() {
		g.It("Should return first value's index, which passes predicate test", func() {
			g.Assert(FindIndex(inSeq, searchPredicate)).Equal(5)
			g.Assert(FindIndex(nil, searchPredicate)).Equal(-1)
			g.Assert(FindIndex(inSeq, nil)).Equal(-1)
			g.Assert(FindIndex(nil, nil)).Equal(-1)
		})
	})

	g.Describe("#FindLastIndex()", func() {
		g.It("Should return last value's index, which passes predicate test", func() {
			g.Assert(FindLastIndex(inSeq, searchPredicate)).Equal(9)
			g.Assert(FindLastIndex(nil, searchPredicate)).Equal(-1)
			g.Assert(FindLastIndex(inSeq, nil)).Equal(-1)
			g.Assert(FindLastIndex(nil, nil)).Equal(-1)
		})
	})

	g.Describe("#Some()", func() {
		g.It("Should return true if some of the elements have passed the predicate test", func() {
			g.Assert(Some(inSeq, searchPredicate)).IsTrue()
			g.Assert(Some(nil, searchPredicate)).IsFalse()
			g.Assert(Some(inSeq, nil)).IsFalse()
			g.Assert(Some(nil, nil)).IsFalse()
		})
	})

	g.Describe("#Every()", func() {
		g.It("Should return true if all of the elements have passed the predicate test", func() {
			g.Assert(Every(inSeq, searchPredicate)).IsFalse()
			g.Assert(Every(nil, searchPredicate)).IsFalse()
			g.Assert(Every(inSeq, nil)).IsFalse()
			g.Assert(Every(nil, nil)).IsFalse()
		})
	})

	g.Describe("#Contains()", func() {
		g.It("Should return true if slice contains target value", func() {
			g.Assert(Contains(inSeq, 7, true, intComparator)).IsTrue()
			g.Assert(Contains(inSeq, 88, true, intComparator)).IsFalse()
			g.Assert(Contains(inSeqDifOrder, 7, false, intComparator)).IsTrue()
			g.Assert(Contains(inSeqDifOrder, 88, false, intComparator)).IsFalse()
			g.Assert(Contains(inSeqDifOrder, 88, false, nil)).IsFalse()
			g.Assert(Contains(nil, 0, false, intComparator)).IsFalse()
			g.Assert(Contains(nil, 0, false, nil)).IsFalse()
		})
	})

	g.Describe("#IndexOf()", func() {
		g.It("Should return first index of target value", func() {
			g.Assert(IndexOf(inSeq, 7, true, intComparator)).Equal(4)
			g.Assert(IndexOf(inSeq, 88, true, intComparator)).Equal(-1)
			g.Assert(IndexOf(inSeqDifOrder, 7, false, intComparator)).Equal(7)
			g.Assert(IndexOf(inSeqDifOrder, 88, false, intComparator)).Equal(-1)
			g.Assert(IndexOf(inSeqDifOrder, 88, false, nil)).Equal(-1)
			g.Assert(IndexOf(nil, 0, false, intComparator)).Equal(-1)
			g.Assert(IndexOf(nil, 0, false, nil)).Equal(-1)
		})
	})

	g.Describe("#LastIndexOf()", func() {
		g.It("Should return last index of target value", func() {
			g.Assert(LastIndexOf(inSeq, 10, intComparator)).Equal(7)
			g.Assert(LastIndexOf(inSeq, 88, intComparator)).Equal(-1)
			g.Assert(LastIndexOf(inSeqDifOrder, 10, intComparator)).Equal(8)
			g.Assert(LastIndexOf(inSeqDifOrder, 88, intComparator)).Equal(-1)
			g.Assert(LastIndexOf(inSeqDifOrder, 88, nil)).Equal(-1)
			g.Assert(LastIndexOf(nil, 0, intComparator)).Equal(-1)
			g.Assert(LastIndexOf(nil, 0, nil)).Equal(-1)
		})
	})

	g.Describe("#EqualsStrict()", func() {
		g.It("Should return true if sliceA == sliceB, e. g. order and elements", func() {
			g.Assert(EqualsStrict(inSeq, inSeq, intComparator)).IsTrue()
			g.Assert(EqualsStrict(inSeq, inSeqDifOrder, intComparator)).IsFalse()
			g.Assert(EqualsStrict(nil, inSeq, intComparator)).IsFalse()
			g.Assert(EqualsStrict(inSeq, nil, intComparator)).IsFalse()
			g.Assert(EqualsStrict(inSeq, inSeqDifOrder, nil)).IsFalse()
			g.Assert(EqualsStrict(nil, nil, intComparator)).IsTrue()
			g.Assert(EqualsStrict(nil, nil, nil)).IsFalse()
		})
	})

	g.Describe("#EqualsNotStrict()", func() {
		g.It("Should return true if sliceA == sliceB, only elements", func() {
			g.Assert(EqualsNotStrict(inSeq, inSeq, intComparator)).IsTrue()
			g.Assert(EqualsNotStrict(inSeq, inSeqDifOrder, intComparator)).IsTrue()
			g.Assert(EqualsNotStrict(nil, inSeq, intComparator)).IsFalse()
			g.Assert(EqualsNotStrict(inSeq, nil, intComparator)).IsFalse()
			g.Assert(EqualsNotStrict(inSeq, inSeqDifOrder, nil)).IsFalse()
			g.Assert(EqualsNotStrict(nil, nil, intComparator)).IsTrue()
			g.Assert(EqualsNotStrict(nil, nil, nil)).IsFalse()
		})
	})
}

func TestMultipleValues(t *testing.T) {
	g := Goblin(t)

	inSeq := Seq{2, 4, 6, 7, 8, 10, 120, 10, 2, 17}
	intComparator := func(l, r Object) int { return l.(int) - r.(int) }
	changingCallback := func(cur, _, _ Object) Object { return cur.(int) - 2 }
	evenPredicate := func(cur, _, _ Object) bool { return cur.(int)%2 == 0 }
	evenCallback := func(cur, _, _ Object) Object {
		if cur.(int)%2 == 0 {
			return "even"
		} else {
			return "odd"
		}
	}

	g.Describe("#Each()", func() {
		g.It("Should call Action on each element of Seq", func() {
			inSeq := Seq{2, 4, 6, 7, 8, 10, 120, 10, 2, 17}
			outSeqResult := NewSeq(len(inSeq))
			outSeqTest := Seq{4, 16, 36, 49, 64, 100, 14400, 100, 4, 289}
			powAction := func(cur, index, _ Object) { outSeqResult[index.(int)] = cur.(int) * cur.(int) }

			Each(inSeq, powAction)
			g.Assert(outSeqResult).Equal(outSeqTest)

			Each(inSeq, nil)
			Each(nil, powAction)
			Each(nil, nil)

		})
	})

	g.Describe("#Map()", func() {
		g.It("Should return changed elements", func() {
			outSeqChanged := Seq{0, 2, 4, 5, 6, 8, 118, 8, 0, 15}

			g.Assert(Map(inSeq, changingCallback)).Equal(outSeqChanged)
			g.Assert(Map(nil, changingCallback)).Equal(Seq(nil))
			g.Assert(Map(inSeq, nil)).Equal(inSeq)
			g.Assert(Map(nil, nil)).Equal(Seq(nil))
		})
	})

	g.Describe("#Filter()", func() {
		g.It("Should return filtered elements", func() {
			outSeqFiltered := Seq{2, 4, 6, 8, 10, 120, 10, 2}

			g.Assert(Filter(inSeq, evenPredicate)).Equal(outSeqFiltered)
			g.Assert(Filter(inSeq, nil)).Equal(inSeq)
			g.Assert(Filter(nil, evenPredicate)).Equal(Seq(nil))
			g.Assert(Filter(nil, nil)).Equal(Seq(nil))
		})
	})

	g.Describe("#Reject()", func() {
		g.It("Should return values which NOT passed predicate test elements", func() {
			outSeqFiltered := Seq{7, 17}

			g.Assert(Reject(inSeq, evenPredicate)).Equal(outSeqFiltered)
			g.Assert(Reject(inSeq, nil)).Equal(inSeq)
			g.Assert(Reject(nil, evenPredicate)).Equal(Seq(nil))
			g.Assert(Reject(nil, nil)).Equal(Seq(nil))
		})
	})

	g.Describe("#SortBy()", func() {
		g.It("Should return sorted Seq", func() {
			outSeqSorted := Seq{2, 2, 4, 6, 7, 8, 10, 10, 17, 120}

			g.Assert(SortBy(inSeq, intComparator)).Equal(outSeqSorted)
			g.Assert(SortBy(inSeq, nil)).Equal(inSeq)
			g.Assert(SortBy(nil, intComparator)).Equal(Seq(nil))
			g.Assert(SortBy(nil, nil)).Equal(Seq(nil))
		})
	})

	g.Describe("#CountBy()", func() {
		g.It("Should return map with countings, e. g. keys - callback result, value - number of same results", func() {
			outMapCounted := map[string]int{"even": 8, "odd": 2}

			g.Assert(CountBy(inSeq, evenCallback)).Equal(outMapCounted)
			g.Assert(CountBy(inSeq, nil)).Equal(map[string]int(nil))
			g.Assert(CountBy(nil, evenCallback)).Equal(map[string]int(nil))
			g.Assert(CountBy(nil, nil)).Equal(map[string]int(nil))
		})
	})

	g.Describe("#Remove()", func() {
		g.It("Should return Seq without value in given index", func() {
			outSeqRemoved := Seq{2, 2, 6, 7, 8, 10, 10, 17, 120}

			g.Assert(Remove(inSeq, 2)).Equal(outSeqRemoved)
			g.Assert(Remove(inSeq, -1)).Equal(inSeq)
			g.Assert(Remove(inSeq, 30)).Equal(inSeq)
			g.Assert(Remove(nil, -1)).Equal(Seq(nil))
		})
	})

	g.Describe("#Shuffle()", func() {
		g.It("Should return shuffled Seq", func() {
			g.Assert(EqualsStrict(Shuffle(inSeq), inSeq, intComparator)).IsFalse()
			g.Assert(EqualsNotStrict(Shuffle(inSeq), inSeq, intComparator)).IsTrue()
			g.Assert(Shuffle(Seq{})).Equal(Seq{})
			g.Assert(Shuffle(nil)).Equal(Seq(nil))
		})
	})

	g.Describe("#Uniq()", func() {
		g.It("Should return Seq with no duplicates", func() {
			inSeq := Seq{2, 4, 6, 7, 8, 10, 120, 10, 2, 17}
			outSeqUniq := Seq{2, 4, 6, 7, 8, 10, 120, 17}

			g.Assert(Uniq(inSeq, intComparator)).Equal(outSeqUniq)
			g.Assert(Uniq(inSeq, nil)).Equal(inSeq)
			g.Assert(Uniq(nil, intComparator)).Equal(Seq(nil))
			g.Assert(Uniq(nil, nil)).Equal(Seq(nil))
		})
	})
}

func TestUtils(t *testing.T) {
	g := Goblin(t)

	g.Describe("#Random()", func() {
		g.It("Should return random number from min to max", func() {
			g.Assert(Random(-100, 0) <= 0).IsTrue()
			g.Assert(0 <= Random(0, 100)).IsTrue()
			g.Assert(Random(0, 0)).Equal(0)
			g.Assert(Random(math.Inf(-1), math.Inf(-1))).Equal(0)
			g.Assert(Random(math.Inf(-1), math.Inf(1)) != int(math.Inf(-1))).IsTrue()
		})
	})

	g.Describe("#IsSlice()", func() {
		g.It("Should check whether it is a Slice from given Object", func() {
			g.Assert(IsSlice(Seq{0, 1, 2})).IsTrue()
			g.Assert(IsSlice([]int{0, 1, 2})).IsTrue()
			g.Assert(IsSlice("not slice")).IsFalse()
			g.Assert(IsSlice(Seq(nil))).IsTrue()
			g.Assert(IsSlice(nil)).IsFalse()
		})
	})

	g.Describe("#IsEmpty()", func() {
		g.It("Should check whether it is an empty sequence", func() {
			g.Assert(IsEmpty(nil)).IsTrue()
			g.Assert(IsEmpty(Seq{})).IsTrue()
			g.Assert(IsEmpty(Seq(nil))).IsTrue()
			g.Assert(IsEmpty(Seq{0, 1, 2})).IsFalse()
		})
	})

	g.Describe("#From()", func() {
		g.It("Should return seq from int slice", func() {
			g.Assert(From([]string{"fst", "snd"}, 2)).Equal(Seq{"fst", "snd"})
			g.Assert(From([]int{}, -1)).Equal(Seq{})
			g.Assert(From([]string{"fst", "snd"}, -1)).Equal(Seq{})
			g.Assert(From("non slice", 2)).Equal(Seq{})
			g.Assert(From(nil, 0)).Equal(Seq{})
		})
	})
}
