[![GoDoc](https://godoc.org/github.com/alxrm/ugo?status.svg)](https://godoc.org/github.com/alxrm/ugo)
[![Go Report Card](https://goreportcard.com/badge/github.com/alxrm/ugo)](https://goreportcard.com/report/github.com/alxrm/ugo)
[![GoCover.io](https://img.shields.io/badge/coverage-99.7%-brightgreen.svg?style=flat)](https://gocover.io/github.com/alxrm/ugo)
[![Travis CI](https://travis-ci.org/alxrm/ugo.svg?branch=master)](https://travis-ci.org/alxrm/ugo)

# ugo

Simple and expressive toolbox written with love and care in Go.

Deeply inspired by [underscore.js](http://underscorejs.org/) and has the same syntax and behaviour

Fully covered with tests, no surprises

### Quick start

__Installation__

```
go get -u  github.com/alxrm/ugo
```

__Import__

``` GO
import (
	u "github.com/alxrm/ugo"
)
```

### Usage

It works with some special type, named `Seq`, which is an alias for `[]interface{}`

So let's make some

```GO
	
// creates the new string Seq
strSeq := u.Seq{ "nineteen", "eigth", "four" } // ["nineteen" "eigth" "four"]
	
// creates the new Seq with given length
emptySeq := u.NewSeq(0); // []
	
// copies reflectively all of the values from int slice to Seq
intSlc := []int{ 4, 6, 2, 7, 8, 10, 9, 9, 120, 10, 2, 17 }
intSeq := u.From(intSlc, len(intSlc)) // [4 6 2 7 8 10 9 9 120 10 2 17]

```

_Actually reflection approach in `u.From` can have an impact on performance :snail: so use it with care._

Okay, now we have some Seq's, 

What if I want to leave only unique elements in my int slice?

Go and get it :zap:

```Go

initialSeq := u.Seq{ 4, 6, 2, 7, 8, 10, 9, 9, 120, 10, 2, 17 }
uniquedSeq := u.Uniq(initial, func(l, r u.Object) int { return l.(int) - r.(int) })

fmt.Println(uniquedSeq) // [4 6 2 7 8 10 9 120 17]

```

And I want to leave only odd ones

Nice choice! :fire:

```Go
oddsSeq := u.Filter(uniquedSeq, func(cur, _, _ u.Object) bool { return cur.(int) % 2 != 0 })

fmt.Println(oddsSeq) // [7 9 17]
```

### Try it by yourself! 

Explore all of the features and get your slice routine done faster

### Contribute

We highly appreciate any of your pull requests and issues!

Be welcome :coffee:


__Some of the TODOs:__

- [ ] Aliases from underscore
- [ ] Chaining like map(...).uniq(...).result()
- [ ] Better documentation 


### License 

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
