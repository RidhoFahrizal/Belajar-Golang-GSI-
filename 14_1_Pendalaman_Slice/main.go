package main

import (
	"fmt"
)

func main(){

	/**
	Slice itu struktur data yang dinamis bahkan bisa dibuat tanpa 
	mengisialisasikan ukurannya
	*/
	var s[]string
	fmt.Println("uninit:", s, s == nil, len(s) == 0 )
	/**
	slice ditas itu nil , dan ukurannya juga nil.        
	*/




	/**
	kita bisa pake kata kunci make dalam membuat slice, agar slice nya tidak memiliki 
	zero length. 
	*/
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("set:", s)
	
	/**
	slice juga bisa di copy
	
	*/

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:",l)
}

