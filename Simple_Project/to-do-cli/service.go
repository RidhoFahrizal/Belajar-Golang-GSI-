package main

import (
	"fmt"
	"os"
	_"slices"
	"time"
	"github.com/olekukonko/tablewriter"
)
/**
1. PASTIKAN DALAM SERVICE INI PASTIKAN DIA BERPAKU PADA INDEKS
2. SEGALA JENIS UPDATE SERTA DELETE BERPAKU PADA INDEKS
*/




func AddTodo(tittle string){	
	todos = append(todos, Todo{
		Title: tittle,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDone: false,
	})
}


func UpdateTittleTodo(index int, title string){
	for i := range todos {
		if i == index{
			todos[i].Title = title
			return   
		}
	}
}


func DeleteTodo(index int) {
	index -= 1 
	if index < 0 || index >= len(todos) {
		return
	}
	todos = append(todos[:index], todos[index+1:]...)
}

func MarkTaskIsDone(index int){
	index -= 1
	if index < 0 || index >= len(todos) {
		return
	}
	todos[index].IsDone = true
}

func ListTodos() {
	
	index :=1
	table := tablewriter.NewWriter(os.Stdout)


	table.Header(
		"ID",
		"TITLE",
		"STATUS",
		"CREATED AT",
		"UPDATED AT",
	)

	for _, t := range todos {
		status := "❌ BELUM"
		if t.IsDone {
			status = "✅️ SELESAI"
		}

		table.Append(
			fmt.Sprint(index),
			t.Title,
			status,
			t.CreatedAt.Format("2006-01-02"),
			t.UpdatedAt.Format("2006-01-02"),
		)

		index++
	}

	table.Render()
}

