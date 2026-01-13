package main

import (
	"fmt"
)

func main(){
	var foo1 = 10+10 
	fmt.Println(foo1)

	var a = 18
	var b = 18
	var c = a + b 
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)


	var foo = 9 
	foo++ // artinya foo + 1
	fmt.Println(foo)
}