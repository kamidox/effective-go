package main

import "fmt"

func update(a []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 1
	}
	return a
}

func allocateWithNew() {
	fmt.Println("allocate with new(), it return Array Pointer")
	a := []int{
		0,
		1,
		2,
		3,
		4,
	}

	fmt.Printf("before: %v\n", a)
	update(a)
	fmt.Printf("after: %v\n", a)
}

func allocateWithMake() {
	fmt.Println("\nallocate with make(), use Array as parameter. Array are passed by Value, not by Reference")
	a := make([]int, 5)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Printf("before: %v\n", a)
	update(a)
	fmt.Printf("after: %v\n", a)
}

func allocateWithMakeBySlice() {
	fmt.Println("\nallocate with make(), but use Slice as parameter")
	a := make([]int, 5)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	fmt.Printf("before: %v\n", a)
	update(a[0:])
	fmt.Printf("after: %v\n", a)
}

/**
 * 1. make() applies only to maps, slices and channels and does not return a pointer.
 * 2. new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T.
 *
 * 1. Arrays are values. Assigning one array to another copies all the elements.
 * 2. In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
 * 3. The size of an array is part of its type. The types [10]int and [20]int are distinct.
 */
func main() {
	allocateWithNew()
	allocateWithMake()
	allocateWithMakeBySlice()
}
