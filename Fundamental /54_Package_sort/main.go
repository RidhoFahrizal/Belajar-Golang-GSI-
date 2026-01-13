package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	age  int
}

type userSlice []user

func (value userSlice) Len() int {
	return len(value)
}

func (value userSlice) Less(i, j int) bool {
	return value[i].age < value[j].age
}

func (value userSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []user{
		{"ridho", 20},
		{"jack", 19},
		{"nani", 12},
		{"fame", 5},
		{"key", 24},
	}

	fmt.Println(users)
	sort.Sort(userSlice(users))
	fmt.Println(users)
}