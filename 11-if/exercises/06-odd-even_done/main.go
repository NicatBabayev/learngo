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
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("\"%v\" is not a number", os.Args[1])
	}
	if num%2 == 0 {
		if num%8 == 0 {
			fmt.Println("Number is even and divisible by 8")
		} else {
			fmt.Println("Number is even but is not divisible by 8")
		}
	} else {
		fmt.Println("Number is odd.")
	}
}
