package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Sudoku Project!")

	test := [3]int{5, 3, 4}

	for t := range test {
		fmt.Println(t)
	}
}
