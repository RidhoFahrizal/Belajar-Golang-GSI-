package main 

import (
	"fmt"
)

type User struct {
	Name  string `validate:"required"`
	Age   int    `validate:"min=18"`
	Email string `validate:"email"`
}

func PrintValidationRules(v interface{}) {
	t := reflect.TypeOf(v)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("validate")

		if tag != "" {
			fmt.Printf("%s → %s\n", field.Name, tag)
		}
	}
}



func isValid


func main (){

}