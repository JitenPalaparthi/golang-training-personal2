package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x any = rand.Intn(100000)
	fmt.Println(x)

	var y = 1231231

	var z = 432423

	var i uint32 = uint32(y)

	sum := x.(int) + y + z + int(i)

	var x1 any = 12313

	sum1 := x.(int) + x1.(int)

	println("sum:", sum)
	println("sum1:", sum1)

	var x2 any = 12312.12312 // float64

	//var num1 int = int(x2.(float64))

	// var num1 int = x2.(int)

	num1, ok1 := x2.(int)

	// if ok1==true {

	if ok1 { // same as above
		println("num1 square:", num1*num1)
	} else {
		num1, ok1 := x2.(float64)
		if ok1 {
			fmt.Printf("num1 square:%.2f\n", (num1 * num1))
		} else {
			println("x2 is neither int not float64")
		}
	}
}

// type assertion
// type assertion is only for empty interface or any. For other type casting, have to use type casting as usual

// any.(type)
