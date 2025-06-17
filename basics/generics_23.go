package main

import "fmt"

// Without Generics
func has(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// With Generics
// type comparable interface{ comparable }
// comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint, not as the type of a variable.
func generic[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
func main() {
	fmt.Printf("Present or not: %v\n", has([]string{"1", "2"}, "1"))
	fmt.Printf("Present or not: %v\n", has([]string{"1", "2"}, "6"))

	// with generics we can pass any type of variable
	fmt.Printf("Present or not: %v\n", generic([]string{"1", "2"}, "1"))
	fmt.Printf("Present or not: %v\n", generic([]float32{1.8, 7, 1.2}, 1.1))
	fmt.Printf("Present or not: %v\n", generic([]uint{1, 3}, 5))

	// Determining T implicitly isn’t always possible
	value1 := NilPtr[int]() // OK
	fmt.Println(value1)
	value2 := NilPtr[struct{}]() // OK
	fmt.Println(value2)
	// value3 := NilPtr()  // The last example gives us the compile error: "cannot infer T" : Compile Error

	f32List1 := MyList[float32]{0.0, 1.5} // OK
	fmt.Println(f32List1)
	// f32List2 := MyList{0.0, 1.5}    // Compile Error

	// map[MyType]struct{} SAME AS Below but with Generics
	// type Set[T any] struct{ m map[T]struct{} }

	f64Set := NewSet(1.2, 3.5, 7.75)
	fmt.Printf("%v: %v \n", 1.2, f64Set.Has(1.2)) // 1.2: true
	fmt.Printf("%v: %v \n", 5.6, f64Set.Has(1.2)) // 5.6: false

}
func NilPtr[T any]() *T {
	return nil
}

type MyList[T any] []T

type MyStruct[T any] struct {
	a T
	// ...
}

type Set[T comparable] struct{ m map[T]struct{} }

func (s *Set[T]) Has(item T) bool {
	_, has := s.m[item]
	return has
}

func NewSet[T comparable](values ...T) Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return Set[T]{m: m}
}

// Can only use int - kind of pointless.
// func Has1[T int](items []T, v T) bool

// // Can only use int or uint
// func Has2[T int | uint](items []T, v T) bool

// // Can only use int or types with string as the
// // underlying type.
// func Has3[T int | ~string](items []T, v T) bool
