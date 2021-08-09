package main

import "fmt"

func main() {
	//Initializing
	a := 3
	b := 5

	//show value before swap
	fmt.Println("Before swap :")
	fmt.Println("a = ",a )
	fmt.Println("b = ",b )

	//swap value and show result
	a,b = b,a
	fmt.Println("\nAfter swap :")
	fmt.Println("a = ",a )
	fmt.Println("b = ",b )
}
