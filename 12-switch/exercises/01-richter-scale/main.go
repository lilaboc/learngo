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
	"strconv"
)

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Give me the magnitude of the earthquake")
		return
	}
	if magnitude, err := strconv.ParseFloat(os.Args[1], 32); err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	} else {
		level := ""
		switch {
		// Less than 2.0                micro
		case magnitude <= 2.0:
			level = "micro"
		// 2.0 and less than 3.0        very minor
		case magnitude <= 3.0:
			level = "very minor"
		// 3.0 and less than 4.0        minor
		case magnitude <= 4.0:
			level = "minor"
		// 4.0 and less than 5.0        light
		case magnitude <= 5.0:
			level = "light"
		// 5.0 and less than 6.0        moderate
		case magnitude <= 6.0:
			level = "moderate"
		// 6.0 and less than 7.0        strong
		case magnitude <= 7.0:
			level = "strong"
		// 7.0 and less than 8.0        major
		case magnitude <= 8.0:
			level = "major"
		// 8.0 and less than 10.0       great
		case magnitude <= 10.0:
			level = "major"
		// 10.0 or more                 massive
		case magnitude > 10.0:
			level = "massive"
		}
		fmt.Printf("%g is %s", magnitude, level)
	}
}
