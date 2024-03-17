package main

import "fmt"

// ---scanf for clint input---

// %s for string
// %d for int
// %f for float

func main() {
	var firstName string
	var lastName string
	var grade float32

	var x string = "กรุณาป้อนชื่อของท่าน"
	var y string = "กรุณาป้อนนามสกุลของท่าน"
	var z string = "กรุณาป้อนเกรดของท่าน"

	fmt.Print(x)
	fmt.Scanf("%s", &firstName)
	fmt.Print(y)
	fmt.Scanf("%s", &lastName)
	fmt.Print(z)
	fmt.Scanf("%d", &grade)
	fmt.Println("ชื่อ", firstName, lastName, "เกรด : ", grade)
}
