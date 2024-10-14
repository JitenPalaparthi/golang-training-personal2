package main

import "fmt"

func main() {

	var (
		num1 uint8 = 131
		//num2 uint64 = num1 // not okay, there is no implicit type casting

		num2 uint64 = uint64(num1) // type casting
	)
	println("num2:", num2)

	var (
		num3 uint64 = 123123123145
		num4 uint8  = uint8(num3) // 8 bytes are trying to fit to one byte
		num5 uint16 = uint16(num3)

		float1 float32 = 2131231.123

		num6 uint32 = uint32(float1)

		// ok1 bool = true // 1 byte

		//num7 int8 = int8(ok1) // cannot convert non numbers to numbers

		byte1 byte = 123 // byte is nothing but uint8

		num7 int64 = int64(byte1)

		char1 rune = '挨' // rune is nothing but int32

		num8 int64 = int64(char1)
	)
	// if ok1 { // nothing but ok1 == true
	// 	num7 = 1
	// } else {
	// 	num7 = 0
	// }

	println("num4:", num4)
	println("num5:", num5)
	println("num6:", num6)
	println("num7:", num7)
	println("num7:", num8)

	c1 := complex(123.234, 123.23) // c1 could be complex64 or complex128

	var r1, i1 float32 = 12.123, 123.343

	c2 := complex(r1, i1)

	c3 := c1 + complex128(c1)

	var c4 complex64 = complex64(c1) - c2

	c5 := c1 * complex128(c2)

	c6 := 1231231.34 + 231.34i // c6 is a complex128 number .. without a complex functin, can create a complex number with
	// short hand declaration r+i

	fmt.Println(c1, c2, c3, c4, c5, c6)

}

// long num1 = 0; // 8 bytes

// short num2 = 12312; // 2 bytes

// num1 = num2

// type casting or conversion

// can convert / cast number of any type to numbers of any type
