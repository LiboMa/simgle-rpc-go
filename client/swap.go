package main

import "fmt"

func swap(a *int, b *int) error {
	*a, *b = *b, *a
	return nil

}
func main() {

	A := 1
	B := 99

	fmt.Printf("before pointer swap func -> A: %d, B: %d\n", A, B)
	swap(&A, &B)
	fmt.Printf("after pointer swap func -> A: %d, B: %d\n", A, B)
}
