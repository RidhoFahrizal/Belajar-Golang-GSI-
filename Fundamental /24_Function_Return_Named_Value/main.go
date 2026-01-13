package main

import(
	"fmt"
)

func getFullName1()(firstname string, lastname string, nrp string){
	firstname = "Ridho"
	lastname  = "Fahrizal"
	nrp 	  = "3123600043"
	return
}

func main (){
	foo,_,_ := getFullName1()
	fmt.Print(foo)
}