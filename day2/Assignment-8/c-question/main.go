package main

import (
	"c-question/model"
	"c-question/person"
)

func main() {

	u := model.Person{
		Name: "Surya",
		Age:  22,
	}

	person.PrintPerson(u)

}
