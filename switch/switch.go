package main

import (
	"fmt"
	"math/rand"
	"time"
)

func returnDynamicType() interface{} {
	rets := []interface{}{
		true,
		"demo",
		43,
		3.14,
		nil,
	}
	rand.Seed(time.Now().UnixNano())
	return rets[rand.Intn(len(rets))]
}

func main() {
	var t interface{}
	t = returnDynamicType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type: %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case float64:
		fmt.Printf("float %f\n", t) // t has type float64
	case string:
		fmt.Printf("string %v\n", t) // t has type string
	}
}
