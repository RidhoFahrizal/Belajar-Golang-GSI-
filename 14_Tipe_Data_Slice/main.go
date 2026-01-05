package main

import (
	"fmt"
)

func main(){
	
	s := []int{1,2,3}
	t := s 

	t[0] = 99
	fmt.Println(s)

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

	fmt.Println(months)
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[2:4]
	fmt.Println("Ini slice month[10:]", slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println("Ini slice2:",slice2)





	fmt.Println("ini months:", months)



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