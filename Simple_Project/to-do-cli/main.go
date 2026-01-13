package main

import (
	"bufio"
	"fmt"
	_ "go/printer"
	"os"
	"strconv"
	_ "strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin) 
	fmt.Println("Disini akan di tulis project todo-sederhana")
	AddTodo("Pelajari Struct method")
	AddTodo("Pelajari Interface Kosong")
	AddTodo("Pelajari Nil")
	MarkTaskIsDone(2)
	ListTodos()
	
	for  {
		fmt.Println("====== TODO ======" )
		fmt.Println("1. Add todo")		
		fmt.Println("2. Update todo")
		fmt.Println("3. Mark Taks Is Done" )
		fmt.Println("4. List Todo")
		fmt.Println("5. Delete Todo")
		fmt.Println("Tekan enter untuk keluar")
		fmt.Print("Pilih Menu :")
		input, err := ReadLine(reader)
		if err != nil {
			println("terjadi error saat membuat readline")
		}	
		if input == "" {
			fmt.Println("Keluar dari program")
			break
		}
		switch input{
		case "1": 
			fmt.Print("Isi Task : ")
			foo, err := ReadLine(reader)
			if  err != nil{
				fmt.Println("Terjadi kesalahan saat membuat ReadLine")
			}
			AddTodo(foo)
		case "2": 
			fmt.Println("isi menu")
			foo, err := ReadLine(reader)
			if  err != nil{
				fmt.Println("Terjadi kesalahan saat membuat ReadLine")
			}
			UpdateTittleTodo(1, foo)
		case "3":
			var fooInt int
			for{
				fmt.Print("Tugas Mana yang sudah :")
				foo, err := ReadLine(reader)
				if err != nil {
					fmt.Println("gagal membaca input")
					continue
				}
				fooInt, err = strconv.Atoi(foo)
				if err != nil {
					fmt.Print(fmt.Errorf("Tolong masukkan angka "))
					continue
				}
				
				break
			}
			if fooInt > len(todos) {
				fmt.Println("tolong isi angka yang sesuai")
				break
			}
			MarkTaskIsDone(fooInt)
		case "4": 
			ListTodos()
		case "5":
			var fooInt int
			for{
				fmt.Print("Masukkan tugas yang Ingin di hapus:")
				foo, err := ReadLine(reader)
				if err != nil {
					fmt.Println("gagal membaca input")
					continue
				}
				fooInt, err = strconv.Atoi(foo)
				if err != nil {
					fmt.Print(fmt.Errorf("Tolong masukkan angka "))
					continue
				}
				break
			}
			if fooInt > len(todos) {
				fmt.Println("tolong isi angka yang sesuai")
				break
			} 
			DeleteTodo(fooInt)
		default: 
			fmt.Println("Tolong pilih sesuai menu ")
		}
	
	}
}
