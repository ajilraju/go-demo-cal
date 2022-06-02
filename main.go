package main

import (
	"fmt"
	"log"
)

func add(f int, s int) int {
	return f + s
}

func multi(f int, s int) int {
	return f * s
}

func main() {

	var f_op int
	var s_op int

	fmt.Println("Golang Demo calculator")
	fmt.Println("======================\n")
	fmt.Println("Enter the first digit")
	if _, err := fmt.Scan(&f_op); err != nil {
		log.Print("Failed to scan first number", err)
		return
	}

	fmt.Println("Enter the second digit")
	if _, err := fmt.Scan(&s_op); err != nil {
		log.Print("Failed to scan second number", err)
		return
	}

	fmt.Printf("Sum of %d and %d is %d\n", f_op, s_op, add(f_op, s_op))
	fmt.Printf("Product of %d and %d is %d\n", f_op, s_op, multi(f_op, s_op))
}
