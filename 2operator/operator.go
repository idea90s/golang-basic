package main

import "fmt"

func main() {
	const A int = 1 // constant variable
	const B int = 2 // constant variable

	fmt.Println(A + B) // output
	fmt.Println(A - B) // output
	fmt.Println(A * B) // output
	fmt.Println(A / B) // output
	fmt.Println(A % B) // output

	var y int = 3 // variable
	var z int = 3 // variable

	fmt.Println(y == z) // output
	fmt.Println(y != z) // output
	fmt.Println(y < z)  // output
	fmt.Println(y > z)  // output
	fmt.Println(y <= z) // output
	fmt.Println(y >= z) // output

}
