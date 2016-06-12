package main

import (
	"fmt"
)

type A struct {
	aa int
	bb int
}

type B struct {
	cc int
}

type C struct {
	A
	B
}

func main() {
	var c interface{} = new(C)
	var x, ok = c.(*A)
	if !ok {
		fmt.Println("asdasd")

	} else {
				x.aa = 34

	}
	//	switch n := c.(type) {
	//	case *A:
	//		n.aa = 122
	//		n.bb = 343
	//	case *B:
	//		n.cc = 3434
	//	case *C:
	//		n.aa = 34
	//		n.bb = 45
	//		n.cc = 35
	//	default:

	//	}
	fmt.Println(c)

}
