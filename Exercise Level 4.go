package main

import "fmt"

type froyo int
var bx froyo

func main(){
	fmt.Println(bx)
	fmt.Printf("%T\n",bx)

	bx = 42

	fmt.Println(bx)
}
