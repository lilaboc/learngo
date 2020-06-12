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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.

	for _, v := range strings.Split(header, separator) {
		fmt.Printf("%-15s", v)
	}
	fmt.Printf("\n")

	locations := []string{}
	sizes := []int{}
	beds := []int{}
	baths := []int{}
	prices := []int{}
	fmt.Println("===========================================================================")
	sumSize := 0
	sumBed := 0
	sumBaths := 0
	sumPrice := 0
	for _, row := range strings.Split(data, "\n") {
		data := strings.Split(row, separator)
		locations = append(locations, data[0])
		size, _ := strconv.Atoi(data[1])
		sizes = append(sizes, size)
		sumSize += size
		bed, _ := strconv.Atoi(data[2])
		beds = append(beds, bed)
		sumBed += bed
		bath, _ := strconv.Atoi(data[3])
		baths = append(baths, bath)
		sumBaths += bath
		price, _ := strconv.Atoi(data[4])
		prices = append(prices, price)
		sumPrice += price

	}
	for i := 0; i < len(locations); i++ {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}
	fmt.Println("===========================================================================")
	fmt.Printf("%-15s%-15.2f%-15.2f%-15.2f%-15.2f\n", "", float64(sumSize)/float64(len(sizes)), float64(sumBed)/float64(len(beds)), float64(sumBaths)/float64(len(baths)), float64(sumPrice)/float64(len(prices)))

}
