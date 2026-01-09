package main

import (
	"fmt"
)

func random()interface{}{
	return "foo"
}

func main (){


	/**
	
	Kemampuan menrubah tipe data yang diinginkan 
	sering ditemukan saat bertemu data interface kosong
	*/

	result := random()

	switch value := result.(type){
	case string :
		fmt.Println("Value", value, "is String")
	case int :
		fmt.Println("Value", value, "is int" )
	case float64 :
		fmt.Println("Value", value, "is float" )
	case bool :
		fmt.Println("Value", value, "is bool" )
	default:
		fmt.Println("unkknown")
	}
}