package main

import (
	"time"
)

type Todo struct{
	ID int
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time 
	IsDone bool
}

//gunakan slice sebagai database
var todos []Todo 


