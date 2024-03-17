package main

import "fmt"

func main() {

	var a string = "string" // "string"
	var b int = 9           // int , int8 , int16 , int32 , int64 , uint8 , uint16 , uint32 , uint64 , uintptr
	var c float32 = 88.88   // float32 , float64
	var d bool = false      // boolean

	fmt.Printf("a Type = %T\n", a)
	fmt.Printf("b Type = %T\n", b)
	fmt.Printf("c Type = %T\n", c)
	fmt.Printf("d Type = %T\n", d)
}
