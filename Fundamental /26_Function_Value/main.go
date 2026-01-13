package main 

import(
	"fmt"
)
/**
	disini kita menyimpan function di dalam variable 
	disini variable ini bisa mewakili si function ini 
*/

func foo2(fooString string)string {
	return fooString
}

func main(){
	foo := func()  {
		fmt.Println("Say Hellow")
	}

	foo()
	foo3 := foo2 
	fmt.Println(foo3("String"))
	
}