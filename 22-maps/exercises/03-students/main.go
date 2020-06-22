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
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	if len(os.Args) == 1 {
		fmt.Println("gimme a house name")
		return
	}
	house := os.Args[1]
	houses := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},

		"slytherin": {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":      {"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	if val, ok := houses[house]; ok {
		clone := append([]string(nil), val...)
		sort.Strings(clone)

		fmt.Printf("~~~ %s students ~~~\n\n", house)
		for _, student := range clone {
			fmt.Printf("\t+ %s\n", student)
		}

	} else {
		fmt.Printf("Sorry. I don't know anything about \"%s\".\n", house)
	}
}
