package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num uint64 = 12312312323123
	var any1 any = num
	switch v := any1.(type) { // type switch
	case int, uint8, uint16, uint32, uint64, int8, int32, int64, uint:
		fmt.Println(v, "is", reflect.TypeOf(v), "type")
	case float64:
		println(v, "is float64 type")
	case bool:
		println(v, "is bool type")
	case string:
		println(v, "is string type")
	default:
		println(v, "does not belong to int,float64,bool or string")
	}
}
