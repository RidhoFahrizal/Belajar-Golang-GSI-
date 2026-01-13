package main 

import (
	"fmt"
)

/**
	Varags bisa digunakan untuk mengimput dynamic data
	semacam slice 

*/


// variadic parameter harus diletakkan di posisi paling belakang 
// (foo ...int, fooVar ) ❌ ini gaboleh dan si variadic harus di letakkan di yang terakhir
// (fooVar, foo ...int)  ✅ Ini benar 

func Total(foo ...int)int{
	total := 0
	for _, fooVar:= range foo {
		total += fooVar 
	}
	return total
}

func main(){
	fmt.Println(Total())

	/**
	kalau kamu sudah memiliki slice 
	kita bisa masukkan dia sebagai parameter 
	dengan menggunakan variable slice ditambah ...

	contohnya : func(foo1,foo2,foo...)
	*/
	slice := []int{1,2,3,4,5,6} 	
	fmt.Println(Total(slice...))
	 
}