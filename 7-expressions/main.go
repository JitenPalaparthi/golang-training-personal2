package main

func main() {
	a := 100

	b := 200

	c := 300

	// (a+b) whole square
	// a square+ b square+ 2ab

	d := 2*a*b + (a * a) + (b * b)

	println("a+b whole square:", d)

	// t := a
	// a = b
	// b = t

	//a, b = b, a // direct swaping in go

	a, b, c = c, a, b // swapping three variables

	println(a, b, c)

	// cant do this unless use anonymous funcs
	// var e :={
	// 	return a * a
	// }

	//println(e)

	k1 := a == b
	k2 := a >= b
	k3 := a <= b

	println(k1, k2, k3)

	k4 := (a+b) <= 500 && k2 || a <= c
	// (400)<=500 = true
	// true && true || false
	// true || false
	// true

	println(k4)

	println((a+b) <= 500 && k2 || a <= c)

	// if k4 {
	// 	println(true)
	// } else {
	// 	println(false)
	// }
}

// precedence

// ()
// *,/, %, <<, >> , &,|, &^
// +, - , `
// == != , <,<=,>,>=
// && , ||
//
