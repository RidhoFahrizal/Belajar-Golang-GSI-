package main

import (
	"fmt"
)



func SayHelloTo(firstname string, lastname string){
	fmt.Println("Hello", firstname, lastname)
}

func main(){
	SayHelloTo("Ridho", "Fahrizal")
}


