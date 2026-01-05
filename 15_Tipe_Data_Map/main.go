package main 

import (
	"fmt"
)

func main(){
	ages := map[string]int{
		"A": 30,
		"B": 25,
		"C": 35,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))
	delete(ages, "A")
	fmt.Println(len(ages))
	fmt.Println(ages)
	
	
	var book = make(map[string]string)
	book["tittle"] = "Belajar Golang"
	book["Author"] = "Ridho Fahrizal"
	book["foo"] = "foo"

	fmt.Println(book)
	delete(book, "foo")

	fmt.Println(book)
}