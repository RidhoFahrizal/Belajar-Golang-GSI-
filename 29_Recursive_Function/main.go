package main 

import (
	"fmt"
)
// dengan recursive pastikan ada kondisi untuk berhenti 
func factorialLoop(value int) int {
	result := 1
	for i := value; i>0 ; i-- {
		result *= i 
	}
	return result
}

func factorialLoopRecursive(value int)int{
	if value == 1 {
		return 1
	}else {
		return value * factorialLoopRecursive(value-1)
	}
}
func main(){

	foo:= factorialLoop(6)
	fmt.Println(foo)
	// kasus factorial
	fooRecursive := factorialLoopRecursive(6)
	fmt.Println(fooRecursive)	
}