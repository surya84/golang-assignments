package person

import (
	"c-question/model"
	"fmt"
)

func PrintPerson(u model.Person) {
	fmt.Println("Name : ", u.Name)
	fmt.Println("Age : ", u.Age)
}
