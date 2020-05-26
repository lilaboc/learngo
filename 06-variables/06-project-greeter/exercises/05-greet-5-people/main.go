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
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet 5 People
//
//  Greet 5 people this time.
//
//  Please do not copy paste from the previous exercise!
//
// RESTRICTION
//  This time do not use variables.
//
// INPUT
//  bilbo balbo bungo gandalf legolas
//
// EXPECTED OUTPUT
//  There are 5 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Hello great gandalf !
//  Hello great legolas !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	fmt.Println("There are", len(os.Args)-1, "people!")
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Hello great", os.Args[i], "!")
	}
	fmt.Println("Nice to meet you all.")

}
