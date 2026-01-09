package main

import(
	"fmt"	
)
// kita pake util untuk mengambil input me
func main(){
	fmt.Println("Disini akan di tulis project todo-sederhana")
	AddTodo("Pelajari Struct method")
	AddTodo("Pelajari Interface Kosong")
	AddTodo("Pelajari Nil")
	ListTodos()
}