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
			inSeq := Seq{ 39, 92, 2, math.MinInt32}
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
			inSeq := Seq{ 39, 92, 2, math.MinInt32}
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
			inSeq := Seq{ 39, 92, 2, math.MinInt32}
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
			inSeq := Seq{ 39, 92, 2, math.MinInt32}
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
			inSeq := Seq{ -12, 2, 6, 33, -12, 8, 92 }
			inSeqDif := Seq{ 3, 2, 1, 0, 99 }


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
			inSeq := Seq{ 39, 92, 2, math.MaxInt32 }
			inSeqDif := Seq{ 7, 9, 0, -12000 }

			g.Assert(Chain(inSeq).Reduce(clctr, 0).Value()).Equal(2147483780)
			g.Assert(Chain(inSeqDif).Reduce(clctr, nil).Value()).Equal(-11984)
			g.Assert(Chain(nil).Reduce(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).Reduce(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Reduce(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#Inject()", func() {
		g.It("Should clone #Reduce behaviour", func() {
			inSeq := Seq{ 39, 34, 3, 230 }
			inSeqDif := Seq{ 7, 230, 0, -1 }

			g.Assert(Chain(inSeq).Inject(clctr, 0).Value()).Equal(306)
			g.Assert(Chain(inSeqDif).Inject(clctr, nil).Value()).Equal(236)
			g.Assert(Chain(nil).Inject(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).Inject(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).Inject(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FoldL()", func() {
		g.It("Should clone #Reduce behaviour", func() {
			inSeq := Seq{ 44, 2, -10, 0 }
			inSeqDif := Seq{ -7, 0, 53, -20 }

			g.Assert(Chain(inSeq).FoldL(clctr, 0).Value()).Equal(36)
			g.Assert(Chain(inSeqDif).FoldL(clctr, nil).Value()).Equal(26)
			g.Assert(Chain(nil).FoldL(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).FoldL(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).FoldL(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#ReduceRigth()", func() {
		g.It("Should return sum of slice elements, starts from right", func() {
			inSeq := Seq{ 39, 2, math.MaxInt32, -12 }
			inSeqDif := Seq{ 8237, 7, 9, 0, -1493, math.MinInt16 }

			g.Assert(Chain(inSeq).ReduceRight(clctr, 0).Value()).Equal(2147483676)
			g.Assert(Chain(inSeqDif).ReduceRight(clctr, nil).Value()).Equal(-26008)
			g.Assert(Chain(nil).ReduceRight(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).ReduceRight(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).ReduceRight(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#FoldR()", func() {
		g.It("Should clone #ReduceRight behaviour", func() {
			inSeq := Seq{ 39, 2, 32 }
			inSeqDif := Seq{ 7 }

			g.Assert(Chain(inSeq).FoldR(clctr, 0).Value()).Equal(73)
			g.Assert(Chain(inSeqDif).FoldR(clctr, nil).Value()).Equal(7)
			g.Assert(Chain(nil).FoldR(clctr, nil).Value()).Equal(nil)
			g.Assert(Chain(inSeq).FoldR(nil, nil).Value()).Equal(nil)
			g.Assert(Chain(nil).FoldR(nil, nil).Value()).Equal(nil)
		})
	})

	g.Describe("#Some()", func() {
		g.It("Should return true if at least one of the elements have passed the predicate test", func() {
			inSeq := Seq{ 77, 34, 242, 99, -20, 0, 3, 0 }
			inSeqDif := Seq{ 3, 1, 1, -12, 0, math.MinInt16, 2 }


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
			inSeq := Seq{ 7, -8, 0, 0 }
			inSeqDif := Seq{ 3, 7, 0, math.MinInt16, 8 }


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
			inSeq := Seq{ 10, 9, 30, 0 }
			inSeqDif := Seq{ 219, 92, 10, 8 }


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
			inSeq := Seq{ 10, 9, 30, 299 }
			inSeqDif := Seq{ 219, 92, -10, 8 }


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
			inSeq := Seq{ 0, 9, 21, 299, 300, math.MaxInt32, math.MaxInt32 + 100 }
			inSeqDif := Seq{ 29, 92, 123, 8, 123 }


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
			inSeq := Seq{ 211, 325, 420, 999, math.MaxInt32, math.MaxInt32 + 100 }
			inSeqDif := Seq{ 1, -8, 910, 0, 19, 2, 3, 8 }


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
			inSeq := Seq{ -100, 0, 100, 32521 }
			inSeqDif := Seq{ 1, 8 }


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
			inSeq := Seq{ -100, 0, 8, 8, 100,  32521 }
			inSeqDif := Seq{ 1, 8, 2, -12, 2, 0 }


			g.Assert(Chain(inSeq).LastIndexOf(8,comp).Value()).Equal(3)
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
			leftSeq := Seq{ 23, 3, 9, 9, 1, -12, math.MinInt32, 9 }
			rightSeq := Seq{ 23, math.MinInt32, 9, 3, 9, 9, 1, -12 }
			rightSeqIdent:= Seq{ 23, 3, 9, 9, 1, -12, math.MinInt32, 9 }

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
			leftSeq := Seq{ 23, 3, 9, 9, 1, -12, math.MinInt32, 9 }
			rightSeq := Seq{ 23, math.MinInt32, 9, 3, 9, 9, 1, -12 }
			rightSeqIdent:= Seq{ 23, 3, 9, 9, 1, -12, math.MinInt32, 9 }

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
