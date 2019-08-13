package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
pass by value
func (v Vertex) Abs() float64 {
	v.X = 4
	v.Y = 6
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/


// pass by reference
func (v Vertex) Abs() float64 {
	v.X = 4
	v.Y = 6
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.X)

}






/*
Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
*/
