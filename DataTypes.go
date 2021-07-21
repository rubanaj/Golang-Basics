package main

import "fmt"
var number1 = 10

var a = "I wish i know what's rune type"
// GO is a static programming language not a dynamic one


//facing an error here
//var b string = 'Ruban said, "He wants to know"'

func main(){

	fmt.Println(number1)
	fmt.Printf("%T\n",number1)

	fmt.Println(a)
	fmt.Printf("%T\n",a)

	//fmt.Println(b)
	//fmt.Printf("%T\n",b)

}
