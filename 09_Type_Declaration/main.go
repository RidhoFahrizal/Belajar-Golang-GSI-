package main

import (
	"fmt"
	"reflect"
)



func main(){
	
	type (
		NoKTP string 
		Married bool
		Graduate bool
	)

	var KtpRidho NoKTP = "3123600043"
	fmt.Println(KtpRidho)
	fmt.Println("Type KtpRidho:", reflect.TypeOf(KtpRidho))

	var isMarried Married = false
	fmt.Println(isMarried)
	fmt.Println("Type isMarried:", reflect.TypeOf(isMarried),)
	
	var isGraduate Graduate= false
	fmt.Println(isGraduate)
	fmt.Println("Type isGraduate", reflect.TypeOf(isGraduate))
	
	// dia mengulang variable yang sudah kita buat dengan variable bawaan golang
}