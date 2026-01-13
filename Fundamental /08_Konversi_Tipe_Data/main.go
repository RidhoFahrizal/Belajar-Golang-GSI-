package main

import (

	"fmt"
	"reflect"
)

func main (){
	//disini kita akan mempelajari konversi tipe data
	
	var foo1 int64 = 128 
	var foo2 int32 = int32(foo1)
	var foo3 int16 = int16(foo2)
	var foo4 int8  = int8(foo3)


	// disini kita bisa lihat bahwa type datanya terkonversi ke hal yang lain 

	fmt.Println(foo1,reflect.TypeOf(foo1) )
	fmt.Println(foo2,reflect.TypeOf(foo2) )
	fmt.Println(foo3,reflect.TypeOf(foo3) )
	fmt.Println(foo4,reflect.TypeOf(foo4) )// 


	// practice 
	var name = "Ridho"
	var e  byte  = name[0]


	fmt.Println("this is e:",e, `dia adalah karakter aschii dari name`, reflect.TypeOf(e))

	var eString = string(e)
	fmt.Println("this is eString:", eString, `dia adalah representasi dari aschii e (82)`, reflect.TypeOf(eString))


}