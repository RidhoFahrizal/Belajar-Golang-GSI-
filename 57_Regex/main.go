package main

import (
	"fmt"
	_ "fmt"
	"regexp"
)

func main (){
	var regex = regexp.MustCompile(`e(a-z)o`)

	fmt.Println(regex.MatchString("ridho"))
	fmt.Println(regex.MatchString("ridha"))
	fmt.Println(regex.MatchString("steven"))

	fmt.Println(regex.FindAllString("ridho joko ridha Steven", 10))
}