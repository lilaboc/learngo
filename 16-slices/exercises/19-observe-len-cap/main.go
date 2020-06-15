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
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//
	//games := []string{}
	// 2. print the length and capacity of the games slice
	//fmt.Printf("length: %d\n", len(games))
	//fmt.Printf("capacity: %d\n", cap(games))
	//
	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	// 4. print the length and capacity of the games slice
	//
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//games = append(games, "pacman")
	//games = append(games, "mario")
	//games = append(games, "tetris")
	//games = append(games, "doom")
	//
	// 6. print the length and capacity of the games slice
	//fmt.Printf("length: %d\n", len(games))
	//fmt.Printf("capacity: %d\n", cap(games))
	//
	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)

	games := []string{"pacman", "mario", "tetris", "doom"}
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	for i := 0; i < 4; i++ {
		fmt.Println(games[i])
	}

	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	for i := 0; i < 4; i++ {
		s := games[:i]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}
	// for ... {
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", ...)
	// }

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]
	//
	// 2. print the games and the new slice's len and cap
	fmt.Printf("games's len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's len: %d cap: %d\n", len(zero), cap(zero))

	//
	// 3. append a new element to the new slice
	zero = append(zero, "botw")
	//
	// 4. print the new slice's lens and caps
	fmt.Printf("zero's len: %d cap: %d\n", len(zero), cap(zero))
	//
	// 5. repeat the last two steps 5 times (use a loop)
	for i := 0; i < 5; i++ {
		zero = append(zero, "botw")
		fmt.Printf("zero's len: %d cap: %d\n", len(zero), cap(zero))
	}
	//
	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	to := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	for _, v := range to {
		zero = append(zero, v)
	}

	//fmt.Printf("games's len: %d cap: %d\n",  len(games), cap(games))
	fmt.Printf("zero's len: %d cap: %d\n", len(zero), cap(zero))
	fmt.Println()

	// zero := ...
	// fmt.Printf("games's     len: %d cap: %d\n", ...)
	// fmt.Printf("zero's      len: %d cap: %d\n", ...)

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	// zero = ...
	// for ... {
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", ...)
	// }

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
}
