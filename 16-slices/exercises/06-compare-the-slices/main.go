// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	names := strings.Split(namesA, ", ")
	sort.Strings(names)
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sort.Strings(namesB)
	equal := true
	if len(names) == len(namesB) {
		for i, v := range names {
			if v != namesB[i] {
				equal = false
				break
			}
		}
	} else {
		equal = false
	}
	if equal {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}
}
