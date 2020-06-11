// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{"a", "b", "c"}
	distances := [...]int{1, 2, 3, 4, 5}
	data := [...]byte{1, 2, 3, 4, 5}
	ratios := [...]float32{1.2}
	alives := [...]bool{true, false, true, false}
	zero := [...]byte{}

	for _, v := range names {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	for _, v := range distances {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for _, v := range data {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for _, v := range ratios {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for _, v := range alives {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	for _, v := range zero {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
