package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var val1 string = "12312"

	//var num1 uint32 = uint32(val1) // not possible

	num1, _ := strconv.Atoi(val1)

	fmt.Printf("Value %v Type:%T\n", num1, num1)

	// What if error

	var val2 string = "12a312"

	//var num1 uint32 = uint32(val1) // not possible

	num2, err := strconv.Atoi(val2)

	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Value %v Type:%T\n", num2, num2)
	}

	val3 := "12312.123"

	float1, _ := strconv.ParseFloat(val3, 32)

	fmt.Println("Value:", float1, "Type:", reflect.TypeOf(float1))

	val4 := "T"

	ok1, _ := strconv.ParseBool(val4)

	fmt.Println("Value:", ok1, "Type:", reflect.TypeOf(ok1))

	//ParseInt(s, 10, 0)

	var num3 = 12312

	val5 := strconv.Itoa(num3)

	fmt.Println("Value:", val5, "Type:", reflect.TypeOf(val5))

	val6 := fmt.Sprintf("num1:%d float1: %f ok1:%v", num1, float1, ok1)

	fmt.Println("Value:", val6, "Type:", reflect.TypeOf(val6))

	//val7 := fmt.Sprint(num1)

}

// slice, map , pointer, interface, channel
