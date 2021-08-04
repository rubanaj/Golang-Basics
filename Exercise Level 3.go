package main

import "fmt"

var exx = 42
var exy = "James Bond"
var exz = true

func main(){

	s := fmt.Sprintf("%d is %s age and its %t", exx, exy, exz)
	fmt.Println(s)


}
