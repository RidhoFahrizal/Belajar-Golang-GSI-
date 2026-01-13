package main

import (
	"fmt"
)

func main(){
	counter := 1

	for counter <= 10  {
		fmt.Println(counter)
		counter++
	}

	slice := []string{"eko", "Ridho", "Flamming"}
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}


	person := make(map[string]string)
	person ["name"] = "ridho"
	person ["title"] = "Engineer"

	for key, value := range person{
		fmt.Println(key, "=", value)
	}
}