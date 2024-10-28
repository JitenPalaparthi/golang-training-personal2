package main

func main() {

	num := 6
	println("Successful fallthrough")
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("Failure and false negative fallthrough")
	num = 6
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough // since there is fallthrough whether the next case is satisfied or not, the statements are executed
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

}

// 1. only one case condition is evaluated.
// 2. want the same effect what if break is removed in other programming languages, have to include fallthrough in switch case
// 3. However fallthrough can also give false negative, when you dont write then order or proper cases
