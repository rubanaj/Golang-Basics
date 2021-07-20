package main

import "fmt"
//go automatically adds a ; at the end of a statement so main() { is the correct way of declaring a main()0

var z = 3
//declares variable z of type int and value 0 is assigned to it
var r int

func main() {
	//short declaration
	//works only inside a func()
	x := 10
	fmt.Println(x)

	// Var can be declared outside a func() and can be used within a func()
	var y = 20
	fmt.Println(y)
	fmt.Println(z)

	foo()

	fmt.Println(r)
}


func foo(){
	fmt.Println("test it again",z)
}
