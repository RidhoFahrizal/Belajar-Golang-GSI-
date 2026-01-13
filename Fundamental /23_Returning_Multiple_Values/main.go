package main 


import (
	"fmt"
)

func foo ()(string, string){
	return "jack", "harlow"
}

func main(){
	firstname, lastname := foo()
	fmt.Println(firstname, lastname)

	foo1, _ := foo()
	fmt.Println(foo1)	
}