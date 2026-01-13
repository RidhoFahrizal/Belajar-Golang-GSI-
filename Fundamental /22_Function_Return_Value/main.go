package main 

import (
	"fmt"
)

func foo (firstname string, lastname string)string{
	return firstname + lastname;
}

func main (){
	var fooName = foo("Ridho", "Fahrizal")
	fmt.Println(fooName)
}