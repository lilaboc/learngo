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
// EXERCISE: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 42
//
//     (! You should use redeclaration)
//
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  42, 20, 3.14
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	//
	age, yourAge := 42, 40
	age, yourAge, ratio := 42, 20, 3.14

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(age, yourAge, ratio)
}
