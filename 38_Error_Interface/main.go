package main

import(
	"errors"
	"fmt"
)

func Pembagian(nilai int , pembagi int)(int , error){
	if pembagi == 0{
		return 0, errors.New("Pembagian tidak boleh dengan NOL")
	}else{
		return nilai/pembagi,nil
	}
}

func main(){
	foo, err := Pembagian(4,0)
	if err != nil {
		fmt.Println("terjadi kesalahan:", err.Error())
		return
	}
	fmt.Println(foo)
}