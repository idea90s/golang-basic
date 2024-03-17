package main

import "fmt"

func main() {
	fmt.Println("Hello, Golang! ...") // output

	// create constants variables
	const money string = "ThaiBath"

	fmt.Println(money) //output constants

	// create variables
	var firstName string = "Chwit" // string
	lastName := "Chwit Ruangsiri"  // string
	var age int = 27               // int || number
	score := 112.21                // int || number
	var isActive bool = false      // boolean
	isPass := true                 // boolean

	fmt.Println("ชื่อ", firstName, lastName) //output
	fmt.Println("อายุ", age)                 //output
	fmt.Println("มีคะแนน", score)            //output
	fmt.Println("สถานะ", isActive)           //output
	fmt.Println("เป็นสมาชิกหรือไม่", isPass) //output

	fmt.Printf("ชื่อจริงของ %v", firstName) // output and add values variables

	fmt.Printf("Type Of %T", score) // check type variables
}
