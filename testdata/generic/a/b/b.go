package b

import "a/c"

type A[T any] struct {
	ID string
	T  T
}

type B struct {
	B string
}

type C struct {
	C string
}

type AB = A[B]
type AC = A[c.C]
type Astring = A[string]
type AstringPtr = A[*string]
