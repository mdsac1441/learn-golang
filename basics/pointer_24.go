package main

import "fmt"

// &i gives the address with type *int, which is a pointer and expected by the function func increase(i *int).
// *i is the value the pointer points to.

func increase(i int) {
	i = i + 1
}

type Rectangle struct {
	a, b int
}

func (r *Rectangle) doubleIt() {
	// r.b is the same as (*r).b in this context, but it is easier to read.
	// Pointers are important.
	(*r).a *= 2 // r.a *= 2
	(*r).b *= 2 // r.b *= 2
}

func main() {
	fmt.Printf("HI\n")

	i := 0
	fmt.Println(i)
	increase(i)
	fmt.Println(i)

	r1 := Rectangle{a: 1, b: 2}
	fmt.Printf("%d\n%d\n", r1.a, r1.b)
	r1.doubleIt()
	fmt.Printf("%d\n%d\n", r1.a, r1.b)

}
