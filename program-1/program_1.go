package main

import "fmt"

func swapvalue(a, b int) (int, int) {
	a, b = b, a

	return a, b
}

func main() {
	//Initializing
	a := 3
	b := 5

	//show value before swap
	fmt.Println("Before swap :")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//swap value and show result
	a, b = swapvalue(a, b)

	// show value after swap
	fmt.Println("\nAfter swap :")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
