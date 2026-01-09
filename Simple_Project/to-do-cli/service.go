package main

import (
	"fmt"
	"time"
	"github.com/olekukonko/tablewriter"
	"os"
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


func ListTodos() {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header(
		"ID",
		"TITLE",
		"STATUS",
		"CREATED AT",
		"UPDATED AT",
	)

	for _, t := range todos {
		status := "Belum"
		if t.IsDone {
			status = "Selesai"
		}

		table.Append(
			fmt.Sprint(t.ID),
			t.Title,
			status,
			t.CreatedAt.Format("2006-01-02"),
			t.UpdatedAt.Format("2006-01-02"),
		)
	}

	table.Render()
}

