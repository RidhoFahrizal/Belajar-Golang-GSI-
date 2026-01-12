package main


import(
	"bufio"
	"strings"
)



func ReadLine(reader * bufio.Reader)(string, error){
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err 
	}

	return strings.TrimSpace(text), nil 
}