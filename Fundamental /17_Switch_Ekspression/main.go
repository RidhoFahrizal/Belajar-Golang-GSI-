package main 


import (
	"fmt"
)


func main(){
	name := "Ridho"

	switch name {
	case "Ridho":
		fmt.Println("Hellow Ridho")
		fmt.Println("Hellow Ridho")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Mari Berkenalan")
	}

	switch length := len(name); length > 5 {
	case true: 
		fmt.Println("Nama terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
