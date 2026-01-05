

package main

import (

	"fmt"
	"reflect"
)

func main(){
	var name string 

	name ="ridho fahrizal"

	fmt.Println(name)

	name = "FOO NAME THIS IS FOOOOO"

	fmt.Println(name)


	foo := 933

	fmt.Print("This is foo")
	fmt.Print(" Disni foo sebagai ")
	fmt.Println(reflect.TypeOf(foo))





	//Deklarasi Multiple variable
	var (
		firstName = "Ridho"
		lastName = "Fahrizal"
	)

	fmt.Println("First name:",firstName)
	fmt.Println("Last name:",lastName)
}