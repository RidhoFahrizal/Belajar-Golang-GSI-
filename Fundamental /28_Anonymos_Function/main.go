package main

import(
	"fmt"
)

type Blacklist func(string) bool 
func registerUser (name string, blacklist Blacklist){
	
	if blacklist(name){
		fmt.Println("Your are Blocked", name)
	}else {
		fmt.Println("Welcom", name)
	}
}
func main (){
	/**
	disini kita menggunakan function tanpa nama jadi langsung di 
	masukkan ke variable.  
	contohnya ini : 
	*/

	blacklist := func (name string) bool  {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", blacklist)
	registerUser("ridho", blacklist)

}