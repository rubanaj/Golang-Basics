package main

import "fmt"

var ab int

type burger int
var b burger

func main(){
	ab = 42
	b = 50
	fmt.Println(ab)
	fmt.Printf("%T\n",ab)

	fmt.Println(b)
	fmt.Printf("%T\n",b)

	ab=int(b)

	fmt.Println(ab)
	fmt.Printf("%T\n",ab)

}
