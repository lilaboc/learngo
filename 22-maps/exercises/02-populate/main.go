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
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 0 from Product ID #576872813.
//
// ---------------------------------------------------------

func main() {
	const name = "dulin"
	phones := map[string]string{
		"bowen": "202-555-0179",
		"dulin": "03.37.77.63.06",
		"greco": "03489940240",
	}
	if _, ok := phones[name]; ok {
		fmt.Printf("%s's phone number: %s\n", name, phones["dulin"])
	} else {
		fmt.Printf("%s's phone number: is not available\n", name)

	}
	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	availability := map[string]bool{
		"617841573": true,
		"879401371": false,
		"576872813": true,
	}
	const id = "879401371"
	_, ok := availability[id]
	if ok {
		fmt.Printf("Product ID #%s is available\n", id)
	} else {
		fmt.Printf("Product ID #%s is not available\n", id)
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	mphone := map[string][]string{
		"bowen": {"202-555-0179"},
		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"greco": {"03489940240", "03489900120"},
	}
	found := false
	const sname = "greco"
	if val, ok := mphone[sname]; ok {
		if len(val) > 1 {
			found = true
			fmt.Printf("%s's 2nd phone number: %s\n", sname, val[1])
		}
	}
	if !found {
		fmt.Printf("%s's 2nd phone number: N/A\n", sname)
	}

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	//     100 [617841573:4 576872813:2]
	//     101 [576872813:5 657473833:20]
	//     102 []
	//
	basket := map[int]map[int]int{
		100: {
			617841573: 4,
			576872813: 2,
		},
		101: {
			576872813: 5,
			657473833: 20,
		},
		102: {},
	}
	customer := 101
	product := 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", customer, basket[customer][product], product)

}
