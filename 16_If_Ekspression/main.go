package main



import (
	"fmt"
)

func main(){
	var name = "ridho"

	if name == "ridho"{
		fmt.Println("DO YOUR TASK!!")
	}else if name == "michael"{
		fmt.Println("Do you know ridho? i'm looking for him")
	}else{
		fmt.Println("Who are you? ")
	}

	if lenght  := len(name); lenght > 5{
		fmt.Println("Terlalu Panjang")
	}else {
		fmt.Println("nama sudah benar ")
	}
}



