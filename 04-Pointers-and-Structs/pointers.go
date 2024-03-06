package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	pVertex  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	i, j := 42, 2701

	ptr := &i // point to i
	fmt.Println(*ptr) // read i through the pointer
	fmt.Println(ptr)
	*ptr = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	ptr = &j         // point to j
	*ptr = *ptr / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex {X: 1, Y: 2}
	v.X = 4
	fmt.Println(v);
	pStruct := &v
	// Not needed: Go automatically dereferences
	// (*pStruct).X = 1e9
	pStruct.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, pVertex, v2, v3);
}