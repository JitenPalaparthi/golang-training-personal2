package main

import (
	"fmt"
	"reflect"
)

//type any = interface{}

func main() {

	//var any1 interface{}

	var x any // x can hold any type

	x = "Hello world"

	fmt.Println(x)

	fmt.Println("type of x:", reflect.TypeOf(x))

	x = true

	fmt.Println(x)

	fmt.Println("type of x:", reflect.TypeOf(x))

	x = 12312
	fmt.Println(x)

	fmt.Println("type of x:", reflect.TypeOf(x))

	x = 1233122.32

	fmt.Println(x)

	fmt.Println("type of x:", reflect.TypeOf(x))

	var num uint8 = 123

	x = num

	fmt.Println(x)

	fmt.Println("type of x:", reflect.TypeOf(x))

	var y interface{} = x

	fmt.Println(y)

	fmt.Println("type of y:", reflect.TypeOf(y))

	var z any

	fmt.Println(z)

	fmt.Println("type of z:", reflect.TypeOf(z))

}

// empty interface interface{}
// any

// Usecases:

// 1. Generic data structures
// 2. Dynamic type handling
