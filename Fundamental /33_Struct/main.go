package main

import (
	"fmt"
)

type Customer struct{
	Name 			string
	Age 			uint8
	Address 		string
	PremiumAccess   bool
}

func main(){
	var Ridho Customer

	Ridho.Name 			= "Ridho Fahrizal"
	Ridho.Age  			= 20
	Ridho.Address		= "Simo Kwagean Kuburan no 26"
	Ridho.PremiumAccess	= true

	fmt.Println(Ridho)


	Joko := Customer{
		Name		  : "Joko Widodo",
		Age 		  : 30 ,
		Address		  : "Kupang Jaya no 69",
		PremiumAccess : false,
	}

	fmt.Println(Joko)
}


