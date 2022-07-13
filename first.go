package main

import (
	"fmt"
)

func swap(x *int, y *int) (*int, *int) {
	*x, *y = *y, *x
	return x, y
}

func main() {
	a, b := 5, 10
	fmt.Printf("a: %d b: %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("a: %d b: %d", a, b)
}
