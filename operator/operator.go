package main

import "fmt"

func main() {
	const A int = 1
	const B int = 2

	fmt.Println(A + B)
	fmt.Println(A - B)
	fmt.Println(A * B)
	fmt.Println(A / B)
	fmt.Println(A % B)

	const y int = 3
	const z int = 3

	fmt.Println(y == z)
	fmt.Println(y != z)
	fmt.Println(y < z)
	fmt.Println(y > z)
	fmt.Println(y <= z)
	fmt.Println(y >= z)

}
