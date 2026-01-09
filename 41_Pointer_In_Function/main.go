package main

import (
	"fmt"

	"github.com/google/uuid"
)


type Student struct{
	Id string
	FirstName  string 
	LastName   string
	NRP string
	Class string
	Major string	
	IsGraduate bool
}

var studentDatabase []*Student

func NewStudent(	
	firstName, lastName, nrp, class, major string, isGraduate bool,
)*Student{
	return &Student{
		Id: uuid.NewString(),
		FirstName: firstName,
		LastName: lastName,
		NRP: nrp,
		Class: class,
		Major: major,
		IsGraduate: isGraduate,
	}
}

func AddStudent(
	firstName, lastName, nrp, class, major string,
	isGraduate bool,
) {
	studentDatabase = append(
		studentDatabase,
		NewStudent(firstName, lastName, nrp, class, major, isGraduate),
	)
}

func ShowStudents() {
	fmt.Println("\n=== Student Database ===")

	if len(studentDatabase) == 0 {
		fmt.Println("No student data")
		return
	}

	for i, student := range studentDatabase {
		fmt.Printf(
			"%d. %s %s | NRP: %s | Class: %s | Major: %s | Graduate: %t\n",
			i+1,
			student.FirstName,
			student.LastName,
			student.NRP,
			student.Class,
			student.Major,
			student.IsGraduate,
		)
	}
}



func main (){

	var firstName, lastName, nrp, class, major string
	var isGraduate bool 
	fmt.Println("masukkan data:")
	fmt.Print("FirstName:")
	fmt.Scanf("%s", &firstName)
	fmt.Print("LastName:")
	fmt.Scanf("%s", &lastName)
	fmt.Print("NPR:")
	fmt.Scanf("%s", &nrp)
	fmt.Print("Class:")
	fmt.Scanf("%s", &class)
	fmt.Print("Major")
	fmt.Scanf("%s", &major)
	fmt.Print("isGraduate")
	fmt.Scanf("%t", &isGraduate)
	

	AddStudent(firstName,lastName,nrp,class,major,isGraduate)	
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	AddStudent(firstName,lastName,nrp,class,major,isGraduate)
	ShowStudents()
	
}


/**
input : 
ridho
fahrizal
3123600043
D4 IT B
Informatics Engineering
true

*/