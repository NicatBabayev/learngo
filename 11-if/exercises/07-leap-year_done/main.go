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
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	input := os.Args
	if len(input) == 1 {
		fmt.Println("Give me a year number")
		os.Exit(3)
	}
	if year, err := strconv.Atoi(input[1]); err != nil {
		fmt.Printf("%v is not a valid number", year)
	} else {
		if year%4 == 0 {
			if year%100 == 0 {
				if year%400 == 0 {
					fmt.Println(year, "is a leap year")
				} else {
					fmt.Println(year, "is not a leap year")
				}
			} else {
				fmt.Println(year, "is a leap year")
			}
		} else {
			fmt.Println(year, "is a leap year")
		}
	}

}