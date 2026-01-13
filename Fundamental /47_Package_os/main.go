package main

import (
	"fmt"
	_ "fmt"
	"os"
	_ "os"
)

func main (){
	/**
	bisa ngotak ngatik sistem operasi 
	*/

	args := os.Args
	println(args)


	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("Hostname", hostname)

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWOTD")

	fmt.Println(username)
	fmt.Println(password)
}