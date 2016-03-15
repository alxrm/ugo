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

	g.Describe("#Min()", func() {
		g.It("Should return min value by chaining ", func() {
			inSeq := Seq{2, 99, -12, 884, 8}
			inSeqDif := Seq{1, 0, -math.MaxInt32, -math.MinInt16}
			inSingle := Seq{0}
			inEmpty := Seq{}
			comp := func(l, r Object) int {
				return l.(int) - r.(int)
			}

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
			inSingle := Seq{0}
			inEmpty := Seq{}
			comp := func(l, r Object) int {
				return l.(int) - r.(int)
			}

			g.Assert(Chain(inSeq).Max(comp).Value()).Equal(-math.MinInt32)
			g.Assert(Chain(inSeqDif).Max(comp).Value()).Equal(1)
			g.Assert(Chain(inSingle).Max(comp).Value()).Equal(0)
			g.Assert(Chain(inEmpty).Max(comp).Value()).Equal(-1)
			g.Assert(Chain(nil).Max(comp).Value()).Equal(-1)
			g.Assert(Chain(inSeqDif).Max(nil).Value()).Equal(-1)
			g.Assert(Chain(nil).Max(nil).Value()).Equal(-1)
		})
	})
}
