package main

import (
	"fmt"
	"reflect"
)

func foo(fooVar interface{}) interface{}{
	return fooVar
}

func main(){
	fmt.Println(foo(8), "Tipe nya",reflect.TypeOf(foo))
}