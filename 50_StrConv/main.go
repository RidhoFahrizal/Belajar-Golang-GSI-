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
	

	num, err := strconv.Atoi("4339284987498274")
	fmt.Println(num)
}