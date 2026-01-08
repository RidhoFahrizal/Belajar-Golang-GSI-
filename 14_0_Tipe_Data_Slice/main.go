package main

import (
	"fmt"
)

func main(){

	var months = [...]string{
		"Januari", 
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli", 
		"Agustus",
		"September",
		"Oktober", 
		"November",
		"Desember",
	}

	// aku melakukan 
	var slice2 = months[2:4:4]
	fmt.Println("data slice2","len", len(slice2),"cap", cap(slice2))
	fmt.Println("Ini slice month[10:]", slice2)


	// ini praktik yang sangat penting dalam penggunaan slice 
	
	var slice3 = append([]string{}, slice2...)// ini adalah praktik yang bagus karena kita membuat array baru 
	// walaupun si array capacity nya masih bagus dia akan membuat array baru .









	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"

	fmt.Println(slice3)
	fmt.Println("Ini slice2:",slice2)
	fmt.Println("terjadi perubahan di moths:", months)


	newSlice := make([]string, 2,5)
	newSlice[0] = "Ridho"
	newSlice[1] = "Fahrizal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice), cap(newSlice))

	copySlice := make([]string, 0, cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)


	//kesimpulan : hati - hati dalam membuat array 
	iniArray := [5]int{1,2,1,2,1}
	iniSlice := []int{1,2,1,2,1}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}