package main

func main() {

	// var num1 uint8 = 100

	// var num2 int = 200

	// sum := int(num1) + num2

	// ok := int(num1) <= num2

	var num1 int = 12313

	if true {

	}

	if num1%2 == 0 {
		println(num1, "is even number")
	} else {
		println(num1, "is odd number")
	}

	var (
		age    uint8
		gender rune
	)

	// if age is greater than equal to 18 and gendre is female , then eligible for marriage
	// if age is greater then equal to 21 and gender is male , then eligible for marriage
	// not eligible to marry

	age = 17
	gender = 'F'

	// false && (false || true)
	// false && true
	// false

	if age >= 18 && (gender == 'f' || gender == 'F') {
		println("she is eligible to marry, because her age is", age)

		// fasle && (false || false)
		// false && false
		// false

	} else if age >= 21 && (gender == 'm' || gender == 'M') {
		println("he is eligible to marry,becase his age is", age)
	} else {

		// fasle || true
		// true
		if gender == 'f' || gender == 'F' {
			println("she is not elible to marry")
		} else if gender == 'm' || gender == 'M' {
			println("he is not elible to marry")
		} else {
			println("unknown gender but not eligible to marry")
		}
	}
}
