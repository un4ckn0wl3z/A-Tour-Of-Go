package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X)
	fmt.Println(v.Y)
}

/** 
Struct Fields
Struct fields are accessed using a dot.
*/
