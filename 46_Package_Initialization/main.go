package main

import (
	"fmt"
	"Database/database"
	_"os"		 // dipanggil tapi tidak digunakan, bisa dengan cara _ (blank identifier)
	_"net/http"  // dipanggil tapi tidak digunakan, bisa dengan cara _ (blank identifier)
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
