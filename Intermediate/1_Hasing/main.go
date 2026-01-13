package main


import (
	"crypto/sha256"
	"fmt"
)

func main(){
	text := "hello world"
	hash := sha256.Sum256([]byte(text))
	fmt.Printf("Sha-256: %x\n", hash)
	
}