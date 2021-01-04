package main

import "fmt"

func variableParameters(v ...interface{}) {
	fmt.Println(v...)
	for i, a := range v {
		fmt.Printf("\t%d: %v\n", i, a)
	}
}

// 可变参数
func main() {
	fmt.Println("\nvariable parameters with empty parameters")
	variableParameters()
	fmt.Println("\nvariable parameters with primitive types")
	variableParameters(1, nil, "hello")
	fmt.Println("\nvariable parameters with construction types")
	variableParameters(1, nil, []int{1, 3, 6})
}
