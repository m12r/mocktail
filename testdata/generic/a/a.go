package a

import "a/b"

type A struct {
	Foo string
}

type AA = b.A[A]

type AAer interface {
	GetAA(id string) (*AA, error)
}

type ABer interface {
	GetAB(id string) (*b.AB, error)
}

type ACer interface {
	GetAC(id string) (*b.AC, error)
}

type AStringer interface {
	GetAString(id string) (*b.Astring, error)
}

type AStringPtrer interface {
	GetAStringPtr(id string) (*b.AstringPtr, error)
}
