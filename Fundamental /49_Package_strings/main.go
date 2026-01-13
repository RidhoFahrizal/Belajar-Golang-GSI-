package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Ridho Fahrizal", "Ridho Fahrizal")) // boolean, true
	fmt.Println(strings.Contains("Ridho Fahrizal", "foo")) // boolean, flase
	fmt.Println(strings.Split("Ridho Fahrizal", " "))      // [Ridho Fahrizal]
	fmt.Println(strings.ToUpper("ridho fahrizal")) 		   // RIDHO FAHRIZAL
	fmt.Println(strings.ToLower("RIDHO FAHRIZAL")) 		   // ridho fahrizal
	fmt.Println(strings.Trim("    Ridho Fahrizal  ", " ")) // Ridho Fahrizal
	fmt.Println(strings.ReplaceAll("RIDHO FOO SHIT GOO", "RIDHO", "JACK"))
}
