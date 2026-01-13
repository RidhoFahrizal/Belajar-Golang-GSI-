package main 

import (
	"fmt"
)

func main(){
	var name1 = "foo"
	var name2 = "foo"
	var name3 = "foo1"
	var compare bool = name1 == name2
	var compare2 bool = name1 == name3
	
	fmt.Println(compare)
	fmt.Println(compare2)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}