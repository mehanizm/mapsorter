Mapsorter
================

[![GoDoc](https://godoc.org/github.com/mehanizm/mapsorter?status.svg)](https://pkg.go.dev/github.com/mehanizm/mapsorter)
![Go](https://github.com/mehanizm/mapsorter/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/mehanizm/mapsorter/branch/master/graph/badge.svg)](https://codecov.io/gh/mehanizm/mapsorter)
[![Go Report](https://goreportcard.com/badge/github.com/mehanizm/mapsorter)](https://goreportcard.com/badge/github.com/mehanizm/mapsorter)

Golang map sorter that frees you of writing boilerplate sorting code every time

- [Mapsorter](#mapsorter)
  - [Installation](#installation)
  - [Basic usage](#basic-usage)
    - [Using options](#using-options)
  - [Benchmarking](#benchmarking)

## Installation

```
go get github.com/mehanizm/mapsorter
```

## Basic usage

You can see an example in cmd/main.go

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
```

### Using options

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

## Benchmarking

You can see some benchmark tests comparing mapsorter with straightforward boilerplate go code. Them shows approximately 5 times decreasing of all metrics.
But save time, planet and keys on your keyboard! :)

With love ❤️