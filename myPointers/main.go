package main

import "fmt"

func main() {
	fmt.Println("Hello pointers!")
	var ptr *int
	fmt.Println(ptr)
	number := 53
	myptr := &number
	fmt.Println("the actual pointer :", myptr)
	fmt.Println("the value of the pointer :", *myptr)
	*myptr = *myptr + 3
	fmt.Println("the new value of pointer :", *myptr)
}
