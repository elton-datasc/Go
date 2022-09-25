package main

import (
	"fmt"
)
// varivável iniciada e não instanciada sempre terá valor zero
var a int
var b float
var c string
var d bool

func main() {

	fmt.Println("%v, %T\n", a, a)
	fmt.Println("%v, %T\n", b, b)
	fmt.Println("%v, %T\n", c, c)
	fmt.Println("%v, %T\n", d, d)

}