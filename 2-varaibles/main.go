package main

import "fmt"

type integer = int // this is an alias for int

func main() {
	var num1 int = -123123
	println("num1:", num1)

	var num2 uint = 31312
	println("num2:", num2)

	var num3 uint8 = 39

	var num4 float32 = 12312.123

	var num5 float64 = 123123.12312

	fmt.Println("num3:", num3, "num4:", num4, "num5:", num5)

	fmt.Printf("num3:%d num4:%f num5:%.2f\n", num3, num4, num5)

	var age = 45 // type inference

	var float1 = 12.34 // it could be float32 but it is inferred as float64

	var ok = false

	var ok1 = true // type inference

	var ok2 bool // false is inferred

	var num6 uint64 // 0 is inferred

	var num7 float32 // 0 is inferred

	fmt.Printf("age:%v type of age: %T\n", age, age)
	fmt.Printf("float1:%v type of float1: %T\n", float1, float1)

	fmt.Println(ok1, ok2, num6, num7)

	// shorthand way of creating variable
	ok3 := true // for shorthand notation , a value must be given

	fmt.Printf("ok3:%v type of ok3:%T\n", ok3, ok3)

	{
		var (
			num1       uint8
			num2       uint64
			num3       float32
			num4       float64
			ok         bool
			num5, num6 int = 100, 200
		)

		fmt.Println(num1, num2, num3, num4, num5, num6, ok)
	}
	fmt.Println(num1, num2, num3, num4, num5, num6, ok)

	var char1 rune = 95
	var char2 rune = 'A'
	var char3 int32 = '界'
	var char4 rune = '\u1200'

	var num9 integer = 12312

	type integer = int // this is an alias for int

	type char = rune

	var num8 integer = 12312

	var char5 char = 'B'

	println(num8, num9)

	var b1 byte = 123

	println(char1, char2, char3, char4, char5, b1)

}

/* numbers
   int,uint,int8,int16,int32,int64 ,uint8,uint16,uint32,uint64
   float32, float64
   rune, byte
   uintptr
*/
/* boolean
bool
*/

/* string
string -> utf8 chars
*/

/* empty interface
interface{} or any --> both of them are same
*/
