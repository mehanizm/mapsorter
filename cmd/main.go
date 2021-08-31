package main

import (
	"fmt"

	"github.com/mehanizm/mapsorter"
)

func main() {
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
}
