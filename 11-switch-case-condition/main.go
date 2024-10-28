package main

func main() {

	var num = 31

	switch { // empty switch , when cases are conditions, then the switch should be empty switch
	case num%2 == 0:
		println(num, "is an even number")
	default:
		println(num, "is an odd number")
	}

	// another example
	num = 50 //mutating a variable
	println(num >= 50 && num < 100 && true && (true || false) && false)
	switch { // empty switch
	case num < 0:
		println(num, "is less than 0")
	case num >= 0 && num < 50:
		{
			println(num, "is between 0 and 50")
		}
	//case num >= 50 && num < 100 && true && (true || false) && true && (100 == (50 * 2)):
	case num >= 50 && num < 100:
		// true && true && true && true && true
		// true && true && true
		// true && true
		// true
		println(num, "is between 50 and 100")

	default:
		println(num, "is greater than or equal to 100")
	}

}
