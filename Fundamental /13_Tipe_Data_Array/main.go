package main

import "fmt"

func main() {

	// array di dalam golang itu tidak berisikan pointer 
	// di dalam array dia adalah copy dari value yang ada di memori 
	var foo = [3]string{
		"foo1", 
		"foo2", 
		"foo3",
	}

	for i := 0; i < 3; i++ {
		fmt.Println(foo[i])
	}

	// ini akan mengeprint panjang array 
	fmt.Println(len(foo)) 
}