package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(0)
	r := Random()
	fmt.Printf("random generated by using cgo: %d\n", r)
}
