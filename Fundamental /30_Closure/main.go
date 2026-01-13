package main

import (
	"fmt"
	"reflect"
	"net/http"
)


func makeCounter(n *int) func() int {
     // The variable "n" is captured by the inner function
	 
	 return func() int {
        *n++
        return *n
    }
}

func makeHelloHandler(appName string) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from ", appName)
	}
}

func main() {

	 n := 0; 
	/**
	counter1 menyimpan sebuah closure yang terdiri dari satu function dan satu environment.
	Function-nya sama untuk semua instance closure, sedangkan environment menyimpan variabel n yang di-allocate di heap.
	Setiap pemanggilan makeCounter() membuat environment baru sehingga state n bersifat independen.
	*/
    counter1 := makeCounter(&n)
    fmt.Println(counter1()) // Output: 1
    fmt.Println(counter1()) // Output: 2

	
	println("Penting di perhatikan bahwa Counter 1 typenya adalah: ", reflect.TypeOf(counter1))


    counter2 := makeCounter(&n)
    fmt.Println(counter2()) // Output: 1 (counter2 has its own independent 'n')



	/**
	
	salah satu pengimplementasian dari closure adalah  http handler,
	HandleFunc adalah fungsi biasa yang bisa mengembalikan sebuah fungsi http 
	jadi de
	*/
	http.HandleFunc("/hello", makeHelloHandler("MyApp"))
	http.ListenAndServe(":8080", nil)

}
