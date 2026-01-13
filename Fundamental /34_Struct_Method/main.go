package main

import(
	"fmt"
)

type Foo struct {
	fooVar string
}


/**
	disni saya mempelajari bagainana si struct bisa mempuyai method nya sendiri 
	jadi si struct mempunyai beheaviornya sendiri 
*/
func (foo Foo) sayFoo(){
	fmt.Println("hallo", foo.fooVar)
}

func main (){
	var ridho Foo
	ridho.fooVar = "jack"
	ridho.sayFoo()
}