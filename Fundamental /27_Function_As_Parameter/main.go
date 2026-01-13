package main 


import (
	"fmt"
)

type Filter func(string)string 

func sayHellowWithFillter(name string,filter Filter){
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string)string{
	if name == "anjing"{
		return "..."
	}else {
		return name 
	}
}

func main (){
	sayHellowWithFillter("Eko", spamFilter)
	sayHellowWithFillter("anjing",spamFilter )
}

