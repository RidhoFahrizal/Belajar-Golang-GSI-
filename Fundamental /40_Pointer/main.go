// package main

// import (
// 	"errors"
// 	"fmt"
// )


// func Devide(num1 int, num2 int)(int, error){
// 	if num2 == 0 {
// 		return 0, errors.New("slaah")
// 	}else{
// 		return int(num1/num2), nil 
// 	}
// }

// func main (){
	
	
// 	var num1, num2 int 
// 	fmt.Print("Angka1:")
// 	fmt.Scanf("%d", &num1)
// 	fmt.Print("Angka2:")
// 	fmt.Scanf("%d", &num2)
// 	print("num 1:",num1, "num 2:",num2)
// 	value, err := Devide(num1, num2)
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	fmt.Println("\n","Ini value",value)
// }

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var angka1, angka2 int
	fmt.Print("Masukkan angka1: ")
	fmt.Scanf("%d", &angka1)
	fmt.Print("Masukkan angka2 sebagai pembagi: ")
	fmt.Scanf("%d", &angka2)

	hasil, err := Pembagian(angka1, angka2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}