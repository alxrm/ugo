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

package ugo_test

import (
	. "github.com/alxrm/ugo"
	. "github.com/franela/goblin"
	"math"
	"testing"
)

func TestChainingSingleValues(t *testing.T) {
	g := Goblin(t)

	inEmpty := Seq{}
	inSingle := Seq{0}

	clctr := func(memo, cur, _, _ Object) Object {
		return memo.(int) + cur.(int)
	}

	pred := func(cur, _, _ Object) bool {
		return cur.(int) > 7
	}

	comp := func(l, r Object) int {
		return l.(int) - r.(int)
	}

	g.Describe("#Min()", func() {
		g.It("Should return min value by chaining ", func() {
			inSeq := Seq{2, 99, -12, 884, 8}
			inSeqDif := Seq{1, 0, -math.MaxInt32, -math.MinInt16}

			g.Assert(Chain(inSeq).Min(comp).Value()).Equal(-12)
			g.Assert(Chain(inSeqDif).Min(comp).Value()).Equal(math.MinInt32 + 1)
			g.Assert(Chain(inSingle).Min(comp).Value()).Equal(0)
			g.Assert(Chain(inEmpty).Min(comp).Value()).Equal(-1)
			g.Assert(Chain(nil).Min(comp).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).Min(nil).Value()).Equal(-1)
			g.Assert(Chain(nil).Min(nil).Value()).Equal(-1)
		})
	})

	g.Describe("#Max()", func() {
		g.It("Should return max value by chaining ", func() {
			inSeq := Seq{433, 39, 92, -12, -math.MinInt32}
			inSeqDif := Seq{1, 0, -1, 0}

			g.Assert(Chain(inSeq).Max(comp).Value()).Equal(-math.MinInt32)
			g.Assert(Chain(inSeqDif).Max(comp).Value()).Equal(1)
			g.Assert(Chain(inSingle).Max(comp).Value()).Equal(0)
			g.Assert(Chain(inEmpty).Max(comp).Value()).Equal(-1)
			g.Assert(Chain(nil).Max(comp).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).Max(nil).Value()).Equal(-1)
			g.Assert(Chain(nil).Max(nil).Value()).Equal(-1)
		})
	})

	g.Describe("#Find()", func() {
		g.It("Should return first value, which passes predicate test", func() {
			inSeq := Seq{39, 92, 2, math.MinInt32}
			inSeqDif := Seq{7, 9, 0, -12000}

			g.Assert(Chain(inSeq).Find(pred).Value()).Equal(39)
			g.Assert(Chain(inSeqDif).Find(pred).Value()).Equal(9)
			g.Assert(Chain(inSingle).Find(pred).Value()).Equal(nil)
			g.Assert(Chain(inEmpty).Find(pred).Value()).Equal(nil)
			g.Assert(Chain(nil).Find(pred).Value()).Equal(nil)
			g.Assert(Chain(inSeqDif).Find(nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Find(nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FindLast()", func() {
		g.It("Should return last value, which passes predicate test", func() {
			inSeq := Seq{39, 92, 2, math.MinInt32}
			inSeqDif := Seq{7, 9, 0, 8, -12000}

			g.Assert(Chain(inSeq).FindLast(pred).Value()).Equal(92)
			g.Assert(Chain(inSeqDif).FindLast(pred).Value()).Equal(8)
			g.Assert(Chain(inSingle).FindLast(pred).Value()).Equal(nil)
			g.Assert(Chain(inEmpty).FindLast(pred).Value()).Equal(nil)
			g.Assert(Chain(nil).FindLast(pred).Value()).Equal(nil)
			g.Assert(Chain(inSeqDif).FindLast(nil).Value()).Equal(nil)
			g.Assert(Chain(nil).FindLast(nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FindIndex()", func() {
		g.It("Should return first value's index, which passes predicate test", func() {
			inSeq := Seq{39, 92, 2, math.MinInt32}
			inSeqDif := Seq{7, 9, 0, 8, -12000}

			g.Assert(Chain(inSeq).FindIndex(pred).Value()).Equal(0)
			g.Assert(Chain(inSeqDif).FindIndex(pred).Value()).Equal(1)
			g.Assert(Chain(inSingle).FindIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(inEmpty).FindIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(nil).FindIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).FindIndex(nil).Value()).Equal(-1)
			g.Assert(Chain(nil).FindIndex(nil).Value()).Equal(-1)
		})
	})

	g.Describe("#FindLastIndex()", func() {
		g.It("Should return last value's index, which passes predicate test", func() {
			inSeq := Seq{39, 92, 2, math.MinInt32}
			inSeqDif := Seq{7, 9, 0, 8, -12000}

			g.Assert(Chain(inSeq).FindLastIndex(pred).Value()).Equal(1)
			g.Assert(Chain(inSeqDif).FindLastIndex(pred).Value()).Equal(3)
			g.Assert(Chain(inSingle).FindLastIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(inEmpty).FindLastIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(nil).FindLastIndex(pred).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).FindLastIndex(nil).Value()).Equal(-1)
			g.Assert(Chain(nil).FindLastIndex(nil).Value()).Equal(-1)
		})
	})

	g.Describe("#Detect()", func() {
		g.It("Should clone #Find behaviour", func() {
			inSeq := Seq{-12, 2, 6, 33, -12, 8, 92}
			inSeqDif := Seq{3, 2, 1, 0, 99}

			g.Assert(Chain(inSeq).Detect(pred).Value()).Equal(33)
			g.Assert(Chain(inSeqDif).Detect(pred).Value()).Equal(99)
			g.Assert(Chain(inSingle).Detect(pred).Value()).Equal(nil)
			g.Assert(Chain(inEmpty).Detect(pred).Value()).Equal(nil)
			g.Assert(Chain(nil).Detect(pred).Value()).Equal(nil)
			g.Assert(Chain(inSeqDif).Detect(nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Detect(nil).Value()).Equal(nil)
		})
	})

	g.Describe("#Reduce()", func() {
		g.It("Should return sum of slice elements", func() {
			inSeq := Seq{39, 92, 2, math.MaxInt32}
			inSeqDif := Seq{7, 9, 0, -12000}

			g.Assert(Chain(inSeq).Reduce(clctr, 0).Value()).Equal(2147483780)
			g.Assert(Chain(inSeqDif).Reduce(clctr, nil).Value()).Equal(-11984)
			g.Assert(Chain(nil).Reduce(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).Reduce(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Reduce(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#Inject()", func() {
		g.It("Should clone #Reduce behaviour", func() {
			inSeq := Seq{39, 34, 3, 230}
			inSeqDif := Seq{7, 230, 0, -1}

			g.Assert(Chain(inSeq).Inject(clctr, 0).Value()).Equal(306)
			g.Assert(Chain(inSeqDif).Inject(clctr, nil).Value()).Equal(236)
			g.Assert(Chain(nil).Inject(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).Inject(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Inject(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FoldL()", func() {
		g.It("Should clone #Reduce behaviour", func() {
			inSeq := Seq{44, 2, -10, 0}
			inSeqDif := Seq{-7, 0, 53, -20}

			g.Assert(Chain(inSeq).FoldL(clctr, 0).Value()).Equal(36)
			g.Assert(Chain(inSeqDif).FoldL(clctr, nil).Value()).Equal(26)
			g.Assert(Chain(nil).FoldL(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).FoldL(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).FoldL(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#ReduceRigth()", func() {
		g.It("Should return sum of slice elements, starts from right", func() {
			inSeq := Seq{39, 2, math.MaxInt32, -12}
			inSeqDif := Seq{8237, 7, 9, 0, -1493, math.MinInt16}

			g.Assert(Chain(inSeq).ReduceRight(clctr, 0).Value()).Equal(2147483676)
			g.Assert(Chain(inSeqDif).ReduceRight(clctr, nil).Value()).Equal(-26008)
			g.Assert(Chain(nil).ReduceRight(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).ReduceRight(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).ReduceRight(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FoldR()", func() {
		g.It("Should clone #ReduceRight behaviour", func() {
			inSeq := Seq{39, 2, 32}
			inSeqDif := Seq{7}

			g.Assert(Chain(inSeq).FoldR(clctr, 0).Value()).Equal(73)
			g.Assert(Chain(inSeqDif).FoldR(clctr, nil).Value()).Equal(7)
			g.Assert(Chain(nil).FoldR(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).FoldR(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).FoldR(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#Some()", func() {
		g.It("Should return true if at least one of the elements have passed the predicate test", func() {
			inSeq := Seq{77, 34, 242, 99, -20, 0, 3, 0}
			inSeqDif := Seq{3, 1, 1, -12, 0, math.MinInt16, 2}

			g.Assert(Chain(inSeq).Some(pred).Value()).IsTrue()
			g.Assert(Chain(inSeqDif).Some(pred).Value()).IsFalse()
			g.Assert(Chain(inSingle).Some(pred).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Some(pred).Value()).IsFalse()
			g.Assert(Chain(nil).Some(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Some(nil).Value()).IsFalse()
			g.Assert(Chain(nil).Some(nil).Value()).IsFalse()
		})
	})

	g.Describe("#Any()", func() {
		g.It("Should clone #Some behaviour", func() {
			inSeq := Seq{7, -8, 0, 0}
			inSeqDif := Seq{3, 7, 0, math.MinInt16, 8}

			g.Assert(Chain(inSeq).Any(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Any(pred).Value()).IsTrue()
			g.Assert(Chain(inSingle).Any(pred).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Any(pred).Value()).IsFalse()
			g.Assert(Chain(nil).Any(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Any(nil).Value()).IsFalse()
			g.Assert(Chain(nil).Any(nil).Value()).IsFalse()
		})
	})

	g.Describe("#Every()", func() {
		g.It("Should return true if all of the elements have passed the predicate test", func() {
			inSeq := Seq{10, 9, 30, 0}
			inSeqDif := Seq{219, 92, 10, 8}

			g.Assert(Chain(inSeq).Every(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Every(pred).Value()).IsTrue()
			g.Assert(Chain(inSingle).Every(pred).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Every(pred).Value()).IsFalse()
			g.Assert(Chain(nil).Every(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Every(nil).Value()).IsFalse()
			g.Assert(Chain(nil).Every(nil).Value()).IsFalse()
		})
	})

	g.Describe("#All()", func() {
		g.It("Should clone #Every behaviour", func() {
			inSeq := Seq{10, 9, 30, 299}
			inSeqDif := Seq{219, 92, -10, 8}

			g.Assert(Chain(inSeq).All(pred).Value()).IsTrue()
			g.Assert(Chain(inSeqDif).All(pred).Value()).IsFalse()
			g.Assert(Chain(inSingle).All(pred).Value()).IsFalse()
			g.Assert(Chain(inEmpty).All(pred).Value()).IsFalse()
			g.Assert(Chain(nil).All(pred).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).All(nil).Value()).IsFalse()
			g.Assert(Chain(nil).All(nil).Value()).IsFalse()
		})
	})

	g.Describe("#Contains()", func() {
		g.It("Should return true if slice contains target value", func() {
			inSeq := Seq{0, 9, 21, 299, 300, math.MaxInt32, math.MaxInt32 + 100}
			inSeqDif := Seq{29, 92, 123, 8, 123}

			g.Assert(Chain(inSeq).Contains(300, true, comp).Value()).IsTrue()
			g.Assert(Chain(inSeqDif).Contains(8, false, comp).Value()).IsTrue()
			g.Assert(Chain(inSingle).Contains(-12, false, comp).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Contains(nil, false, comp).Value()).IsFalse()
			g.Assert(Chain(nil).Contains(300, true, nil).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Contains(nil, true, nil).Value()).IsFalse()
			g.Assert(Chain(nil).Contains(nil, true, nil).Value()).IsFalse()
		})
	})

	g.Describe("#Includes()", func() {
		g.It("Should clone #Contains behaviour", func() {
			inSeq := Seq{211, 325, 420, 999, math.MaxInt32, math.MaxInt32 + 100}
			inSeqDif := Seq{1, -8, 910, 0, 19, 2, 3, 8}

			g.Assert(Chain(inSeq).Includes(777, true, comp).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Includes(3, false, comp).Value()).IsTrue()
			g.Assert(Chain(inSingle).Includes(-1, false, comp).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Includes(nil, false, comp).Value()).IsFalse()
			g.Assert(Chain(nil).Includes(300, true, nil).Value()).IsFalse()
			g.Assert(Chain(inSeqDif).Includes(nil, true, nil).Value()).IsFalse()
			g.Assert(Chain(nil).Includes(nil, true, nil).Value()).IsFalse()
		})
	})

	g.Describe("#IndexOf()", func() {
		g.It("Should return first index of target value", func() {
			inSeq := Seq{-100, 0, 100, 32521}
			inSeqDif := Seq{1, 8}

			g.Assert(Chain(inSeq).IndexOf(0, true, comp).Value()).Equal(1)
			g.Assert(Chain(inSeqDif).IndexOf(3, false, comp).Value()).Equal(-1)
			g.Assert(Chain(inSingle).IndexOf(-1, false, comp).Value()).Equal(-1)
			g.Assert(Chain(inEmpty).IndexOf(nil, false, comp).Value()).Equal(-1)
			g.Assert(Chain(nil).IndexOf(300, true, nil).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).IndexOf(nil, true, nil).Value()).Equal(-1)
			g.Assert(Chain(nil).IndexOf(nil, true, nil).Value()).Equal(-1)
		})
	})

	g.Describe("#LastIndexOf()", func() {
		g.It("Should return last index of target value", func() {
			inSeq := Seq{-100, 0, 8, 8, 100, 32521}
			inSeqDif := Seq{1, 8, 2, -12, 2, 0}

			g.Assert(Chain(inSeq).LastIndexOf(8, comp).Value()).Equal(3)
			g.Assert(Chain(inSeqDif).LastIndexOf(2, comp).Value()).Equal(4)
			g.Assert(Chain(inSingle).LastIndexOf(-1, comp).Value()).Equal(-1)
			g.Assert(Chain(inEmpty).LastIndexOf(nil, comp).Value()).Equal(-1)
			g.Assert(Chain(nil).LastIndexOf(300, nil).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).LastIndexOf(nil, nil).Value()).Equal(-1)
			g.Assert(Chain(nil).LastIndexOf(nil, nil).Value()).Equal(-1)
		})
	})

	g.Describe("#EqualsStrict()", func() {
		g.It("Should return true if sliceA == sliceB, e. g. order and elements", func() {
			leftSeq := Seq{23, 3, 9, 9, 1, -12, math.MinInt32, 9}
			rightSeq := Seq{23, math.MinInt32, 9, 3, 9, 9, 1, -12}
			rightSeqIdent := Seq{23, 3, 9, 9, 1, -12, math.MinInt32, 9}

			g.Assert(Chain(leftSeq).EqualsStrict(rightSeq, comp).Value()).IsFalse()
			g.Assert(Chain(leftSeq).EqualsStrict(rightSeqIdent, comp).Value()).IsTrue()
			g.Assert(Chain(nil).EqualsStrict(rightSeqIdent, comp).Value()).IsFalse()
			g.Assert(Chain(nil).EqualsStrict(nil, comp).Value()).IsTrue()
			g.Assert(Chain(leftSeq).EqualsStrict(rightSeq, nil).Value()).IsFalse()
			g.Assert(Chain(leftSeq).EqualsStrict(nil, nil).Value()).IsFalse()
			g.Assert(Chain(nil).EqualsStrict(nil, nil).Value()).IsFalse()
		})
	})

	g.Describe("#EqualsNotStrict()", func() {
		g.It("Should return true if sliceA == sliceB, only elements", func() {
			leftSeq := Seq{23, 3, 9, 9, 1, -12, math.MinInt32, 9}
			rightSeq := Seq{23, math.MinInt32, 9, 3, 9, 9, 1, -12}
			rightSeqIdent := Seq{23, 3, 9, 9, 1, -12, math.MinInt32, 9}

			g.Assert(Chain(leftSeq).EqualsNotStrict(rightSeq, comp).Value()).IsTrue()
			g.Assert(Chain(leftSeq).EqualsNotStrict(rightSeqIdent, comp).Value()).IsTrue()
			g.Assert(Chain(nil).EqualsNotStrict(rightSeqIdent, comp).Value()).IsFalse()
			g.Assert(Chain(nil).EqualsNotStrict(nil, comp).Value()).IsTrue()
			g.Assert(Chain(leftSeq).EqualsNotStrict(rightSeq, nil).Value()).IsFalse()
			g.Assert(Chain(leftSeq).EqualsNotStrict(nil, nil).Value()).IsFalse()
			g.Assert(Chain(nil).EqualsNotStrict(nil, nil).Value()).IsFalse()
		})
	})
}

func TestChainingMultipleValues(t *testing.T) {
	g := Goblin(t)

	inEmpty := Seq{}
	inSingle := Seq{0}

	pred := func(cur, _, _ Object) bool {
		return cur.(int) < 7
	}

	comp := func(l, r Object) int {
		return l.(int) - r.(int)
	}

	clb := func(cur, _, _ Object) Object {
		if cur.(int)%2 == 0 {
			return "even"
		}

		return "odd"
	}

	g.Describe("#Each()", func() {
		g.It("Should call Action on each element of Seq", func() {

			inSeq := Seq{3, 4, -9, 0, 11, 1}
			inSeqDif := Seq{0, 0, -1, -2, -7, 6}

			outSeq := Seq{1, 121, 0, 81, 16, 9}
			outSeqDif := Seq{36, 49, 4, 1, 0, 0}

			tmpSeq := NewSeq(6)
			tmpSeqDif := NewSeq(6)

			act := func(cur, index, _ Object) {
				tmpSeq[(len(tmpSeq)-1)-index.(int)] = cur.(int) * cur.(int)
			}

			actDif := func(cur, index, _ Object) {
				tmpSeqDif[(len(tmpSeqDif)-1)-index.(int)] = cur.(int) * cur.(int)
			}

			Chain(inSeq).Each(act)
			g.Assert(tmpSeq).Equal(outSeq)

			Chain(inSeqDif).Each(actDif)
			g.Assert(tmpSeqDif).Equal(outSeqDif)

			Chain(nil).Each(act)
			g.Assert(tmpSeq).Equal(tmpSeq)

			Chain(inSeq).Each(nil)
			g.Assert(tmpSeq).Equal(tmpSeq)

			Chain(nil).Each(nil)
			g.Assert(tmpSeq).Equal(tmpSeq)
		})
	})

	g.Describe("#ForEach()()", func() {
		g.It("Should clone #Each behaviour", func() {

			inSeq := Seq{2, 4, -9, 0, 0, 12}
			inSeqDif := Seq{0, -1}

			outSeq := Seq{4, 16, 81, 0, 0, 144}
			outSeqDif := Seq{0, 1}

			tmpSeq := NewSeq(6)
			tmpSeqDif := NewSeq(2)

			act := func(cur, index, _ Object) {
				tmpSeq[index.(int)] = cur.(int) * cur.(int)
			}

			actDif := func(cur, index, _ Object) {
				tmpSeqDif[index.(int)] = cur.(int) * cur.(int)
			}

			Chain(inSeq).ForEach(act)
			g.Assert(tmpSeq).Equal(outSeq)

			Chain(inSeqDif).ForEach(actDif)
			g.Assert(tmpSeqDif).Equal(outSeqDif)

			Chain(nil).ForEach(act)
			g.Assert(tmpSeq).Equal(tmpSeq)

			Chain(inSeq).ForEach(nil)
			g.Assert(tmpSeq).Equal(tmpSeq)

			Chain(nil).ForEach(nil)
			g.Assert(tmpSeq).Equal(tmpSeq)
		})
	})

	g.Describe("#Uniq()", func() {
		g.It("Should return Seq with no duplicates", func() {
			inSeq := Seq{23, 3, 9, 9, 1, -12, math.MinInt32, 9}
			inSeqDif := Seq{23, 2, 1, 2, math.MinInt32, 9, 3, 9, 9, 1, -12}

			outSeq := Seq{23, 3, 9, 1, -12, math.MinInt32}
			outSeqDif := Seq{23, 2, 1, math.MinInt32, 9, 3, -12}

			g.Assert(Chain(inSeq).Uniq(comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Uniq(comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Uniq(comp).Value()).Equal(inSingle)
			g.Assert(Chain(nil).Uniq(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Uniq(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Uniq(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Uniq(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Unique()", func() {
		g.It("Should clone #Uniq behaviour", func() {
			inSeq := Seq{23, 3, 9, 9, 7, -12, math.MinInt32, 9}
			inSeqDif := Seq{23, 2, 1, 2, math.MinInt32, 9, 0, 0, -0, 9, 9, 1, -12}

			outSeq := Seq{23, 3, 9, 7, -12, math.MinInt32}
			outSeqDif := Seq{23, 2, 1, math.MinInt32, 9, 0, -12}

			g.Assert(Chain(inSeq).Unique(comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Unique(comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Unique(comp).Value()).Equal(inSingle)
			g.Assert(Chain(nil).Unique(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Unique(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Unique(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Unique(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#SortBy()", func() {
		g.It("Should return sorted Seq", func() {
			inSeq := Seq{4, 3, 43, 2, 3, -92, 102, 2, 0}
			inSeqDif := Seq{23, 2, 1, 2, math.MinInt32, 9, 0, 0, -0, 9, 9, 1, -12}

			outSeq := Seq{-92, 0, 2, 2, 3, 3, 4, 43, 102}
			outSeqDif := Seq{math.MinInt32, -12, 0, 0, 0, 1, 1, 2, 2, 9, 9, 9, 23}

			g.Assert(Chain(inSeq).SortBy(comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).SortBy(comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).SortBy(comp).Value()).Equal(inSingle)
			g.Assert(Chain(nil).SortBy(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).SortBy(comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).SortBy(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).SortBy(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#CountBy()", func() {
		g.It("Should return map with countings, e. g. keys - callback result, value - number of same results", func() {
			inSeq := Seq{4, 3, 43, 2, 3, -92, 102, 2, 0}
			inSeqDif := Seq{23, 2, 1, 2, math.MinInt32, 9, 0, 0, -0, 9, 9, 1, -12}

			out := map[string]int{"odd": 3, "even": 6}
			outSingle := map[string]int{"even": 1}
			outDif := map[string]int{"odd": 6, "even": 7}
			outEmpty := map[string]int{}

			g.Assert(Chain(inSeq).CountBy(clb).Value()).Equal(out)
			g.Assert(Chain(inSeqDif).CountBy(clb).Value()).Equal(outDif)
			g.Assert(Chain(inSingle).CountBy(clb).Value()).Equal(outSingle)
			g.Assert(Chain(nil).CountBy(clb).Value()).Equal(outEmpty)
			g.Assert(Chain(inEmpty).CountBy(clb).Value()).Equal(outEmpty)
			g.Assert(Chain(inEmpty).CountBy(nil).Value()).Equal(outEmpty)
			g.Assert(Chain(nil).CountBy(nil).Value()).Equal(outEmpty)
		})
	})

	g.Describe("#GroupBy()", func() {
		g.It("Should return map where keys are Callback result and values are elements, that produce such result", func() {
			inSeq := Seq{4, 3, 43, 2, 3, -92, 102, 2, 0}
			inSeqDif := Seq{23, 2, 1, 2, math.MinInt32, 9, 0, 0, -0, 9, 9, 1, -12}

			out := map[Object]Seq{"odd": {3, 43, 3}, "even": {4, 2, -92, 102, 2, 0}}
			outSingle := map[Object]Seq{"even": {0}}
			outDif := map[Object]Seq{
				"odd":  {23, 1, 9, 9, 9, 1},
				"even": {2, 2, math.MinInt32, 0, 0, 0, -12}}
			outEmpty := map[Object]Seq{}

			g.Assert(Chain(inSeq).GroupBy(clb).Value()).Equal(out)
			g.Assert(Chain(inSeqDif).GroupBy(clb).Value()).Equal(outDif)
			g.Assert(Chain(inSingle).GroupBy(clb).Value()).Equal(outSingle)
			g.Assert(Chain(nil).GroupBy(clb).Value()).Equal(outEmpty)
			g.Assert(Chain(inEmpty).GroupBy(clb).Value()).Equal(outEmpty)
			g.Assert(Chain(inEmpty).GroupBy(nil).Value()).Equal(outEmpty)
			g.Assert(Chain(nil).GroupBy(nil).Value()).Equal(outEmpty)
		})
	})

	g.Describe("#Map()", func() {
		g.It("Should return changed elements", func() {
			inSeq := Seq{-1, 0, 11, 1, 11}
			inSeqDif := Seq{0, 0, 0, 1, 2, 7, 67, -50000}
			//
			outSeq := Seq{"odd", "even", "odd", "odd", "odd"}
			outSeqDif := Seq{"even", "even", "even", "odd", "even", "odd", "odd", "even"}

			outSingle := Seq{"even"}

			g.Assert(Chain(inSeq).Map(clb).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Map(clb).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Map(clb).Value()).Equal(outSingle)
			g.Assert(Chain(nil).Map(clb).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Map(clb).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Map(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Map(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Shuffle()", func() {
		g.It("Should return shuffled Seq", func() {
			inSeq := Seq{-1, 0, 11, 1, 11}
			inSeqDif := Seq{0, 0, 0, 1, 2, 7, 67, -50000}

			g.Assert(Chain(inSeq).Shuffle().EqualsNotStrict(inSeq, comp).Value()).IsTrue()
			g.Assert(Chain(inSeqDif).Shuffle().EqualsNotStrict(inSeqDif, comp).Value()).IsTrue()
			g.Assert(Chain(inSingle).Shuffle().EqualsNotStrict(inSingle, comp).Value()).IsTrue()
			g.Assert(Chain(nil).Shuffle().EqualsNotStrict(inSeqDif, comp).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Shuffle().EqualsNotStrict(inSingle, comp).Value()).IsFalse()
			g.Assert(Chain(inEmpty).Shuffle().EqualsNotStrict(inEmpty, comp).Value()).IsTrue()
			g.Assert(Chain(nil).Shuffle().EqualsNotStrict(nil, comp).Value()).IsTrue()
		})
	})

	g.Describe("#Reverse()", func() {
		g.It("Should return reversed Seq", func() {
			inSeq := Seq{-1, 0, 11, 1, 11}
			inSeqDif := Seq{0, 0, 0, 1, 2, 7, 67, -50000}

			outSeq := Seq{11, 1, 11, 0, -1}
			outSeqDif := Seq{-50000, 67, 7, 2, 1, 0, 0, 0}

			g.Assert(Chain(inSeq).Reverse().Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Reverse().Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Reverse().Value()).Equal(inSingle)
			g.Assert(Chain(nil).Reverse().Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Reverse().Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Reverse().Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Reverse().Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Remove()", func() {
		g.It("Should return Seq without value in given index", func() {
			inSeq := Seq{3, 34, 23, 333, -12}
			inSeqDif := Seq{0, 0, 0, 1, 1, 92, -19, 42}

			outSeqEnd := Seq{3, 34, 23, 333}
			outSeqDif := Seq{0, 0, 0, 1, 92, -19, 42}
			outSeqDifStart := Seq{0, 0, 1, 1, 92, -19, 42}

			g.Assert(Chain(inSeq).Remove(123).Value()).Equal(outSeqEnd)
			g.Assert(Chain(inSeqDif).Remove(4).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSeqDif).Remove(-12).Value()).Equal(outSeqDifStart)
			g.Assert(Chain(inSingle).Remove(1).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Remove(112).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Remove(-15).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Remove(32).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Remove(0).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Insert()", func() {
		g.It("Should return Seq with new value inserted to given index", func() {
			inSeq := Seq{3, 34, 23, 333, -12}
			inSeqDif := Seq{0, 0, 0, 1, 1, 92, -19, 42}

			outSeqEnd := Seq{3, 34, 23, 333, -12, 999}
			outSeqDif := Seq{0, 0, 0, 1, 1, 9, 92, -19, 42}
			outSeqDifStart := Seq{-122, 0, 0, 0, 1, 1, 92, -19, 42}

			g.Assert(Chain(inSeq).Insert(999, 129).Value()).Equal(outSeqEnd)
			g.Assert(Chain(inSeqDif).Insert(9, 5).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSeqDif).Insert(-122, -122).Value()).Equal(outSeqDifStart)
			g.Assert(Chain(inEmpty).Insert(0, 192).Value()).Equal(inSingle)
			g.Assert(Chain(nil).Insert(112, 9).Value()).Equal(Seq{112})
			g.Assert(Chain(inEmpty).Insert(-15, 102).Value()).Equal(Seq{-15})
			g.Assert(Chain(inEmpty).Insert(32, -12).Value()).Equal(Seq{32})
			g.Assert(Chain(nil).Insert(0, 0).Value()).Equal(inSingle)
		})
	})

	g.Describe("#Concat()", func() {
		g.It("Should return slice, with appended another slice", func() {
			inSeq := Seq{232, -123, 0, math.MaxInt32}
			inSeqDif := Seq{-21, 0, 0, 0, 12}

			nextSeq := Seq{23, 20, 21}
			nextSeqDif := Seq{-1, -1, -2}

			outSeq := Seq{232, -123, 0, math.MaxInt32, 23, 20, 21}
			outSeqDif := Seq{-21, 0, 0, 0, 12, -1, -1, -2}
			outSeqSingle := Seq{0, -1, -1, -2}
			outSeqEmpty := Seq{23, 20, 21}

			g.Assert(Chain(inSeq).Concat(nextSeq).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Concat(nextSeqDif).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Concat(nextSeqDif).Value()).Equal(outSeqSingle)
			g.Assert(Chain(nil).Concat(nextSeq).Value()).Equal(outSeqEmpty)
			g.Assert(Chain(inEmpty).Concat(nextSeq).Value()).Equal(outSeqEmpty)
			g.Assert(Chain(nil).Concat(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Concat(nil).Value()).Equal(inSeq)
			g.Assert(Chain(nil).Concat(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Collect()", func() {
		g.It("Should clone #Map behaviour", func() {
			inSeq := Seq{0, -1, 0}
			inSeqDif := Seq{-199999, 0, 0, 0, 1000001}

			outSeq := Seq{"even", "odd", "even"}
			outSeqDif := Seq{"odd", "even", "even", "even", "odd"}

			outSingle := Seq{"even"}

			g.Assert(Chain(inSeq).Collect(clb).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Collect(clb).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Collect(clb).Value()).Equal(outSingle)
			g.Assert(Chain(nil).Collect(clb).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Collect(clb).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Collect(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Collect(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Filter()", func() {
		g.It("Should return filtered elements(only < 7)", func() {
			inSeq := Seq{0, -1, 0}
			inSeqDif := Seq{-199999, 8, 8, 9, 7, 0, 0, 0, 1000001}

			outSeq := Seq{0, -1, 0}
			outSeqDif := Seq{-199999, 0, 0, 0}

			outSingle := Seq{0}

			g.Assert(Chain(inSeq).Filter(pred).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Filter(pred).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Filter(pred).Value()).Equal(outSingle)
			g.Assert(Chain(nil).Filter(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Filter(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Filter(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Filter(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Select()", func() {
		g.It("Should clone #Filter behaviour", func() {
			inSeq := Seq{0, 6, 91, 8, 6}
			inSeqDif := Seq{273, 888, 923}

			outSeq := Seq{0, 6, 6}
			outSeqDif := Seq{}

			outSingle := Seq{0}

			g.Assert(Chain(inSeq).Select(pred).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Select(pred).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Select(pred).Value()).Equal(outSingle)
			g.Assert(Chain(nil).Select(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Select(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Select(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Select(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Reject()", func() {
		g.It("Should return values which NOT passed predicate test elements", func() {
			inSeq := Seq{0, 6, 91, 8, 6}
			inSeqDif := Seq{273, 888, 923}

			outSeq := Seq{91, 8}
			outSeqDif := Seq{273, 888, 923}

			outSingle := Seq{}

			g.Assert(Chain(inSeq).Reject(pred).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Reject(pred).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Reject(pred).Value()).Equal(outSingle)
			g.Assert(Chain(nil).Reject(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Reject(pred).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Reject(nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Reject(nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Difference()", func() {
		g.It("Should return Seq with elements, that are not in the other Seq", func() {
			inSeq := Seq{12, 33, 22, 22}
			inSeqDif := Seq{3, 4, 5}

			nextSeq := Seq{21, 22, 4, 12}
			nextSeqDif := Seq{3, 5, 5, 4, 5}

			outSeq := Seq{33}
			outSeqDif := Seq{}

			g.Assert(Chain(inSeq).Difference(nextSeq, comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Difference(nextSeqDif, comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Difference(nextSeqDif, comp).Value()).Equal(inSingle)
			g.Assert(Chain(nil).Difference(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Difference(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Difference(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Difference(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Difference(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Difference(nil, nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Difference(nil, nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Intersection()", func() {
		g.It("Should return elements, which are in both of the slices", func() {
			inSeq := Seq{12, 33, 22, 22}
			inSeqDif := Seq{3, 4, 5}

			nextSeq := Seq{21, 22, 4, 12}
			nextSeqDif := Seq{3, 5, 5, 4, 5}

			outSeq := Seq{12, 22}
			outSeqDif := Seq{3, 4, 5}

			g.Assert(Chain(inSeq).Intersection(nextSeq, comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Intersection(nextSeqDif, comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Intersection(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Intersection(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Intersection(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Intersection(nextSeqDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Intersection(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Intersection(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Intersection(nil, nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Intersection(nil, nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Union()", func() {
		g.It("Should return unique values, that at least once appeared in any of slices", func() {
			inSeq := Seq{12, 33, 22, 22}
			inSeqDif := Seq{3, 4, 5}

			nextSeq := Seq{21, 22, 4, 12}
			nextSeqDif := Seq{3, 5, 5, 4, 5}

			outSeq := Seq{12, 33, 22, 21, 4}
			outSeqDif := Seq{3, 4, 5}
			outSingleSeqDif := Seq{0, 3, 5, 4}
			outEmptySeqDif := Seq{3, 5, 4}

			g.Assert(Chain(inSeq).Union(nextSeq, comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Union(nextSeqDif, comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Union(nextSeqDif, comp).Value()).Equal(outSingleSeqDif)
			g.Assert(Chain(nil).Union(nextSeqDif, comp).Value()).Equal(outEmptySeqDif)
			g.Assert(Chain(inEmpty).Union(nextSeqDif, comp).Value()).Equal(outEmptySeqDif)
			g.Assert(Chain(nil).Union(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(nextSeqDif).Union(nil, comp).Value()).Equal(outEmptySeqDif)
			g.Assert(Chain(inSeq).Union(nil, nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Union(nil, nil).Value()).Equal(inEmpty)
		})
	})

	g.Describe("#Without()", func() {
		g.It("Should return the Seq without all instances of passed value", func() {
			inSeq := Seq{12, 33, 22, 22}
			nonGrata := 44

			inSeqDif := Seq{3, 4, 5}
			nonGrataDif := 4

			outSeq := Seq{12, 33, 22, 22}
			outSeqDif := Seq{3, 5}

			g.Assert(Chain(inSeq).Without(nonGrata, comp).Value()).Equal(outSeq)
			g.Assert(Chain(inSeqDif).Without(nonGrataDif, comp).Value()).Equal(outSeqDif)
			g.Assert(Chain(inSingle).Without(nonGrataDif, comp).Value()).Equal(inSingle)
			g.Assert(Chain(nil).Without(nonGrataDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Without(nonGrataDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inEmpty).Without(nonGrataDif, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Without(nil, comp).Value()).Equal(inEmpty)
			g.Assert(Chain(inSeq).Without(nil, comp).Value()).Equal(inSeq)
			g.Assert(Chain(inSeq).Without(nil, nil).Value()).Equal(inEmpty)
			g.Assert(Chain(nil).Without(nil, nil).Value()).Equal(inEmpty)
		})
	})
}
