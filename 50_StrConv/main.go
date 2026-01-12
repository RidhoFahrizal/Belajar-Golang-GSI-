package main

import (
	"fmt"
	"strconv"
)

func main(){
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}else{
		fmt.Println(err.Error())
	}
	
	var foo []int

	foo = append(foo, 1,2,3,4,5,6,7)

	num, err := strconv.Atoi("jsdfdfkdfkd")

	if err != nil {
		fmt.Println(fmt.Errorf("tolong masukkan karakter angka"))
	}
	fmt.Println(len(foo))
	fmt.Println(num)
}