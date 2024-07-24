package models

import "fmt"

type Customer struct {
	Id       int
	Login    string
	Password string
	Name     string
	Surname  string
	Age      int
}

func (c Customer) String() string {
	return fmt.Sprintf("id: %d, login: %s, password: %s, name: %s, surname: %s, age: %d",
		c.Id, c.Login, c.Password, c.Name, c.Surname, c.Age)
}
