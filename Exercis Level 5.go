package main

import "fmt"

type boba int
var cx boba
var cy int


func main(){


	fmt.Println(cx)
	fmt.Printf("%T\n",cx)

	cx = 42

	fmt.Println(cx)

	cy = int(cx)
	fmt.Printf("%T\n",cy)


}