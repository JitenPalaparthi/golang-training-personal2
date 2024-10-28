package main

import "fmt"

func main() {
	type integer = int
	var num uint64 = 1000
	var num1 integer = 2000
	var any1, any2 any = num, num1

	var sum float64 = 0 // any1 + any2

	switch v := any1.(type) {
	case int:
		sum += float64(v)
	case int8:
		sum += float64(v)
	case int16:
		sum += float64(v)
	case int32:
		sum += float64(v)
	case int64:
		sum += float64(v)
	case uint8:
		sum += float64(v)
	case uint16:
		sum += float64(v)
	case uint32:
		sum += float64(v)
	case uint64:
		sum += float64(v)
	case uint:
		sum += float64(v)
	case float32:
		sum += float64(v)
	case float64:
		sum += v
	default:
		sum += 0
	}

	switch v := any2.(type) {
	case int:
		sum += float64(v)
	case int8:
		sum += float64(v)
	case int16:
		sum += float64(v)
	case int32:
		sum += float64(v)
	case int64:
		sum += float64(v)
	case uint8:
		sum += float64(v)
	case uint16:
		sum += float64(v)
	case uint32:
		sum += float64(v)
	case uint64:
		sum += float64(v)
	case uint:
		sum += float64(v)
	case float32:
		sum += float64(v)
	case float64:
		sum += v
	default:
		sum += 0
	}

	fmt.Printf("sum:%.3f", sum)

}
