Mapsorter
================

[![GoDoc](https://godoc.org/github.com/mehanizm/mapsorter?status.svg)](https://pkg.go.dev/github.com/mehanizm/mapsorter)
![Go](https://github.com/mehanizm/mapsorter/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/mehanizm/mapsorter/branch/master/graph/badge.svg)](https://codecov.io/gh/mehanizm/mapsorter)
[![Go Report](https://goreportcard.com/badge/github.com/mehanizm/mapsorter)](https://goreportcard.com/badge/github.com/mehanizm/mapsorter)

Golang map sorter that frees you of writing boilerplate sorting code every time

When can it be useful? Imagine, you have a map with data and want to get keys slice that:

* in order of keys by string length, string value, datetime, int or float converted from strings or as native types;
* in order of values by string length, string value, datetime, int or float converted from strings or as native types;
* in reverse order or only top N of the sorted results.

Not clear enough? See examples below.

## Table of contents

- [Mapsorter](#mapsorter)
	- [Installation](#installation)
	- [Basic usage](#basic-usage)
		- [The functional style](#the-functional-style)
			- [Using options](#using-options)
		- [The struct style](#the-struct-style)
			- [Using options](#using-options-1)
	- [Benchmarking](#benchmarking)

## Installation

```
go get github.com/mehanizm/mapsorter
```

## Basic usage

You can see an example in cmd/main.go

The lib has two different api. Please, use whichever is more convenient for you.

### The functional style

```go
in := map[int]string{
	1: "a",
	2: "a",
	4: "c",
	3: "b",
}
sortedKeys, err := mapsorter.Sort(in, mapsorter.ByKeys, mapsorter.AsInt, true, 3)
if err != nil {
	panic(err)
}
for _, key := range sortedKeys {
	fmt.Println(key)
}
// >> 4
// >> 3
// >> 2
```

#### Using options

Sort function has four extending options.

`SortBy`
* ByKeys
* ByValues

`SortAs`
* AsString
* AsStringByLength
* AsInt
* AsFloat
* AsDatetime

These options are defined as enum package constants for easy using.

And two more:

`reverse` – bool flag to choose reverse sorting order if needed.

`top` – int count of top result to return.

### The struct style

```go
in := map[int]string{
	1: "a",
	2: "a",
	4: "c",
	3: "b",
}
sortedKeys, err := mapsorter.Map(in).ByValues().AsStringByLength().Reverse().Top(1).Sort()
if err != nil {
	panic(err)
}
for _, key := range sortedKeys {
	fmt.Println(key)
}
// >>  4
```

Or you can use `MustSort()` wrapper that panic if there is an internal error, be careful. But not so verbose.
```go
for _, key := range mapsorter.Map(in).ByValues().AsStringByLength().Reverse().Top(1).MustSort() {
	fmt.Println(key)
}
// >>  4
```

#### Using options

All options can be changes with struct methods:

* `ByKeys()` – sort by keys,
* `ByValues()` – sort by values,
* `AsString()` – sort as strings,
* `AsStringByLength()` – sort as string lengths,
* `AsInt()` – sort as ints,
* `AsFloat()` – sort as floats,
* `AsDatetime()` – sort as datetime (with smart conversion),
* `Forward()` – forward order,
* `Reverse()` – reverse order,
* `Top(top)` – top N of sorted results,
* `All()` – all results.

## Benchmarking

You can see some benchmark tests comparing mapsorter with straightforward boilerplate go code. Them shows approximately 5 times decreasing of all metrics.
But save time, planet and keys on your keyboard! :)

With love ❤️