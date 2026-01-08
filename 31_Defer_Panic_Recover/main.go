package main

import (
	"errors"
	"fmt"
)


func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(x int, y int ){
	defer logging()
	fmt.Println("Run Application")
	if y == 0{
		panic(errors.New("Tidak boleh 0"))
	}
}

// func sayHello(){
// 	fmt.Println("hello")
// }

func main(){
	runApplication(4,0)
	sayHello()
	message := recover()
	fmt.Println(message)

}