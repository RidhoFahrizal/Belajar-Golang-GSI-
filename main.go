package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Farrel")
	data.PushBack("Zikri")
	data.PushBack("Lerri")

	var cin string

	for cin != "." {
		for e := data.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value)
			if e.Next() != nil {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()

		fmt.Printf("Masukkan nama (Ketik . untuk keluar): ")
		fmt.Scanf("%s", &cin)
		data.PushBack(cin)
		fmt.Println()
	}

	fmt.Println("\nData depan:", data.Front().Value)
	fmt.Println("Data belakang:", data.Back().Prev().Value)
}