package main

import (
	"fmt"
	"time"
)

func AddTodo(tittle string){	
	todos = append(todos, Todo{
		ID: len(todos) + 1,
		Title: tittle,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDone: false,
	})
}

func ListTodos(){
	for _, t := range todos{
		fmt.Println("id:", t.ID )
		fmt.Println("title:", t.Title)
		if t.IsDone == true{
			fmt.Println("Status: Sudah Selesai ")
		}else{
			fmt.Println("Status: Belum Dikerjakan")
		}
		fmt.Println("Crated at:", t.CreatedAt)
		fmt.Println("Updated at:", t.UpdatedAt)
	}
}