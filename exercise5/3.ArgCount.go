// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

// func main() {
// 	args := os.Args
// 	l := len(args) - 1

// 	if l == 0 {
// 		fmt.Println("give me args")
// 	} else if l == 1 {
// 		fmt.Printf("There is one: %s\n", args[1])
// 	} else if l == 2 {
// 		fmt.Printf("There are two: %s %s\n", args[1], args[2])
// 	} else if l == 5 {
// 		fmt.Println("There are 5 arguments")
// 	} else {
// 		fmt.Println("gotcha!")
// 	}
// }
